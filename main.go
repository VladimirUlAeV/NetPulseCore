package main

import (
	storage "NetPulse/Storage"
	core "NetPulse/Core"
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
	Scanner := bufio.NewScanner(os.Stdin)
	for {
		menu()
		Scanner.Scan()
		InputUser := Scanner.Text()
		switch InputUser {
		case "1":
			core.PingGo()
			time.Sleep(1 * time.Second)
		case "2":
			storage.OpenFile("Result.log")
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
