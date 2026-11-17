package api

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func Appel_API() []byte {
	client := http.Client{Timeout: time.Duration(time.Second * 10)}
	res, err := client.Get("https://f1-companion-api-ba5k.onrender.com/drivers")
	if err != nil {
			fmt.Println(err)
			return nil
		}
	if res.Status == "200 OK" {
		body, err := io.ReadAll(res.Body)
		defer res.Body.Close()
		if err != nil {
			fmt.Println(err)
			return nil
		}
		return body
	}
	fmt.Println("Erreur: code status ", res.Status)
	return nil	
}
