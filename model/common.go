package model

type Name struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type Address struct {
	Line1  string `json:"line1"`
	Line2  string `json:"line2"`
	Nearby string `json:"nearBy"`
}

type Number struct {
	Primary   int `json:"primary"`
	Secondary int `json:"secondary"`
	Landline  int `json:"landline"`
}

type Parents struct {
	FathersName          string `json:"fatherName"`
	MotherName           string `json:"motherName"`
	GaurdianName         string `json:"gaurdianName"`
	GaurdianRealtionship string `json:"gaurdianRelationship`
}
