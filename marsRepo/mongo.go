package marsRepo

import (
	"context"
	"errors"
	"github.com/marsli9945/mars-go/marsLog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BaseRepo struct {
	Repository *MongoRepository
}

type MongoRepository struct {
	Database   string
	Collection string
	MgoColl    *mongo.Collection
}

func (repository *MongoRepository) InsetOne(model interface{}) (string, error) {
	one, err := repository.MgoColl.InsertOne(context.Background(), model)
	if err != nil {
		marsLog.Logger().ErrorF("InsertOne error: %v", err)
		return "", err
	}
	id := one.InsertedID.(primitive.ObjectID)
	return id.Hex(), nil
}

func (repository *MongoRepository) InsertMany(models []interface{}) error {
	_, err := repository.MgoColl.InsertMany(context.Background(), models)
	if err != nil {
		marsLog.Logger().ErrorF("InsertMany error: %v", err)
		return err
	}
	return nil
}

func (repository *MongoRepository) CountDocuments(filter interface{}, opts ...*options.CountOptions) (int64, error) {
	documents, err := repository.MgoColl.CountDocuments(context.Background(), filter, opts...)
	if err != nil {
		marsLog.Logger().ErrorF("CountDocuments error: %v", err)
		return 0, err
	}
	return documents, nil
}

func (repository *MongoRepository) Find(results interface{}, filter interface{}, opts ...*options.FindOptions) error {
	ctx := context.Background()
	cursor, err := repository.MgoColl.Find(ctx, filter, opts...)
	if err != nil {
		marsLog.Logger().ErrorF("Find error: %v", err)
		return err
	}
	//延迟关闭游标
	defer func() {
		if err = cursor.Close(ctx); err != nil {
			marsLog.Logger().ErrorF("Close cursor error: %v", err)
		}
	}()
	if err = cursor.All(ctx, results); err != nil {
		return err
	}
	return nil
}

func (repository *MongoRepository) FindAndSort(results interface{}, filter interface{}, direction int, sorts ...string) error {
	opts := &options.FindOptions{}
	setSortOpts(opts, direction, sorts...)
	return repository.Find(results, filter, opts)
}

func (repository *MongoRepository) FindAndPage(results interface{}, filter interface{}, skip int64, limit int64) error {
	opts := &options.FindOptions{}
	setPageOpts(opts, &skip, &limit)
	return repository.Find(results, filter, opts)
}

func (repository *MongoRepository) FindAndPageAndSort(results interface{}, filter interface{}, skip int64, limit int64, direction int, sorts ...string) error {
	if limit < 1 || skip < 0 {
		return errors.New("invalid skip or limit value")
	}
	opts := &options.FindOptions{}
	setSortOpts(opts, direction, sorts...)
	setPageOpts(opts, &skip, &limit)
	return repository.Find(results, filter, opts)
}

func setSortOpts(findOptions *options.FindOptions, direction int, sorts ...string) {
	d := bson.D{}
	for _, sort := range sorts {
		d = append(d, bson.E{Key: sort, Value: direction})
	}
	findOptions.Sort = d
}

func setPageOpts(findOptions *options.FindOptions, skip *int64, limit *int64) {
	findOptions.Skip = skip
	findOptions.Limit = limit
}

func (repository *MongoRepository) FindByIds(results interface{}, ids []string) error {
	idList := make([]primitive.ObjectID, len(ids))
	for i, id := range ids {
		hex, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			marsLog.Logger().ErrorF("ObjectIDFromHex error: %v", err)
			return err
		}
		idList[i] = hex
	}
	filter := bson.M{"_id": bson.M{"$in": idList}}
	err := repository.FindAndPage(results, filter, 0, int64(len(ids)))
	if err != nil {
		marsLog.Logger().ErrorF("FindAndPage error: %v", err)
		return err
	}
	return nil
}

func (repository *MongoRepository) Aggregate(results interface{}, pipeline interface{}) error {
	ctx := context.Background()
	cursor, err := repository.MgoColl.Aggregate(ctx, pipeline)
	if err != nil {
		marsLog.Logger().ErrorF("Aggregate error: %v", err)
		return err
	}
	//延迟关闭游标
	defer func() {
		if err = cursor.Close(ctx); err != nil {
			marsLog.Logger().ErrorF("Close cursor error: %v", err)
		}
	}()
	if err = cursor.All(ctx, results); err != nil {
		return err
	}
	return nil
}

func (repository *MongoRepository) AggregateByMatch(results interface{}, matchStage bson.D, groupStage bson.D) error {
	return repository.Aggregate(results, mongo.Pipeline{bson.D{
		{"$match", matchStage},
	}, groupStage})
}

func (repository *MongoRepository) InsertOrUpdateOne(filter interface{}, update interface{}) (int64, error) {
	set := bson.M{"$set": update}
	one, err := repository.MgoColl.UpdateOne(context.Background(), filter, set, options.Update().SetUpsert(true))
	if err != nil {
		marsLog.Logger().ErrorF("InsertOrUpdateOne error: %v", err)
		return 0, err
	}
	return one.MatchedCount, nil
}

func (repository *MongoRepository) UpdateOne(filter interface{}, update interface{}) (int64, error) {
	set := bson.M{"$set": update}
	one, err := repository.MgoColl.UpdateOne(context.Background(), filter, set)
	if err != nil {
		marsLog.Logger().ErrorF("UpdateOne error: %v", err)
		return 0, err
	}
	return one.MatchedCount, nil
}

func (repository *MongoRepository) UpdateByPrimitiveId(id primitive.ObjectID, update interface{}) (int64, error) {
	filter := bson.M{"_id": id}
	return repository.UpdateOne(filter, update)
}

func (repository *MongoRepository) UpdateByStringId(id string, update interface{}) (int64, error) {
	hex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		marsLog.Logger().ErrorF("ObjectIDFromHex error: %v", err)
		return 0, err
	}
	filter := bson.M{"_id": hex}
	return repository.UpdateOne(filter, update)
}

func (repository *MongoRepository) UpdateMany(filter interface{}, update interface{}) (int64, error) {
	many, err := repository.MgoColl.UpdateMany(context.Background(), filter, update)
	if err != nil {
		marsLog.Logger().ErrorF("UpdateMany error: %v", err)
		return 0, err
	}
	return many.MatchedCount, nil
}

func (repository *MongoRepository) DeleteOne(filter interface{}) (int64, error) {
	one, err := repository.MgoColl.DeleteOne(context.Background(), filter)
	if err != nil {
		marsLog.Logger().ErrorF("DeleteOne error: %v", err)
		return 0, err
	}
	return one.DeletedCount, nil
}

func (repository *MongoRepository) DeleteMany(filter interface{}) (int64, error) {
	many, err := repository.MgoColl.DeleteMany(context.Background(), filter)
	if err != nil {
		marsLog.Logger().ErrorF("DeleteMany error: %v", err)
		return 0, err
	}
	return many.DeletedCount, nil
}
