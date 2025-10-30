package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*1234567890")

type Account struct {
	login    string
	password string
	url      string
}

func (acc *Account) OutputData() {
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}

func NewAccount(login, password, urlString string) (*Account, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	newAcc := &Account{
		url:      urlString,
		password: password,
		login:    login,
	}
	if password == "" {
		newAcc.generatePassword(12)
	}
	return newAcc, nil
}
