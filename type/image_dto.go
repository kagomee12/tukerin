package types

type ImageDTO struct {
	ID   int    `json:"id" form:"id"`
	URL  string `json:"url" form:"url"`
}