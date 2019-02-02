package models

type User struct {
	AuthUid string `bson:"auth_uid" json:"auth_uid"`
	Token string `bson:"token" json:"token"`
}
