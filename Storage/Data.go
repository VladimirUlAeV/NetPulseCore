package storage

import (
	"os"
	"os/exec"
	"runtime"
	"sync"
	"time"
)

type result struct {
	URL          string
	StatusCode   int
	TimeDuration time.Duration
	Err error
}

type collector struct{
	Mu sync.Mutex
	DataReport []result
}

var colector = newcollector()

func newcollector() *collector{
	return &collector{
		DataReport: make([]result, 0),
	}
}

func PointerColector()(*collector){
	return colector
}

func (c *collector)NewResult(Url string, StCode int, timeDur time.Duration, err error) {
	defer c.Mu.Unlock()
	c.Mu.Lock()
	c.DataReport = append(colector.DataReport,result{
			URL: Url,
			StatusCode: StCode,
			TimeDuration: timeDur,
			Err: err,
		})
}

func (c *collector)PrintResult(){
	
}

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

func EditLog()error{
	filename := "Result.log"
	if !createOrNot(filename){
		createFile(filename)
	}

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil{
		return err
	}
	defer file.Close()

	_, err = file.WriteString(filename + "\n") //реализовать запись
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