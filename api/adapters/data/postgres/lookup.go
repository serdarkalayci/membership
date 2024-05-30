package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/serdarkalayci/membership/api/adapters/data/postgres/dao"
	"github.com/serdarkalayci/membership/api/adapters/data/postgres/mappers"
	"github.com/serdarkalayci/membership/api/domain"
)

// LookupRepository holds the gorm db for methods to use
type LookupRepository struct {
	cp     *pgxpool.Pool
	dbName string
}

func newLookupRepository(pool *pgxpool.Pool, databaseName string) LookupRepository {
	return LookupRepository{
		cp:     pool,
		dbName: databaseName,
	}
}

func (lr LookupRepository) GetCities() ([]domain.City, error) {
	var cities []dao.City
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rows, err := lr.cp.Query(ctx, "SELECT id, name FROM cities")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	cities, err = pgx.CollectRows(rows, pgx.RowToStructByName[dao.City])
	if err != nil {
		return nil, err
	}
	return mappers.MapCitydaos2Cities(cities), nil
}