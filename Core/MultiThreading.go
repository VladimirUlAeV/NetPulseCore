package core

import (
	storage "NetPulse/Storage"
	"sync"
	"bufio"
	"fmt"
	"os"
)



func PingGo(HowMany int){
	Pointer := storage.PointerColector()
	Scanner := bufio.NewScanner(os.Stdin)
	var wg sync.WaitGroup
	for i := 0; i < HowMany ; i++{
		wg.Add(1)
		fmt.Print("Введите адресс или домен: ")
		Scanner.Scan()
		Url := Scanner.Text()
		go func(u string){
			defer wg.Done()
			url, code, timee, err := ping(u) 
			Pointer.NewResult(url, code, timee, err)
		}(Url)
	}
	wg.Wait()
}