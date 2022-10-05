package model

type Name struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Address struct {
	Line1  string `json:"line1"`
	Line2  string `json:"line2"`
	Nearby string `json:"near_by"`
}

type Number struct {
	Primary   int `json:"primary"`
	Secondary int `json:"secondary"`
	Landline  int `json:"landline"`
}

type Parents struct {
	FatherName          string `json:"father_name"`
	MotherName           string `json:"mother_name"`
	GaurdianName         string `json:"gaurdian_name"`
	GaurdianRealtionship string `json:"gaurdian_relationship`
}
