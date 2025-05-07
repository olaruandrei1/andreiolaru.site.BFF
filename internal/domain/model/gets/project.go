package model

type ProjectTechnology struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type Project struct {
	Title        string              `json:"title"`
	Technologies []ProjectTechnology `json:"technologies"`
	RepoURL      string              `json:"repoUrl"`
}
