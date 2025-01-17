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
	*MongoRepository
}

type MongoRepository struct {
	Database   string
	Collection string
	MgoColl    *mongo.Collection
}

func (repository *MongoRepository) InsertOne(model any) (string, error) {
	return repository.InsertOneContext(context.Background(), model)
}

func (repository *MongoRepository) InsertMany(models []any) error {
	return repository.InsertManyContext(context.Background(), models)
}

func (repository *MongoRepository) CountDocuments(filter any, opts ...*options.CountOptions) (int64, error) {
	return repository.CountDocumentsContext(context.Background(), filter, opts...)
}

func (repository *MongoRepository) Find(results any, filter any, opts ...*options.FindOptions) error {
	return repository.FindContext(context.Background(), results, filter, opts...)
}

func (repository *MongoRepository) FindAndSort(results any, filter any, direction int, sorts ...string) error {
	return repository.FindAndSortContext(context.Background(), results, filter, direction, sorts...)
}

func (repository *MongoRepository) FindAndPage(results any, filter any, skip int64, limit int64) error {
	return repository.FindAndPageContext(context.Background(), results, filter, skip, limit)
}

func (repository *MongoRepository) FindAndPageAndSort(results any, filter any, skip int64, limit int64, direction int, sorts ...string) error {
	return repository.FindAndPageAndSortContext(context.Background(), results, filter, skip, limit, direction, sorts...)
}

func (repository *MongoRepository) FindByIds(results any, ids []string) error {
	return repository.FindByIdsContext(context.Background(), results, ids)
}

func (repository *MongoRepository) Aggregate(results any, pipeline any) error {
	return repository.AggregateContext(context.Background(), results, pipeline)
}

func (repository *MongoRepository) InsertOrUpdateOne(filter any, update any) (int64, error) {
	return repository.InsertOrUpdateOneContext(context.Background(), filter, update)
}

func (repository *MongoRepository) UpdateOne(filter any, update any) (int64, error) {
	return repository.UpdateOneContext(context.Background(), filter, update)
}

func (repository *MongoRepository) UpdateByPrimitiveId(id primitive.ObjectID, update any) (int64, error) {
	return repository.UpdateByPrimitiveIdContext(context.Background(), id, update)
}

func (repository *MongoRepository) UpdateByStringId(id string, update any) (int64, error) {
	return repository.UpdateByStringIdContext(context.Background(), id, update)
}

func (repository *MongoRepository) UpdateMany(filter any, update any) (int64, error) {
	return repository.UpdateManyContext(context.Background(), filter, update)
}

func (repository *MongoRepository) DeleteOne(filter any) (int64, error) {
	return repository.DeleteOneContext(context.Background(), filter)
}

func (repository *MongoRepository) DeleteMany(filter any) (int64, error) {
	return repository.DeleteManyContext(context.Background(), filter)
}

func (repository *MongoRepository) InsertOneContext(ctx context.Context, model any) (string, error) {
	one, err := repository.MgoColl.InsertOne(ctx, model)
	if err != nil {
		marsLog.Logger().ErrorF("InsertOne error: %v", err)
		return "", err
	}
	id := one.InsertedID.(primitive.ObjectID)
	return id.Hex(), nil
}

func (repository *MongoRepository) InsertManyContext(ctx context.Context, models []any) error {
	_, err := repository.MgoColl.InsertMany(ctx, models)
	if err != nil {
		marsLog.Logger().ErrorF("InsertMany error: %v", err)
		return err
	}
	return nil
}

func (repository *MongoRepository) CountDocumentsContext(ctx context.Context, filter any, opts ...*options.CountOptions) (int64, error) {
	documents, err := repository.MgoColl.CountDocuments(ctx, filter, opts...)
	if err != nil {
		marsLog.Logger().ErrorF("CountDocuments error: %v", err)
		return 0, err
	}
	return documents, nil
}

