package dao

// Province is a data access object for provinces
type Province struct {
	ID int `db:"id"`
	Name string  `db:"name"`
  }
  
  // City is a data access object for cities
  type City struct {
	ID int  `db:"id"`
	ProvinceID int `db:"province_id"`
	ProvinceName string `db:"province_name"`
	Name string 
  }
  
  // Area is a data access object for areas
  type Area struct {
	ID int  `db:"id"`
	Name string  `db:"name"`
  }
  
  // MembershipType is a data access object for membership types
  type MembershipType struct {
	ID int  `db:"id"`
	Name string  `db:"name"`
  }
  