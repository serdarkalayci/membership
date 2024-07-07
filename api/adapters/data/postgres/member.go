package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
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

func (mr MemberRepository) ListMembers(pageSize, pageNum int, searchName string, searchCity int, searchArea int) ([]domain.Member, int, error) {
	var listmembers []dao.ListMember
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	query := "SELECT members.id, first_name, last_name, email, phone, city_id, cities.name AS city_name from members INNER JOIN cities ON members.city_id = cities.id"
	query = fmt.Sprintf("%s WHERE (first_name ILIKE '%%%s%%' OR last_name ILIKE '%%%s%%')", query, searchName, searchName)
	query = fmt.Sprintf("%s AND (city_id = %d OR %d = 0)", query, searchCity, searchCity)
	query = fmt.Sprintf("%s AND (area_id = %d OR %d = 0)", query, searchArea, searchArea)
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
	query = "SELECT count(id) as count from members"
	query = fmt.Sprintf("%s WHERE (first_name ILIKE '%%%s%%' OR last_name ILIKE '%%%s%%')", query, searchName, searchName)
	query = fmt.Sprintf("%s AND (city_id = %d OR %d = 0)", query, searchCity, searchCity)
	query = fmt.Sprintf("%s AND (area_id = %d OR %d = 0)", query, searchArea, searchArea)
	if err := mr.cp.QueryRow(ctx, query).Scan(&count); err != nil {
		return []domain.Member{}, 0, err
	}
	return mappers.MapListMemberdaos2Members(listmembers), count, nil
}

func (mr MemberRepository) GetMember(id string) (domain.Member, error) {
	var member dao.Member
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	row, err := mr.cp.Query(ctx, `SELECT members.id as id, first_name, last_name, passive, email, 
	phone, provinces.id as province_id, city_id, area_id, membership_type_id, membership_start_date, 
	last_contact_date, occupation, education, date_of_birth, COALESCE(notes, '') as notes,
	provinces.name as province_name, cities.name as city_name, 
	areas.name as area_name, membership_types.name as membership_type_name
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
}

func (mr MemberRepository) UpdateMember(member domain.Member) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := mr.cp.Exec(ctx, `UPDATE members SET first_name = $1, last_name = $2, email = $3, phone = $4, 
	city_id = $5, area_id = $6, membership_type_id = $7, membership_start_date = $8, last_contact_date = $9, 
	occupation = $10, education = $11, date_of_birth = $12, notes = $13 WHERE id = $14`,
		member.FirstName, member.LastName, member.Email, member.Phone, member.City.ID, member.Area.ID, member.MembershipType.ID,
		member.MembershipStartDate, member.LastContactDate, member.Occupation, member.Education, member.DateOfBirth, member.Notes, member.ID)
	if err != nil {
		return err
	}
	return nil
}

func (mr MemberRepository) CreateMember(member domain.Member) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var id string
	id = uuid.New().String()
	err := mr.cp.QueryRow(ctx, `INSERT INTO members (id, first_name, last_name, email, phone, city_id, area_id, membership_type_id, 
	membership_start_date, last_contact_date, occupation, education, date_of_birth, notes) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING id`,
		id, member.FirstName, member.LastName, member.Email, member.Phone, member.City.ID, member.Area.ID, member.MembershipType.ID,
		member.MembershipStartDate, member.LastContactDate, member.Occupation, member.Education, member.DateOfBirth, member.Notes).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}