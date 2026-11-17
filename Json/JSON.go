package read_json

import (
	"encoding/json"
	"fmt"
)

type data_json struct {
	Season        string   `json:"season"`
	Total_drivers int      `json:"total_drivers"`
	Drivers       []Driver `json:"drivers"`
}

type Driver struct {
	Driverid    string `json:"driverid"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Nationality string `json:"nationality"`
	Number      string `json:"number"`
	Code        string `json:"code"`
}

func Read_data(data []byte) []Driver {
	var res data_json
	err := json.Unmarshal(data, &res)
	if err != nil {
		fmt.Println(err)
	}
	return res.Drivers
}
