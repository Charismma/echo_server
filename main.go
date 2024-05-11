package main

import (
	"fmt"
	"log"
)

func main() {
	var data string
	fmt.Print("Введите данные: ")
	_, err := fmt.Scan(&data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Вы ввели следующие данные: %s", data)
}
