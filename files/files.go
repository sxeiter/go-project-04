package files

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func ReadFile() {}

func WriteFile(content string, name string) {
	file, err := os.Create(name)
	if err != nil {
		color.Red("Ошибка")
	}
	_, err = file.WriteString(content)
	if err != nil {
		file.Close()
		color.Red(err.Error())
		return
	}
	fmt.Println("Запись успешна")
	file.Close()
}
