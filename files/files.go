package files

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)
	if err != nil {
		color.Red("Ошибка")
	}
	_, err = file.Write(content)
	defer file.Close()
	if err != nil {
		color.Red(err.Error())
		return
	}
	fmt.Println("Запись успешна")
	file.Close()
}
