package core

import (
	storage "NetPulse/Storage"
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func PingGo(){
	Scanner := bufio.NewScanner(os.Stdin)
	var wg sync.WaitGroup
	fmt.Print("Введите адресс или домен: ")
	Scanner.Scan()
	UrlText := Scanner.Text()
	Url := strings.Fields(UrlText)
	for _, v := range Url{
		wg.Add(1)
		go func(string){
			defer wg.Done()
			url, code, timee, err := ping(v) 
			storage.EditLog(url, code, timee, err)
			}(v)
		}
	wg.Wait()
}
