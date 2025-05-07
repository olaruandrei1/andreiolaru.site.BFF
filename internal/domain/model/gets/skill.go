package model

type SkillCategory struct {
	ID           string            `json:"id"`
	CategoryName string            `json:"categoryName"`
	Skills       map[string]string `json:"skills"` // name: iconUrl
}
