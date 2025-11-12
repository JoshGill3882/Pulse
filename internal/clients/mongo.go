package clients

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// MongoClientInterface exposes just the entry point repositories need.
// Keeping this narrow makes it trivial to mock when writing unit tests.
type MongoClientInterface interface {
	Collection(dbName, collName string) MongoCollectionInterface
}

// MongoCollectionInterface mirrors the small subset of collection operations our repositories perform.
// Tests can provide their own implementation that records calls or returns canned data.
type MongoCollectionInterface interface {
	InsertOne(ctx context.Context, document any) (*mongo.InsertOneResult, error)
	FindOne(ctx context.Context, filter any) MongoSingleResultInterface
	Find(ctx context.Context, filter any) (MongoCursorInterface, error)
	UpdateOne(ctx context.Context, filter any, update any) (*mongo.UpdateResult, error)
	DeleteOne(ctx context.Context, filter any) (*mongo.DeleteResult, error)
}

// MongoSingleResultInterface lets repositories decode the single document returned by FindOne without binding to the concrete driver type.
type MongoSingleResultInterface interface {
	Decode(val any) error
}

// MongoCursorInterface provides the two cursor operations we rely on.
// Additional cursor methods can be added here as the codebase grows.
type MongoCursorInterface interface {
	All(ctx context.Context, results any) error
	Close(ctx context.Context) error
}

// mongoClient wraps the official driver client and satisfies MongoClientInterface by returning wrapped collections.
type mongoClient struct {
	SDK *mongo.Client
}

// NewMongo establishes the connection using environment configuration and returns the ready-to-use concrete client.
func NewMongo() (MongoClientInterface, error) {
	client, err := mongo.Connect(
		options.Client().ApplyURI(
			os.Getenv("MONGO_URI"),
		).SetServerAPIOptions(
			options.ServerAPI(options.ServerAPIVersion1),
		),
	)
	if err != nil { return nil, err }
	return &mongoClient{SDK: client}, nil
}

// Collection returns a lightweight wrapper around the requested collection.
// Repositories call this once in their constructor and keep the reference.
func (c *mongoClient) Collection(dbName, collName string) MongoCollectionInterface {
	return &mongoCollection{coll: c.SDK.Database(dbName).Collection(collName)}
}

// mongoCollection embeds the driver collection and forwards calls so that higher layers never touch the concrete driver types directly.
type mongoCollection struct {
	coll *mongo.Collection
}

func (c *mongoCollection) InsertOne(ctx context.Context, document any) (*mongo.InsertOneResult, error) {
	return c.coll.InsertOne(ctx, document)
}

func (c *mongoCollection) FindOne(ctx context.Context, filter any) MongoSingleResultInterface {
	return &mongoSingleResult{res: c.coll.FindOne(ctx, filter, options.FindOne())}
}

func (c *mongoCollection) Find(ctx context.Context, filter any) (MongoCursorInterface, error) {
	cursor, err := c.coll.Find(ctx, filter, options.Find())
	if err != nil { return nil, err }
	return &mongoCursor{cur: cursor}, nil
}

func (c *mongoCollection) UpdateOne(ctx context.Context, filter any, update any) (*mongo.UpdateResult, error) {
	return c.coll.UpdateOne(ctx, filter, update, options.UpdateOne())
}

func (c *mongoCollection) DeleteOne(ctx context.Context, filter any) (*mongo.DeleteResult, error) {
	return c.coll.DeleteOne(ctx, filter, options.DeleteOne())
}

// mongoSingleResult wraps the driver SingleResult so repositories only depend on the Decode behavior they need for tests.
type mongoSingleResult struct {
	res *mongo.SingleResult
}

func (r *mongoSingleResult) Decode(val any) error {
	return r.res.Decode(val)
}

// mongoCursor mirrors the subset of cursor operations repositories rely on.
type mongoCursor struct {
	cur *mongo.Cursor
}

func (c *mongoCursor) All(ctx context.Context, results any) error {
	return c.cur.All(ctx, results)
}

func (c *mongoCursor) Close(ctx context.Context) error {
	return c.cur.Close(ctx)
}
