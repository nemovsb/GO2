package main

import (
	"flag"
	"fmt"
	"hash/crc32"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
)

//Необходимая информация о файле
type File struct {
	Path        string
	Hash        uint32
	Size        int64
	ISDuplicate bool
}

//Функция принимает на вход путь до файла и возвращает хэш его содержимого
func getFileHash(filePath string) uint32 {
	defer func() {
		if v := recover(); v != nil {
			fmt.Printf("panic. Can't open the file %s: %v\n", filePath, v)
		}
	}()

	file, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Printf("filePath Error: %v\n", err)
		return 0
	}
	defer func() {
		err = file.Close()
		if err != nil {
			fmt.Println("fileName close: ", file, err)
		}
	}()

	//Читаем содержимое файла
	data := make([]byte, 64)
	for {
		_, err := file.Read(data)
		if err == io.EOF {
			break
		}
	}
	//Получаем хэш содержимого и возвращаем его
	var hash = crc32.NewIEEE()
	_, err = hash.Write(data)

	if err != nil {
		panic(err)
	}

	return hash.Sum32()
}

func main() {
	//Флаги
	var (
		dirForCheck string
		needDelete  *bool
	)
	flag.StringVar(&dirForCheck, "d", "", "directory for checking duplicate files")
	needDelete = flag.Bool("rm", false, "delete flag")
	flag.Parse()
	if !flag.Parsed() {
		log.Fatal("Flag not parsed")
	}

	//Проходим по файлам в выбранной директории и записываем нужную
	//для поиска дубликатов информацию в массив структур File
	files := make([]File, 0)
	err := filepath.Walk(dirForCheck, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, File{
				Path: path,
				Size: info.Size(),
				Hash: getFileHash(path),
			})
		}
		//log.Println(path, info.Size())
		return nil
	})
	if err != nil {
		log.Println(err)
	}
	//fmt.Println("files: ", files)

	//Помечаем дубли флагом ISDuplicate=true
	for fileNum := range files {
		if files[fileNum].ISDuplicate == false {
			for i := fileNum + 1; i < len(files); i++ {
				if (files[fileNum].Hash == files[i].Hash) && (files[fileNum].Size == files[i].Size) {
					files[i].ISDuplicate = true
					//log.Printf("Original : %v\n", files[fileNum])
					log.Printf("Duplicate: %v\n", files[i])
				}
			}
		}
	}
	//fmt.Println("files: ", files)
	log.Printf("need del?: %v\n", *needDelete)
	//os.Remove(filePath)

	//Удаляем дубликаты
	wg := sync.WaitGroup{}
	n := 2
	fileToDelPathCh := make(chan string, n)

	if *needDelete == true {
		//Пишем в канал, какие файлы хотим удалить
		go func() {
			for fileNum := range files {
				if files[fileNum].ISDuplicate == true {
					//log.Printf("fileToDelPathCh <- %s", files[fileNum].Path)
					fileToDelPathCh <- files[fileNum].Path
				}
			}
			close(fileToDelPathCh)
		}()

		wg.Add(n)
		for i := 0; i <= n; i++ {
			//Удаляем файлы
			go func(gouroutineID int) {
				defer wg.Done()
				for fileToDel := range fileToDelPathCh {
					//fileToDel := <-fileToDelPathCh
					//log.Printf("FIle %s ready to delete\n", fileToDel)
					if fileToDel == "" {
						break
					}
					err = os.Remove(fileToDel)
					if err != nil {
						log.Println(err)
					}
					log.Printf("FIle %s removed.\n", fileToDel)

				}

				log.Printf("Gorutine %d done\n", gouroutineID)

			}(i)
		}
	}
	wg.Wait()
}
