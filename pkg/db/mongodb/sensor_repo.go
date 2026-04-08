
package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SensorReading represents a single sensor reading document.
// ----------------------------------------------------------------
// SensorReading แทนเอกสารการอ่านค่าจากเซนเซอร์ 1 รายการ
type SensorReading struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	DeviceID   string             `bson:"device_id"`
	SensorType string             `bson:"sensor_type"` // temperature, humidity, water_leak, smoke
	Value      float64            `bson:"value"`
	Unit       string             `bson:"unit"`       // C, %, bool
	Location   string             `bson:"location"`   // rack_a1, room_1, floor_2
	Metadata   map[string]string  `bson:"metadata,omitempty"` // device_name, firmware_version, etc.
	Timestamp  time.Time          `bson:"timestamp"`
}

// SensorRepository defines operations for time-series sensor data.
// ----------------------------------------------------------------
// SensorRepository กำหนดการดำเนินการสำหรับข้อมูลเซนเซอร์แบบ time-series
type SensorRepository interface {
	// InsertOne inserts a single sensor reading.
	InsertOne(ctx context.Context, reading *SensorReading) error

	// InsertMany inserts multiple sensor readings (bulk insert).
	InsertMany(ctx context.Context, readings []SensorReading) error

	// GetReadings retrieves readings for a device within time range.
	GetReadings(ctx context.Context, deviceID, sensorType string, start, end time.Time, limit int) ([]SensorReading, error)

	// GetHourlyAggregation returns aggregated values (avg, min, max) per hour.
	GetHourlyAggregation(ctx context.Context, deviceID, sensorType string, start, end time.Time) ([]HourlyAggregate, error)

	// CreateTimeSeriesCollection creates a time-series collection (if not exists).
	CreateTimeSeriesCollection(ctx context.Context, collectionName string) error
}

// HourlyAggregate represents aggregated sensor data per hour.
// ----------------------------------------------------------------
// HourlyAggregate แทนข้อมูลเซนเซอร์ที่ถูก aggregation รายชั่วโมง
type HourlyAggregate struct {
	Hour       time.Time `bson:"hour"`
	AvgValue   float64   `bson:"avg_value"`
	MinValue   float64   `bson:"min_value"`
	MaxValue   float64   `bson:"max_value"`
	Count      int       `bson:"count"`
}

// sensorRepository implements SensorRepository.
// ----------------------------------------------------------------
// sensorRepository อิมพลีเมนต์ SensorRepository
type sensorRepository struct {
	collection *mongo.Collection
}

// NewSensorRepository creates a new sensor repository.
// ----------------------------------------------------------------
// NewSensorRepository สร้าง sensor repository ใหม่
func NewSensorRepository(db *mongo.Database, collectionName string) SensorRepository {
	return &sensorRepository{
		collection: db.Collection(collectionName),
	}
}

// CreateTimeSeriesCollection creates a time-series collection optimized for sensor data.
// MongoDB 5.0+ only.
// ----------------------------------------------------------------
// CreateTimeSeriesCollection สร้าง time-series collection ที่ optimize สำหรับข้อมูลเซนเซอร์
// ต้องใช้ MongoDB 5.0 ขึ้นไปเท่านั้น
func (r *sensorRepository) CreateTimeSeriesCollection(ctx context.Context, collectionName string) error {
	// Check if collection already exists
	// ตรวจสอบว่ามี collection อยู่แล้วหรือไม่
	names, err := r.collection.Database().ListCollectionNames(ctx, bson.M{"name": collectionName})
	if err != nil {
		return err
	}
	if len(names) > 0 {
		return nil // already exists, มีอยู่แล้ว
	}

	// Create time-series collection
	// สร้าง time-series collection
	// timeField: field that contains the timestamp
	// metaField: field that contains metadata (used for partitioning)
	// granularity: "seconds" for high-frequency data (per minute/second)
	// ----------------------------------------------------------------
	// timeField: ฟิลด์ที่มี timestamp
	// metaField: ฟิลด์ที่มี metadata (ใช้สำหรับการแบ่งพาร์ติชัน)
	// granularity: "seconds" สำหรับข้อมูลความถี่สูง (ทุกนาที/วินาที)
	createCmd := bson.D{
		{Key: "create", Value: collectionName},
		{Key: "timeseries", Value: bson.D{
			{Key: "timeField", Value: "timestamp"},
			{Key: "metaField", Value: "metadata"},
			{Key: "granularity", Value: "seconds"},
		}},
	}

	return r.collection.Database().RunCommand(ctx, createCmd).Err()
}

// InsertOne inserts a single sensor reading.
// ----------------------------------------------------------------
// InsertOne เพิ่มข้อมูลเซนเซอร์ 1 รายการ
func (r *sensorRepository) InsertOne(ctx context.Context, reading *SensorReading) error {
	reading.Timestamp = reading.Timestamp.UTC()
	_, err := r.collection.InsertOne(ctx, reading)
	return err
}

