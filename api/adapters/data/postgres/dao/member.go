package dao

import (
	"time"

	"github.com/google/uuid"
)

// Member is a data access object for members
type Member struct {
  ID                   uuid.UUID     `db:"id"`
  Firstname            string        `db:"first_name"`
  Lastname             string        `db:"last_name"`
  Passive              bool          `db:"passive"`
  Email                string        `db:"email"`
  Phone                string        `db:"phone"`
  CityID               int           `db:"city_id"`
  CityName             string        `db:"city_name"`
  AreaID               int           `db:"area_id"`
  AreaName             string        `db:"area_name"`
  Notes                string        `db:"notes"`
  MembershipTypeID     int           `db:"membership_type_id"`
  MembershipTypeName   string        `db:"membership_type_name"`
  MembershipStartDate  time.Time     `db:"membership_start_date"`
  LastContactDate      time.Time     `db:"last_contact_date"`
  Occupation           string        `db:"occupation"`
  Education            string        `db:"education"`
  DateOfBirth          time.Time     `db:"date_of_birth"`
}

// ListMember is a data acess object for members for listing ppurposes with less fields
type ListMember struct {
  ID        uuid.UUID `db:"id"`
  Firstname string    `db:"first_name"`
  Lastname  string    `db:"last_name"`
  Email     string    `db:"email"`
  Phone     string    `db:"phone"`
  CityID    int       `db:"city_id"`
  CityName  string    `db:"city_name"`
}

