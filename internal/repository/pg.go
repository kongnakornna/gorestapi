package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/kongnakornna/gorestapi/internal"
	"gorm.io/gorm"
)

type PgRepo[M any] struct {
	DB *gorm.DB
}

func CreatePgRepo[M any](db *gorm.DB) PgRepo[M] {
	return PgRepo[M]{DB: db}
}

/*************  ✨ Windsurf Command ⭐  *************/
// CreatePgRepository returns a new instance of PgRepository, which is a wrapper around GORM and implements the internal.PgRepository interface.
// It takes a GORM DB instance as a parameter and returns a pointer to a PgRepository instance.
/*******  54f8d542-0e64-467c-9067-e62f558cd312  *******/
func CreatePgRepository[M any](db *gorm.DB) internal.PgRepository[M] {
	return &PgRepo[M]{DB: db}
}

/*************  ✨ Windsurf Command ⭐  *************/
// Get retrieves the object with the given ID from the database.
// It returns the object if it exists, and an error if it does not exist.
// If an error occurs during the retrieval process, it returns the error.
//
// Parameters:
//   ctx - the context object
//   id - the ID of the object to retrieve
//
// Returns:
//   *M - the retrieved object
//   error - an error if the object does not exist or if an error occurs during the retrieval process
/*******  356680d6-f3f8-4d85-bd3b-504f65c1d016  *******/
func (r *PgRepo[M]) Get(ctx context.Context, id uuid.UUID) (*M, error) {
	var obj *M
	if result := r.DB.WithContext(ctx).First(&obj, "id = ?", id.String()); result.Error != nil {
		return nil, result.Error
	}
	return obj, nil
}

/*************  ✨ Windsurf Command ⭐  *************/
// GetMulti retrieves a list of objects from the database.
// It takes two parameters, limit and offset, which are used to limit the number of objects retrieved and to offset the starting point of the retrieval, respectively.
// The function returns a list of objects and an error. If an error occurs during the retrieval process, it returns the error.
// Otherwise, it returns the list of objects.
//
// Parameters:
//   ctx - the context object
//   limit - the maximum number of objects to retrieve
//   offset - the starting point of the retrieval
//
// Returns:
//   []*M - the list of retrieved objects
//   error - an error if an error occurs during the retrieval process
/*******  b7f393a0-81fd-48db-b5a9-2a041eff1570  *******/
func (r *PgRepo[M]) GetMulti(ctx context.Context, limit, offset int) ([]*M, error) {
	var objs []*M
	r.DB.WithContext(ctx).Limit(limit).Offset(offset).Find(&objs)
	return objs, nil
}

/*************  ✨ Windsurf Command ⭐  *************/
// Create creates a new object in the database.
// It takes two parameters, ctx and exp, where ctx is the context object and exp is the object to be created.
// If an error occurs during the creation process, it returns the error.
// Otherwise, it returns the created object.
/*******  1d6b9b0b-91ef-487c-b200-eb3aa0f3c25e  *******/
func (r *PgRepo[M]) Create(ctx context.Context, exp *M) (*M, error) {
	if result := r.DB.WithContext(ctx).Create(exp); result.Error != nil {
		return nil, result.Error
	}
	return exp, nil
}

func (r *PgRepo[M]) Delete(ctx context.Context, id uuid.UUID) (*M, error) {
	obj, err := r.Get(ctx, id)

	if err != nil {
		return nil, err
	}

	if result := r.DB.WithContext(ctx).Delete(&obj, "id = ?", id.String()); result.Error != nil {
		return nil, result.Error
	}
	return obj, nil
}

func (r *PgRepo[M]) Update(ctx context.Context, exp *M, values map[string]interface{}) (*M, error) {
	if result := r.DB.WithContext(ctx).Model(&exp).Updates(values); result.Error != nil {
		return nil, result.Error
	}
	return exp, nil
}
