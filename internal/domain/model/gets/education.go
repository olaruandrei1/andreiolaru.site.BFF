package model

type EducationVariant string

const (
	Primary   EducationVariant = "PRIMARY"
	Secondary EducationVariant = "SECONDARY"
	Classic   EducationVariant = "CLASSIC"
)

type Education struct {
	Institution string           `json:"institution"`
	Degree      string           `json:"degree"`
	Period      string           `json:"period"`
	Description string           `json:"description"`
	Variant     EducationVariant `json:"variant"`
}
