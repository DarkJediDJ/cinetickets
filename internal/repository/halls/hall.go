package hall

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"go.uber.org/zap"

	"github.com/darkjedidj/cinema-service/internal"
)

// Repository is a struct to store storage and logger connection
type Repository struct {
	DB  *sql.DB
	Log *zap.Logger
}

// Resource is a struct to store data about entity
type Resource struct {
	ID    int64 `json:"ID"`
	VIP   bool  `json:"VIP"`
	Seats int   `json:"seats"`
}

func (r *Resource) GID() int64 {
	return r.ID
}

// Create new entity in storage
func (r *Repository) Create(i internal.Identifiable, ctx context.Context) (internal.Identifiable, error) {
	var id int64

	hall, ok := i.(*Resource)
	if !ok {
		r.Log.Info("Failed to create Hall object.",
			zap.Bool("ok", ok),
		)

		return nil, internal.ErrInternalFailure
	}

	err := sq.
		Insert("halls").
		Columns("vip", "seats").
		Values(hall.VIP, hall.Seats).
		Suffix("RETURNING \"id\"").
		PlaceholderFormat(sq.Dollar).
		RunWith(r.DB).
		QueryRowContext(ctx).
		Scan(&id)

	if err != nil {
		r.Log.Info("Failed to run Create hall query.",
			zap.Error(err),
		)

		return nil, internal.ErrInternalFailure
	}

	return r.Retrieve(id, ctx)
}

// Retrieve entity from storage
func (r *Repository) Retrieve(id int64, ctx context.Context) (internal.Identifiable, error) {
	var res Resource

	err := sq.
		Select("vip", "id", "seats").
		From("halls").
		Where(sq.Eq{
			"id": id,
		}).
		PlaceholderFormat(sq.Dollar).
		RunWith(r.DB).
		QueryRowContext(ctx).
		Scan(&res.VIP, &res.ID, &res.Seats)

	if err == sql.ErrNoRows {

		return nil, nil
	}

	if err != nil {
		r.Log.Info("Failed to run Retrieve hall query.",
			zap.Error(err),
		)

		return nil, internal.ErrInternalFailure
	}

	return &res, nil
}

// Delete entity in storage
func (r *Repository) Delete(id int64, ctx context.Context) error {

	_, err := sq.
		Delete("halls").
		Where(sq.Eq{
			"id": id,
		}).
		PlaceholderFormat(sq.Dollar).
		RunWith(r.DB).
		ExecContext(ctx)

	if err != nil {
		r.Log.Info("Failed to run Delete hall query.",
			zap.Error(err),
		)

		return internal.ErrInternalFailure
	}

	return nil
}

// RetrieveAll entity from storage
func (r *Repository) RetrieveAll(ctx context.Context) ([]internal.Identifiable, error) {

	rows, err := sq.
		Select("vip", "id", "seats").
		From("halls").
		PlaceholderFormat(sq.Dollar).
		RunWith(r.DB).QueryContext(ctx)

	if err != nil {
		r.Log.Info("Failed to run RetrieveAll halls query.",
			zap.Error(err),
		)

		return nil, internal.ErrInternalFailure
	}

	var data []*Resource

	for rows.Next() {
		res := &Resource{}

		err = rows.Scan(&res.VIP, &res.ID, &res.Seats)
		if err == sql.ErrNoRows {
			return nil, nil
		}

		if err != nil {
			r.Log.Info("Failed to scan rows into halls structures.",
				zap.Error(err),
			)

			return nil, internal.ErrInternalFailure
		}

		data = append(data, res)
	}

	var dataSlice []*Resource = data
	var interfaceSlice []internal.Identifiable = make([]internal.Identifiable, len(dataSlice))
	for i, d := range dataSlice {
		interfaceSlice[i] = d
	}

	return interfaceSlice, nil
}
