package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func getTimeErr(err error) error {
	return errors.New(fmt.Sprintf("%s %s", string(time.Now().Format("01-02-2006 15:04:05.000000")), err))

}

func readFileName() (string, error) {
	var filePath string
	fmt.Print("FIle Path: ")
	_, err := fmt.Scanln(&filePath)
	if err != nil {
		//panic(myErr.Error(err))
		fmt.Println(getTimeErr(err))
	}
	return filePath, err
}

func main() {
	fileName, err := readFileName()
	if err != nil {
		err = fmt.Errorf("Get File name error: %v", err)
		fmt.Println(getTimeErr(err))
	}

	file, err := os.Create(fmt.Sprintf("~/%s", fileName))

}
