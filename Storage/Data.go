package storage

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sync"
	"time"
)
var Mu sync.Mutex

func createFile(path string)error{
	file, err := os.Create(path)
	if err != nil{
		return err
	}
	defer file.Close()
	return nil
}

func createOrNot(filename string)bool{
	_, err := os.Stat(filename)
	if err == nil{
		return true
	}else if os.IsNotExist(err){ //нужна именно эта ошибка
		return false 
	} 
	return false
}

func EditLog(url string, code int, ping time.Duration, errPing error)error{
	defer Mu.Unlock()
	Mu.Lock()
	timestamp := time.Now()
	filename := "Result.log"
	if !createOrNot(filename){
		createFile(filename)
	}

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil{
		return err
	}
	defer file.Close()

	_, err = fmt.Fprintf(file,
		"Запрос на: %s\n"+
    	"Дата записи: %v\n"+
    	"Статус запроса: %d\n"+
    	"Пинг: %v\n"+
    	"Ошибка: %v\n\n",
    	url, timestamp, code, ping, errPing,)
	return err
}

func OpenFile(path string)error{
	var cmd *exec.Cmd
	switch runtime.GOOS{ //функция для проверки ОС
	case "windows":
		cmd = exec.Command("cmd", "/c", "", path)
	case "darwin": //MacOS
		cmd = exec.Command("open", path)
	default://Linus и  Unix ОС
		cmd = exec.Command("xdg-open", path)
	}
	return cmd.Start() //запуск окна
}