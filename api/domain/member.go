package domain

import "time"

// Member is the struct that holds the member information
type Member struct {
	ID       			string
	FirstName			string
	LastName			string
	Email    			string
	Phone				string
	City				City
	Area				Area
	MembershipType		MembershipType
	MembershipStartDate	time.Time
	LastContactDate		time.Time
	Occupation			string
	Education			string
	DateOfBirth			time.Time
	Notes				string
	Passive				bool
}

