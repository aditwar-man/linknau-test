package models

import "github.com/dgrijalva/jwt-go"

var JwtKey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type Data struct {
	ID        int    `json:"id"`
	UserId    int    `json:"user_id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type Person struct {
	Name string
	Age  int
}

type Speaker interface {
	Speak() string
}

func (p Person) Speak() string {
	return "Hello, my name is " + p.Name
}
