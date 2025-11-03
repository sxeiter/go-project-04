package main

import (
	"demo/password/account"
	"demo/password/encrypter"
	"demo/password/files"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": findAccountByUrl,
	"3": findAccountByLogin,
	"4": deleteAccount,
}

func main() {
	fmt.Println("Менеджер паролей")
	err := godotenv.Load()
	if err != nil {
		color.Red("Не удалось найти env файл")
	}
	vault := account.NewVault(files.NewJsonDb("data.vault"), *encrypter.NewEncryter())
Menu:
	for {
		variant := promptData([]string{
			"1. Создать аккаунт",
			"2. Найти аккаунт по URL",
			"3. Найти аккаунт по логину",
			"4. Удалить аккаунт",
			"5. Выход",
			"Выберите вариант",
		})
		menuFunc := menu[variant]
		if menuFunc == nil {
			break Menu
		}
		menuFunc(vault)
	}
	createAccount(vault)
}

func findAccountByUrl(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите URl для поиска"})
	accounts := vault.FindAccount(url, checkUrl)
	outputResult(&accounts)
}

func findAccountByLogin(vault *account.VaultWithDb) {
	login := promptData([]string{"Введите логин для поиска"})
	accounts := vault.FindAccount(login, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Login, str)
	})
	outputResult(&accounts)
}

func outputResult(accounts *[]account.Account) {
	if len(*accounts) == 0 {
		color.Red("Аккаунтов не найдено!")
	}
	for _, account := range *accounts {
		account.OutputData()
	}
}

func checkUrl(acc account.Account, str string) bool {
	return strings.Contains(acc.Url, str)
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите URl для поиска"})
	isDeleted := vault.DeleteAccountByUrl(url)
	if isDeleted {
		color.Green("Удалено")
	} else {
		color.Red("Не найдено")
	}
}

func createAccount(vault *account.VaultWithDb) {
	login := promptData([]string{"Введите логин"})
	password := promptData([]string{"Введите пароль"})
	url := promptData([]string{"Введите url"})
	myAcc, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URl или логин")
		return
	}
	vault.AddAccount(*myAcc)
}

func promptData[T any](prompt ...T) string {
	for i, line := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v: ", line)
		} else {
			fmt.Println(line)
		}
	}
	var res string
	fmt.Scanln(&res)
	return res
}
