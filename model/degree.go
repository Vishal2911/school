package model

type Degree struct {
	DegreeName string   `json:"degree_name"`
	DegreeType Duration `json:"degree_type"`
	Duration   string   `json:"duration"`
	Percentage string   `json:"percentage"`
	CGPA       string   `json:"cgpa"`
}
