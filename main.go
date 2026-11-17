package main

import (
	api "TP2/API"
	read_json "TP2/Json"
	affichagepilote "TP2/affichage_pilote"
	"TP2/fetch"
	"fmt"
	"strconv"
)

func demanderNombre() int {
	fmt.Print("sélectionnez un numéro de pilote : ")
	selection := ""
	fmt.Scan(&selection)
	if selection == "all" || selection == "ALL" {
		fmt.Println("\nvous avez choisi d'afficher les informations de tous les pilotes.")
		return -1
	} else {
		if selection == "" {
			return demanderNombre()
		}
		i, err := strconv.Atoi(selection)
		if err != nil {
			fmt.Println("erreur: ", selection, " n'est pas un nombre valide. (doit être un entier positif ou \"all\")")
			return demanderNombre()
		} else {
			if i < 0 {
				fmt.Println("le nombre saisi est invalide. (doit être un entier positif ou \"all\")")
				return demanderNombre()
			}
		}
		return i
	}
}

func main() {
	selection := -1
	selection = demanderNombre()
	if selection != -1 {
		fmt.Println("vous avez choisi: ", selection)
	}
	Json := api.Appel_API()
	liste_a_afficher, err := fetch.Fetch(selection, read_json.Read_data(Json))
	if err != nil {
		fmt.Println(err)
	} else {
		affichagepilote.Affichagepilote(liste_a_afficher)
	}
}