func (repository *MongoRepository) FindContext(ctx context.Context, results any, filter any, opts ...*options.FindOptions) error {
	cursor, err := repository.MgoColl.Find(ctx, filter, opts...)
	if err != nil {
		marsLog.Logger().ErrorF("Find error: %v", err)
		return err
	}
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

func (repository *MongoRepository) FindAndSortContext(ctx context.Context, results any, filter any, direction int, sorts ...string) error {
	opts := &options.FindOptions{}
	setSortOpts(opts, direction, sorts...)
	return repository.FindContext(ctx, results, filter, opts)
}

func (repository *MongoRepository) FindAndPageContext(ctx context.Context, results any, filter any, skip int64, limit int64) error {
	opts := &options.FindOptions{}
	setPageOpts(opts, &skip, &limit)
	return repository.FindContext(ctx, results, filter, opts)
}

func (repository *MongoRepository) FindAndPageAndSortContext(ctx context.Context, results any, filter any, skip int64, limit int64, direction int, sorts ...string) error {
	if limit < 1 || skip < 0 {
		return errors.New("invalid skip or limit value")
	}
	opts := &options.FindOptions{}
	setSortOpts(opts, direction, sorts...)
	setPageOpts(opts, &skip, &limit)
	return repository.FindContext(ctx, results, filter, opts)
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

func (repository *MongoRepository) FindByIdsContext(ctx context.Context, results any, ids []string) error {
	if len(ids) == 0 {
		return nil
	}
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
	err := repository.FindAndPageContext(ctx, results, filter, 0, int64(len(ids)))
	if err != nil {
		marsLog.Logger().ErrorF("FindAndPage error: %v", err)
		return err
	}
	return nil
}

func (repository *MongoRepository) AggregateContext(ctx context.Context, results any, pipeline any) error {
	cursor, err := repository.MgoColl.Aggregate(ctx, pipeline)
	if err != nil {
		marsLog.Logger().ErrorF("Aggregate error: %v", err)
		return err
	}
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

func (repository *MongoRepository) InsertOrUpdateOneContext(ctx context.Context, filter any, update any) (int64, error) {
	set := bson.M{"$set": update}
	one, err := repository.MgoColl.UpdateOne(ctx, filter, set, options.Update().SetUpsert(true))
	if err != nil {
		marsLog.Logger().ErrorF("InsertOrUpdateOne error: %v", err)
		return 0, err
	}
	return one.MatchedCount, nil
}

func (repository *MongoRepository) UpdateOneContext(ctx context.Context, filter any, update any) (int64, error) {
	set := bson.M{"$set": update}
	one, err := repository.MgoColl.UpdateOne(ctx, filter, set)
	if err != nil {
		marsLog.Logger().ErrorF("UpdateOne error: %v", err)
		return 0, err
	}
	return one.MatchedCount, nil
}

func (repository *MongoRepository) UpdateByPrimitiveIdContext(ctx context.Context, id primitive.ObjectID, update any) (int64, error) {
	filter := bson.M{"_id": id}
	return repository.UpdateOneContext(ctx, filter, update)
}

func (repository *MongoRepository) UpdateByStringIdContext(ctx context.Context, id string, update any) (int64, error) {
	hex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		marsLog.Logger().ErrorF("ObjectIDFromHex error: %v", err)
		return 0, err
	}
	filter := bson.M{"_id": hex}
	return repository.UpdateOneContext(ctx, filter, update)
}

func (repository *MongoRepository) UpdateManyContext(ctx context.Context, filter any, update any) (int64, error) {
	many, err := repository.MgoColl.UpdateMany(ctx, filter, update)
	if err != nil {
		marsLog.Logger().ErrorF("UpdateMany error: %v", err)
		return 0, err
	}
	return many.MatchedCount, nil
}

func (repository *MongoRepository) DeleteOneContext(ctx context.Context, filter any) (int64, error) {
	one, err := repository.MgoColl.DeleteOne(ctx, filter)
	if err != nil {
		marsLog.Logger().ErrorF("DeleteOne error: %v", err)
		return 0, err
	}
	return one.DeletedCount, nil
}

func (repository *MongoRepository) DeleteManyContext(ctx context.Context, filter any) (int64, error) {
	many, err := repository.MgoColl.DeleteMany(ctx, filter)
	if err != nil {
		marsLog.Logger().ErrorF("DeleteMany error: %v", err)
		return 0, err
	}
	return many.DeletedCount, nil
}
