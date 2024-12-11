package model

type User struct {
	username string
	password string
}

func (User) TableName() string {
	return "user"
}
