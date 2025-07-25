package models

type Image struct {
	BaseModel
	URL       string `json:"url"`
	OwnerId   int    `json:"owner_id"`
	OwnerType string `json:"owner_type"`
}