// InsertMany inserts multiple sensor readings using bulk insert.
// This is MUCH faster than individual inserts for high-frequency data.
// ----------------------------------------------------------------
// InsertMany เพิ่มข้อมูลเซนเซอร์หลายรายการแบบ bulk insert
// วิธีนี้เร็วกว่าการ insert ทีละรายการมาก เหมาะกับข้อมูลความถี่สูง
func (r *sensorRepository) InsertMany(ctx context.Context, readings []SensorReading) error {
	if len(readings) == 0 {
		return nil
	}
	// Convert to []interface{} for InsertMany
	// แปลงเป็น []interface{} สำหรับ InsertMany
	docs := make([]interface{}, len(readings))
	for i, reading := range readings {
		reading.Timestamp = reading.Timestamp.UTC()
		docs[i] = reading
	}

	// Use ordered: false for better performance (continue on error)
	// ใช้ ordered: false เพื่อประสิทธิภาพที่ดีขึ้น (ดำเนินการต่อแม้มี error)
	opts := options.InsertMany().SetOrdered(false)
	_, err := r.collection.InsertMany(ctx, docs, opts)
	return err
}

// GetReadings retrieves sensor readings within time range with pagination.
// ----------------------------------------------------------------
// GetReadings ดึงข้อมูลเซนเซอร์ในช่วงเวลาที่กำหนดพร้อมการแบ่งหน้า
func (r *sensorRepository) GetReadings(ctx context.Context, deviceID, sensorType string, start, end time.Time, limit int) ([]SensorReading, error) {
	filter := bson.M{
		"device_id":   deviceID,
		"sensor_type": sensorType,
		"timestamp": bson.M{
			"$gte": start.UTC(),
			"$lte": end.UTC(),
		},
	}

	opts := options.Find().
		SetSort(bson.M{"timestamp": 1}).
		SetLimit(int64(limit))

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var readings []SensorReading
	if err := cursor.All(ctx, &readings); err != nil {
		return nil, err
	}
	return readings, nil
}

// GetHourlyAggregation returns aggregated values using MongoDB aggregation pipeline.
// This is much more efficient than aggregating in application code.
// ----------------------------------------------------------------
// GetHourlyAggregation คืนค่าที่ถูก aggregation (avg, min, max) รายชั่วโมง
// ใช้ aggregation pipeline ของ MongoDB ซึ่งมีประสิทธิภาพสูงกว่าการทำ aggregation ใน Go
func (r *sensorRepository) GetHourlyAggregation(ctx context.Context, deviceID, sensorType string, start, end time.Time) ([]HourlyAggregate, error) {
	// Aggregation pipeline:
	// Stage 1: $match - filter by device, sensor type, and time range
	// Stage 2: $group - group by hour and calculate aggregates
	// Stage 3: $project - format output
	// ----------------------------------------------------------------
	// Aggregation pipeline:
	// ขั้นที่ 1: $match - กรองตาม device, sensor type และช่วงเวลา
	// ขั้นที่ 2: $group - จัดกลุ่มรายชั่วโมงและคำนวณค่า aggregate
	// ขั้นที่ 3: $project - จัดรูปแบบ output
	pipeline := mongo.Pipeline{
		// Stage 1: Match filter
		// ขั้นที่ 1: กรองข้อมูล
		{{Key: "$match", Value: bson.M{
			"device_id":   deviceID,
			"sensor_type": sensorType,
			"timestamp": bson.M{
				"$gte": start.UTC(),
				"$lte": end.UTC(),
			},
		}}},
		// Stage 2: Group by hour
		// ขั้นที่ 2: จัดกลุ่มรายชั่วโมง
		{{Key: "$group", Value: bson.M{
			"_id": bson.M{
				"$dateTrunc": bson.M{
					"date":    "$timestamp",
					"unit":    "hour",
					"timezone": "Asia/Bangkok",
				},
			},
			"avg_value": bson.M{"$avg": "$value"},
			"min_value": bson.M{"$min": "$value"},
			"max_value": bson.M{"$max": "$value"},
			"count":     bson.M{"$sum": 1},
		}}},
		// Stage 3: Project to rename fields
		// ขั้นที่ 3: เปลี่ยนชื่อฟิลด์
		{{Key: "$project", Value: bson.M{
			"hour":      "$_id",
			"avg_value": 1,
			"min_value": 1,
			"max_value": 1,
			"count":     1,
			"_id":       0,
		}}},
		// Stage 4: Sort by hour
		// ขั้นที่ 4: เรียงตามชั่วโมง
		{{Key: "$sort", Value: bson.M{"hour": 1}}},
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []HourlyAggregate
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}