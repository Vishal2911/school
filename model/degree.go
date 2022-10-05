package model

type Degree struct {
	Primary       string `json:"primary"`
	Secondary     string `json:"secondary"`
	UnderGraduate string `json:"under_graduate"`
	PostGraduate  string `json:"post_graduate"`
	PHD           string `json:"phd"`
	B_ad          string `json:"b_ad"`
}
