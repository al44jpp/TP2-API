package affichagepilote

import (
	read_json "TP2/Json"
	"fmt"
)

func Affichagepilote(liste []read_json.Driver) {
	if liste == nil {
		return
	}
	for _, driver := range liste {
		fmt.Println("---   pilote n.", driver.Number, "   ---")
		fmt.Println("nom complet: ", driver.Firstname, driver.Lastname)
		fmt.Println("nationalité : ", driver.Nationality)
		fmt.Println("code : ", driver.Code)
	}
}
