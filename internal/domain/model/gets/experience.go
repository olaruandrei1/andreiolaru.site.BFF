package model

type Experience struct {
	Company     string `json:"company"`
	Title       string `json:"title"`
	Period      string `json:"period"`
	Description string `json:"description"`
	PhotoPath   string `json:"photoPath"`
}
