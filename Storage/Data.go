package storage

import (
	"sync"
	"log"
	"os"
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

func PointerColector()(*collector){
	return Colector
}
var Colector = newcollector()

func File(){
	file, err := os.Create("Resut.txt")
	if err != nil{
		log.Fatal(err)}
	defer file.Close()
	//запись в файл
	/*
	writerFile := bufio.NewWriter(file)
	
	writer.WriteString(данные)
	
	Сброс буфера
	if err := writer.Flush(); err != nil{
	log.Fatal(err)}
	*/
}

func newcollector() *collector{
	return &collector{
		DataReport: make([]result, 0),
	}
}

func (c *collector)NewResult(Url string, StCode int, timeDur time.Duration, err error) {
	defer c.Mu.Unlock()
	c.Mu.Lock()
	c.DataReport = append(Colector.DataReport,result{
			URL: Url,
			StatusCode: StCode,
			TimeDuration: timeDur,
			Err: err,
		})
}

func (c *collector)PrintResult(){
	
}
