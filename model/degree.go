package model

type Degree struct {
	Primary       string `json:"primary"`
	Secondary     string `json:"secondary"`
	UnderGraduate string `json:"under_graduate"`
	PostGraduate  string `json:"postGraduate"`
	PHD           string `json:"phd"`
	Bed           string `json:"bed"`
}
