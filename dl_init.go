package main

import (
	"math/rand"
	"os"
	"time"
	_"io/ioutil"
	_"log"
	"fmt"
)

func get(min, max int64) (int64) { // генерация псевдослучайных (!!! - не путать с действительно случайными) чисел
	return rand.Int63n(max-min) + min
}

func main() {
	rand.Seed(time.Now().Unix()) // "сеим" случайности в основоной функции
	d := get(0, 1000000) // генратор случаныхх значений
	arg := os.Args[1] //передача в качестве аргумента первого элемента командной строки
	fd := fmt.Sprintf("%s_%d_%s","init", d, arg) // задаем имя каталога сессии
	os.Mkdir(fd, os.FileMode(0777)) // создание каталога со случайным именем
	fmt.Println(os.Getwd())
	os.Chdir(fd)
	fmt.Println(os.Getwd())
	file, err := os.Create("init.dl")

	if err != nil {
		// handle the error here
		return
	}
	defer file.Close()
	file.Chdir()
	s := fmt.Sprintf("%s%d\n\n","создание файловой системы -", d)
	file.WriteString(s)

}