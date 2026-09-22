package main

import (
	storage "NetPulse/Storage"
	core "NetPulse/Core"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func menu() {
	fmt.Println("Выберите действие")
	fmt.Println("1 - пинг сайтов")
	fmt.Println("2 - показать отчет")
	fmt.Println("0 - выход")
}

func main() {
	Pointer := storage.PointerColector()
	fmt.Println("Приветсвую в NetPulseCore!")
//	storage.File()
	Scanner := bufio.NewScanner(os.Stdin)
	for {
		menu()
		Scanner.Scan()
		InputUser := Scanner.Text()
		switch InputUser {
		case "1":
			fmt.Print("Введите количество запросов: ")
			Scanner.Scan()
			Input := Scanner.Text()
			HowMany, err := strconv.Atoi(Input)
			if err != nil{
				fmt.Println("Ошибка:", err)
			}
			core.PingGo(HowMany)
			time.Sleep(1 * time.Second)
		case "2":
			//Отчет
			fmt.Println(Pointer)
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
