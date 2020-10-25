package model

func TestUser() *User {
	return &User{
		Email:    "test@mail.ru",
		Password: "password",
	}
}
