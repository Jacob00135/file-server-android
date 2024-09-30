package models

type User struct {
	Id         uint   `json:"id" xml:"id" form:"id"`
	Username   string `json:"username" xml:"username" form:"username"`
	Permission uint   `json:"permission" xml:"permission" form:"permission"`
}

type UserInput struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
}
