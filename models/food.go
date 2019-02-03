package models

type Food struct {
	Id string `json:"_id"        bson:"_id,omitempty"`
	Description string `bson:"description" json:"description"`
	AddedDate string `bson:"added_date" json:"added_date"`
	LimitDate string `bson:"limit_date" json:"limit_date"`
	UserUid string `bson:"user_uid" json:"user_uid"`
}