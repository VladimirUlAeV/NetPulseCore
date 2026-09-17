package main

import (
//	storage "NetPulse/Storage"
	"bufio"
	"fmt"
	"os"
	"time"
)

func menu() {
	fmt.Println("Выберите действие")
	fmt.Println("1 - пинг сайтов")
	fmt.Println("2 - показать отчет")
	fmt.Println("0 - выход")
}

func main() {
	fmt.Println("Приветсвую в NetPulseCore!")
//	storage.File()
	Scanner := bufio.NewScanner(os.Stdin)
	for {
		menu()
		Scanner.Scan()
		InputUser := Scanner.Text()
		switch InputUser {
		case "1":
			// Test Method
			/*Url := "Google"
			StCode := 200
			timeDur :=  40 * time.Millisecond
			storage.Colector.NewResult(Url, StCode, timeDur)
			fmt.Println(storage.Colector) */
			//Метод для пинга

			time.Sleep(1 * time.Second)
		case "2":
			//Отчет
			time.Sleep(1 * time.Second)
		case "0":
			fmt.Println("Спасибо за использование моей программы")
			return
		default:
			fmt.Println("Команда не распознана")
			time.Sleep(1 * time.Second)
		}
	}
}
