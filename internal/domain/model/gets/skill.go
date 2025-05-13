package model

type Skill struct {
	Name   string `json:"name"`
	SvgURL string `json:"svgUrl"`
	Order  uint   `json:"order"`
}

type SkillCategory struct {
	ID           string  `json:"id"`
	CategoryName string  `json:"categoryName"`
	Order        uint    `json:"order"`
	Skills       []Skill `json:"skills"`
}
