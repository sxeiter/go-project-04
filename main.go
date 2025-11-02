package main

import (
	"demo/password/account"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("Менеджер паролей")
	vault := account.NewVault()
Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
		default:
			break Menu
		}
	}
	createAccount(vault)
}

func getMenu() int {
	var variant int
	fmt.Println("Выберите вариант:")
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")
	fmt.Scan(&variant)
	return variant
}

func findAccount(vault *account.Vault) {
	url := promptData("Введите URl для поиска")
	accounts := vault.FindAccountByUrl(url)
	if len(accounts) == 0 {
		color.Red("Аккаунтов не найдено!")
	}
	for _, account := range accounts {
		account.OutputData()
	}

}

func deleteAccount(vault *account.Vault) {
	url := promptData("Введите URl для поиска")
	isDeleted := vault.DeleteAccountByUrl(url)
	if isDeleted {
		color.Green("Удалено")
	} else {
		color.Red("Не найдено")
	}
}

func createAccount(vault *account.Vault) {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите url")
	myAcc, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URl или логин")
		return
	}
	vault.AddAccount(*myAcc)
}

func promptData(prompt string) string {
	fmt.Println(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}
