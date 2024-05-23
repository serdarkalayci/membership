package dto

import (
	"time"

	"github.com/google/uuid"
)

// Member is a data transfer object for members
type Member struct {
  ID        uuid.UUID `gorm:"type:uuId;primaryKey"`
  Firstname string    `gorm:"type:varchar(255);not null"`
  Lastname  string    `gorm:"type:varchar(255);not null"`
  Passive   bool      `gorm:"type:boolean;not null;default:false"`
  Email     string    `gorm:"type:varchar(255);unique;not null"`
  Phone     string    `gorm:"type:varchar(255);unique;not null"`
  CityID   int       `gorm:"type:int;not null"`
  City      City     `gorm:"foreignKey:city_id;references:id"`
  AreaID   int       `gorm:"type:int;not null"`
  Area      Area     `gorm:"foreignKey:area_id;references:id"`
  Notes     string    `gorm:"type:text"`
  MembershipTypeID int `gorm:"type:int;not null"`
  MembershipType MembershipType `gorm:"foreignKey:membership_type_id;references:id"`
  MembershipStartDate time.Time `gorm:"type:date"`
  LastContactDate time.Time `gorm:"type:date"`
  Occupation string `gorm:"type:varchar(255)"`
  Education  string `gorm:"type:varchar(255)"`
  DateOfBirth time.Time `gorm:"type:date"`
}

type ListMember struct {
  ID        uuid.UUID `gorm:"type:uuId;primaryKey"`
  Firstname string    `gorm:"type:varchar(255);not null"`
  Lastname  string    `gorm:"type:varchar(255);not null"`
  Email     string    `gorm:"type:varchar(255);unique;not null"`
  Phone     string    `gorm:"type:varchar(255);unique;not null"`
  CityID  int       `gorm:"type:int;not null"`
  City      City     `gorm:"foreignKey:city_id;references:id"`
}

func (m *Member) TableName() string {
  return "members"
}

// Province is a data transfer object for provinces
type Province struct {
  ID int `gorm:"type:int;primaryKey"`
  Name string `gorm:"type:varchar(255);not null"`
}

func (p *Province) TableName() string {
  return "provinces"
}

// City is a data transfer object for cities
type City struct {
  ID int `gorm:"type:int;primaryKey"`
  ProvinceID int `gorm:"type:int;not null"`
  Province Province `gorm:"foreignKey:province_id;references:id"`
  Name string `gorm:"type:varchar(255);not null"`
}

func (c *City) TableName() string {
  return "cities"
}

// Area is a data transfer object for areas
type Area struct {
  ID int `gorm:"type:int;primaryKey"`
  Name string `gorm:"type:varchar(255);not null"`
}

func (a *Area) TableName() string {
  return "areas"
}

// MembershipType is a data transfer object for membership types
type MembershipType struct {
  ID int `gorm:"type:int;primaryKey"`
  Name string `gorm:"type:varchar(255);not null"`
}

func (mt *MembershipType) TableName() string {
  return "membershiptypes"
}
