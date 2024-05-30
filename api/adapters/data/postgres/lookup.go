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

func (lr LookupRepository) GetAreas() ([]domain.Area, error) {
	var areas []dao.Area
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rows, err := lr.cp.Query(ctx, "SELECT id, name FROM areas")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	areas, err = pgx.CollectRows(rows, pgx.RowToStructByName[dao.Area])
	if err != nil {
		return nil, err
	}
	return mappers.MapAreadaos2Areas(areas), nil
}

func (lr LookupRepository) GetProvinces() ([]domain.Province, error) {
	var provinces []dao.Province
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rows, err := lr.cp.Query(ctx, "SELECT id, name FROM areas")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	provinces, err = pgx.CollectRows(rows, pgx.RowToStructByName[dao.Province])
	if err != nil {
		return nil, err
	}
	return mappers.MapProvincedaos2Provinces(provinces), nil
}

func (lr LookupRepository) GetMembershipTypes() ([]domain.MembershipType, error) {
	var membershipTypes []dao.MembershipType
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rows, err := lr.cp.Query(ctx, "SELECT id, name FROM areas")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	membershipTypes, err = pgx.CollectRows(rows, pgx.RowToStructByName[dao.MembershipType])
	if err != nil {
		return nil, err
	}
	return mappers.MapMembershipTypedaos2MembershipTypes(membershipTypes), nil
}