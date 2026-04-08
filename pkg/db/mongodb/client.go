
// Package mongodb provides MongoDB client and utilities for time-series sensor data.
// ----------------------------------------------------------------
// แพ็คเกจ mongodb ให้บริการ MongoDB client และ utilities สำหรับข้อมูลเซนเซอร์แบบ time-series
package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

// Config holds MongoDB connection settings.
// ----------------------------------------------------------------
// Config เก็บค่ากำหนดการเชื่อมต่อ MongoDB
type Config struct {
	URI              string        // MongoDB connection URI, e.g., "mongodb://localhost:27017"
	Database         string        // Database name
	MaxPoolSize      uint64        // Maximum connection pool size (default 100)
	MinPoolSize      uint64        // Minimum connection pool size (default 0)
	MaxIdleTime      time.Duration // Maximum idle time for a connection
	ConnectTimeout   time.Duration // Timeout for initial connection
	SocketTimeout    time.Duration // Timeout for socket read/write
	RetryWrites      bool          // Enable retryable writes
	RetryReads       bool          // Enable retryable reads
}

// DefaultConfig returns recommended config for production.
// ----------------------------------------------------------------
// DefaultConfig คืนค่า config ที่แนะนำสำหรับ production
func DefaultConfig() *Config {
	return &Config{
		URI:            "mongodb://localhost:27017",
		Database:       "cmon_sensor",
		MaxPoolSize:    100,              // MongoDB Go driver default is 100[reference:3]
		MinPoolSize:    10,
		MaxIdleTime:    30 * time.Second,
		ConnectTimeout: 10 * time.Second,
		SocketTimeout:  30 * time.Second,
		RetryWrites:    true,
		RetryReads:     true,
	}
}

// Client wraps MongoDB client with connection management.
// ----------------------------------------------------------------
// Client ห่อหุ้ม MongoDB client พร้อมการจัดการการเชื่อมต่อ
type Client struct {
	*mongo.Client
	Database *mongo.Database
	config   *Config
}

// NewClient creates a new MongoDB client with connection pool.
// ----------------------------------------------------------------
// NewClient สร้าง MongoDB client ใหม่พร้อม connection pool
func NewClient(ctx context.Context, cfg *Config) (*Client, error) {
	if cfg == nil {
		cfg = DefaultConfig()
	}

	// Configure connection pool
	// กำหนดค่า connection pool
	clientOpts := options.Client().
		ApplyURI(cfg.URI).
		SetMaxPoolSize(cfg.MaxPoolSize).
		SetMinPoolSize(cfg.MinPoolSize).
		SetMaxConnIdleTime(cfg.MaxIdleTime).
		SetConnectTimeout(cfg.ConnectTimeout).
		SetSocketTimeout(cfg.SocketTimeout).
		SetRetryWrites(cfg.RetryWrites).
		SetRetryReads(cfg.RetryReads)

	// Optional: Add command monitoring for debugging
	// เพิ่มการตรวจสอบคำสั่งสำหรับการดีบัก
	clientOpts.SetMonitor(&event.CommandMonitor{
		Started: func(ctx context.Context, evt *event.CommandStartedEvent) {
			// logger.Debug("MongoDB command", zap.String("cmd", evt.CommandName))
		},
	})

	// Connect to MongoDB
	// เชื่อมต่อ MongoDB
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, err
	}

	// Verify connection
	// ตรวจสอบการเชื่อมต่อ
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return &Client{
		Client:   client,
		Database: client.Database(cfg.Database),
		config:   cfg,
	}, nil
}

// Close gracefully closes the MongoDB connection.
// ----------------------------------------------------------------
// Close ปิดการเชื่อมต่อ MongoDB อย่างนุ่มนวล
func (c *Client) Close(ctx context.Context) error {
	return c.Client.Disconnect(ctx)
}