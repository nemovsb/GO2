package main

import (
	"fmt"
	"os"
	"time"
)

//Для закрепления практических навыков программирования, напишите программу,
//которая создаёт один миллион пустых файлов в известной, пустой директории файловой системы используя вызов os.Create.
//Ввиду наличия определенных ограничений операционной системы на число открытых файлов,
//такая программа должна выполнять аварийную остановку. Запустите программу и дождитесь полученной ошибки.
//Используя отложенный вызов функции закрытия файла, стабилизируйте работу приложения.
//Критерием успешного выполнения программы является успешное создание миллиона пустых файлов в директории.

//	Пользовательская ошибка
type MyError struct {
	Date string
	Err  error
}

func (e *MyError) Error() string {
	e.Date = string(time.Now().Format("01-02-2006 15:04:05.000000"))
	return e.Date + " " + e.Err.Error()
}

func (e *MyError) Unwrap() error {
	return e.Err
}

//	Функция считывает имя для нового файла.
//  Возвращает строку с именем и
//	ошибку с сохранением времени возникновения.
func writeFileName() (string, MyError) {
	var filePath string
	fmt.Print("FIle Path: ")
	_, err := fmt.Scanln(&filePath)
	return filePath, MyError{Err: err}
}

//	Функция считывает содержимое для нового файла.
//  Паника вызывается если произошла ошибка,
//	например, введена пустая строка (unexpected newline)
func writeFileContent() (string, error) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Printf("panic with value: %v\n", v)
		}
	}()
	var fileContent string
	fmt.Print("FIle Content: ")
	_, err := fmt.Scanln(&fileContent)
	if err != nil {
		panic(err)
	}
	return fileContent, err
}

func main() {

	defer func() {
		if v := recover(); v != nil {
			fmt.Printf("panic with value: %v\n", v)
		}
	}()

	for i := 0; i < 1048576; i++ {

		_, err := os.Create(fmt.Sprintf("file %d", i))
		if err != nil {
			fmt.Println("fileName: ", err)

			//return
		}

		// file, err := os.OpenFile(fmt.Sprintf("file %d", i), os.O_WRONLY, 0666)
		// if err != nil {
		// 	fmt.Println("fileName: ", err)
		// }

		// _, err = file.Write([]byte(fmt.Sprintf("string %d", i)))
		// if err != nil {
		// 	fmt.Println("fileName write string: ", file, err)
		// }

	}

}
