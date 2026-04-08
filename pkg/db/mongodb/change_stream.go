
package mongodb

import (
	"context"
	"encoding/json"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ChangeEvent represents a MongoDB change stream event.
// ----------------------------------------------------------------
// ChangeEvent แทน event ของ MongoDB change stream
type ChangeEvent struct {
	OperationType string                 `bson:"operationType"`
	FullDocument  map[string]interface{} `bson:"fullDocument"`
	DocumentKey   map[string]interface{} `bson:"documentKey"`
	NS            struct {
		DB         string `bson:"db"`
		Coll       string `bson:"coll"`
	} `bson:"ns"`
}

// ChangeStreamHandler is called for each change event.
// ----------------------------------------------------------------
// ChangeStreamHandler ถูกเรียกสำหรับแต่ละ change event
type ChangeStreamHandler func(event ChangeEvent)

// WatchCollection watches a collection for changes and triggers handler.
// ----------------------------------------------------------------
// WatchCollection ติดตามการเปลี่ยนแปลงใน collection และเรียก handler
func WatchCollection(ctx context.Context, collection *mongo.Collection, handler ChangeStreamHandler) error {
	// Open change stream
	// เปิด change stream
	changeStream, err := collection.Watch(ctx, mongo.Pipeline{}, options.ChangeStream().SetFullDocument(options.UpdateLookup))
	if err != nil {
		return err
	}
	defer changeStream.Close(ctx)

	log.Printf("Started change stream on collection %s", collection.Name())

	for changeStream.Next(ctx) {
		var event ChangeEvent
		if err := changeStream.Decode(&event); err != nil {
			log.Printf("Failed to decode change event: %v", err)
			continue
		}
		// Trigger handler asynchronously
		// เรียก handler แบบ asynchronous
		go handler(event)
	}

	return changeStream.Err()
}