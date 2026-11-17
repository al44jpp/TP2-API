package fetch

import (
	read_json "TP2/Json"
	"errors"
	"fmt"
	"strconv"
)

func Fetch(n int, drivers []read_json.Driver) ([]read_json.Driver, error) {
	if drivers == nil{
		return nil,errors.New("impossible de récupérer la liste des pilotes.")
	}
	if n == -1 {
		return drivers, nil
	}
	for _, driver := range drivers {
		if driver.Number == strconv.Itoa(n) {
			return []read_json.Driver{driver}, nil
		}
	}
	err := errors.New(fmt.Sprint("le nombre ", n, " ne correspond à aucun pilote."))
	return nil, err
}
