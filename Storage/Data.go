package storage

import (
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
	DataReport []result
}

var colector = newcollector()

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

func (c *collector)NewResult(Url string, StCode int, timeDur time.Duration) {
c.DataReport = append(colector.DataReport,result{
			URL: Url,
			StatusCode: StCode,
			TimeDuration: timeDur,
			Err: nil,
		})
}

func (c *collector)PrintResult(){
	
}
