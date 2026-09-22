package core

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)


func adressFullOrNot(Url string)(string, error){
	if Url == "" || Url == " "{
		return Url, errors.New("Пустое значение")
	}
	if !strings.HasPrefix(Url,"https://"){
		return "https://"+Url, nil
	}
	return Url, nil
}

func ping(Url string) (string, int, time.Duration, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	fullUrl, err := adressFullOrNot(Url)
	if err != nil{
		fmt.Println("Ошибка:", err)
		return Url, 0, -1, err
	}
	
	req, err := http.NewRequest(http.MethodPost, fullUrl, nil) //параметры запроса
	if err != nil {
		fmt.Println("Ошибка:", err)
		return Url, 0, -1, err
	}

	ping := time.Now()
	fmt.Println("Запрос на..."+fullUrl)
	resp, err := client.Do(req) //Делаем запрос
	pingClose := time.Since(ping)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return Url, 0, -1, err
	}
	defer resp.Body.Close()
	fmt.Println(Url, resp.StatusCode, time.Duration(pingClose.Milliseconds()))
	return Url, resp.StatusCode, time.Duration(pingClose.Milliseconds()), nil
}
