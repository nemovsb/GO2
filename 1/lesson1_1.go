package main

import (
	"fmt"
	"os"
	"time"
)

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
	fileName, myErr := writeFileName()
	if myErr.Err != nil {
		err := fmt.Errorf("Get File name error: %v", myErr.Error())
		fmt.Println(err)
	}

	//	Создание файла
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("fileName: ", err)
	}

	// Открытие файла
	// file, err := os.OpenFile(fileName, os.O_WRONLY, 0666)
	// if err != nil {
	// 	fmt.Println("fileName: ", err)
	// }

	//	Отложенное закрытие файла
	defer func() {
		err = file.Close()
		if err != nil {
			fmt.Println("fileName close: ", file, err)
		}
	}()

	// Запись в файл
	fileContent, err := writeFileContent()
	_, err = file.Write([]byte(fileContent))
	if err != nil {
		fmt.Println("fileName write string: ", file, err)
	}

}
