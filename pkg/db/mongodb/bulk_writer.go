
package mongodb

import (
	"context"
	"sync"
	"time"

	"go.uber.org/zap"
)

// BulkWriter handles batched insert of sensor data for high throughput.
// ----------------------------------------------------------------
// BulkWriter จัดการการ insert แบบ batch สำหรับข้อมูลเซนเซอร์เพื่อปริมาณงานสูง
type BulkWriter struct {
	repo      SensorRepository
	batchSize int
	flushInterval time.Duration
	buffer    []SensorReading
	mu        sync.Mutex
	stopCh    chan struct{}
}

// NewBulkWriter creates a new bulk writer with batching.
// ----------------------------------------------------------------
// NewBulkWriter สร้าง bulk writer ใหม่พร้อมการทำ batch
func NewBulkWriter(repo SensorRepository, batchSize int, flushInterval time.Duration) *BulkWriter {
	return &BulkWriter{
		repo:          repo,
		batchSize:     batchSize,
		flushInterval: flushInterval,
		buffer:        make([]SensorReading, 0, batchSize),
		stopCh:        make(chan struct{}),
	}
}

// Start begins the background flusher goroutine.
// ----------------------------------------------------------------
// Start เริ่ม goroutine ที่ flush ข้อมูลในพื้นหลัง
func (w *BulkWriter) Start(ctx context.Context) {
	ticker := time.NewTicker(w.flushInterval)
	go func() {
		for {
			select {
			case <-ticker.C:
				w.Flush(ctx)
			case <-w.stopCh:
				ticker.Stop()
				w.Flush(ctx) // final flush, flush ครั้งสุดท้าย
				return
			case <-ctx.Done():
				ticker.Stop()
				w.Flush(ctx)
				return
			}
		}
	}()
}

// Stop gracefully stops the bulk writer and flushes remaining data.
// ----------------------------------------------------------------
// Stop หยุด bulk writer อย่างนุ่มนวลและ flush ข้อมูลที่เหลือ
func (w *BulkWriter) Stop() {
	close(w.stopCh)
}

// Add adds a sensor reading to the buffer, flushing if batch is full.
// ----------------------------------------------------------------
// Add เพิ่มข้อมูลเซนเซอร์ลง buffer และ flush ถ้าเต็ม batch
func (w *BulkWriter) Add(ctx context.Context, reading SensorReading) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.buffer = append(w.buffer, reading)
	if len(w.buffer) >= w.batchSize {
		return w.flush(ctx)
	}
	return nil
}

// Flush writes all buffered readings to MongoDB.
// ----------------------------------------------------------------
// Flush เขียนข้อมูลทั้งหมดใน buffer ลง MongoDB
func (w *BulkWriter) Flush(ctx context.Context) error {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.flush(ctx)
}

func (w *BulkWriter) flush(ctx context.Context) error {
	if len(w.buffer) == 0 {
		return nil
	}
	// Copy buffer to avoid holding lock during write
	// คัดลอก buffer เพื่อไม่ต้องถือ lock ขณะเขียน
	toWrite := make([]SensorReading, len(w.buffer))
	copy(toWrite, w.buffer)
	w.buffer = w.buffer[:0] // clear buffer, ล้าง buffer

	// Write to MongoDB
	// เขียนลง MongoDB
	if err := w.repo.InsertMany(ctx, toWrite); err != nil {
		// Log error and maybe retry
		// บันทึก error และอาจ retry
		// For production, consider dead letter queue
		return err
	}
	return nil
}