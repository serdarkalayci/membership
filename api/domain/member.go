package domain

import "time"

// Member is the struct that holds the member information
type Member struct {
	ID       			string
	Email    			string
	FirstName			string
	LastName			string
	City				City
	Area				Area
	Phone				string
	Notes				string
	MembershipType		MembershipType
	MembershipStartDate	time.Time
	LastContactDate		time.Time
	Occupation			string
	Education			string
	DateOfBirth			time.Time
	Passive				bool
}

// City is the struct that holds the city information
type City struct {
	ID   string
	Name string
}

// Province is the struct that holds the province information
type Province struct {
	ID   string
	Name string
}

// Area is the struct that holds the area information
type Area struct {
	ID   string
	Name string
}

// MembershipType is the struct that holds the membership type information
type MembershipType struct {
	ID   string
	Name string
}