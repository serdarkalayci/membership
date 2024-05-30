package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/serdarkalayci/membership/api/adapters/data/postgres/dao"
	"github.com/serdarkalayci/membership/api/adapters/data/postgres/mappers"
	"github.com/serdarkalayci/membership/api/domain"
)

// MemberRepository holds the arangodb client and database name for methods to use
type MemberRepository struct {
	cp     *pgxpool.Pool
	dbName string
}

func newMemberRepository(pool *pgxpool.Pool, databaseName string) MemberRepository {
	return MemberRepository{
		cp:     pool,
		dbName: databaseName,
	}
}

func (mr MemberRepository) ListMembers(pageSize, pageNum int) ([]domain.Member, int, error) {
	var listmembers []dao.ListMember
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	query := "SELECT members.id, first_name, last_name, email, phone, city_id, cities.name AS city_name from members INNER JOIN cities ON members.city_id = cities.id"
	offset := (pageNum - 1) * pageSize
	query = fmt.Sprintf("%s ORDER BY members.id LIMIT %d OFFSET %d", query, pageSize, offset)
	rows, err := mr.cp.Query(ctx, query)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	listmembers, err = pgx.CollectRows(rows, pgx.RowToStructByName[dao.ListMember])
	if err != nil {
		return nil, 0, err
	}
	var count int
	if err := mr.cp.QueryRow(ctx, fmt.Sprintf("SELECT count(ID) from Members")).Scan(&count); err != nil {
		return []domain.Member{}, 0, err
	}
	return mappers.MapListMemberdaos2Members(listmembers), count, nil
}

func (mr MemberRepository) GetMember(id string) (domain.Member, error) {
	var member dao.Member
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	row, err := mr.cp.Query(ctx, `SELECT members.id as id, first_name, last_name, passive, email, 
	phone, city_id, area_id, membership_type_id, membership_start_date, 
	last_contact_date, occupation, education, date_of_birth, COALESCE(notes, '') as notes,
	cities.name as city_name, areas.name as area_name, membership_types.name as membership_type_name
	FROM members INNER JOIN cities ON members.city_id = cities.id
	INNER JOIN provinces ON cities.province_id = provinces.id
	INNER JOIN areas ON members.area_id = areas.id
	INNER JOIN membership_types ON members.membership_type_id = membership_types.id
	WHERE members.id = $1`, id)
	if err != nil {
		return domain.Member{}, err
	}
	member, err = pgx.CollectExactlyOneRow(row, pgx.RowToStructByName[dao.Member])
	if err != nil {
		return domain.Member{}, err
	}
	return mappers.MapMemberdao2Member(member), nil
	// var member dao.Member
	// fmt.Println(mr.db.ToSQL(func(tx *gorm.DB) *gorm.DB {
	// 	return tx.Joins("City").Joins("JOIN provinces on cities.province_id = provinces.id").Joins("Area").Joins("MembershipType").First(&member, "id = ?", id)
	// }))
	
	// if err := mr.db.First(&member, "ID = ?", id).Preload("City").Joins("Area").Joins("MembershipType").Error; err != nil {
	// 	return domain.Member{}, err
	// }
	// return mappers.MapMemberdao2Member(member), nil
}