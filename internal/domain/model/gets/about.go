package model

type AboutBlock struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type About struct {
	WhoIAm  AboutBlock `json:"whoIAm"`
	Mindset AboutBlock `json:"mindset"`
	CVURL   string     `json:"cvDownloadUrl"`
}
