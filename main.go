package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"
)

func main() {
	files.WriteFile("hello i am file", "file.txt")
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите url")
	myAcc, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URl или логин")
		return
	}
	myAcc.OutputData()

}

func promptData(prompt string) string {
	fmt.Println(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}
