package main

import (
	"fmt"
	"math/rand/v2"
)

type account struct {
	login    string
	password string
	url      string
}

var letterRunes = []rune("bcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*1234567890")

func main() {
	fmt.Println(generatePassword(12))

	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите url")

	myAcc := account{
		login:    login,
		password: password,
		url:      url,
	}
	outputData(&myAcc)
}

func promptData(prompt string) string {
	fmt.Println(prompt)
	var res string
	fmt.Scan(&res)
	return res
}

func outputData(acc *account) {
	fmt.Println(acc.login, acc.password, acc.url)
}

func generatePassword(n int) string {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	return string(res)
}
