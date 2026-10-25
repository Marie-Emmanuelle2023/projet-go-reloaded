package main

import (
	"fmt"
	"go-reloaded/rules"
	"os"
)

func main() {
	// Attend : go run . <entrée> <sortie>
	if len(os.Args) != 3 {
		fmt.Println("Usage : go run . <entrée> <sortie>")
		return
	}

	contenu, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Erreur de lecture :", err)
		return
	}

	// Applique toutes les règles puis écrit le résultat
	resultat := rules.Traitement(string(contenu))
	err = os.WriteFile(os.Args[2], []byte(resultat), 0644)
	if err != nil {
		fmt.Println("Erreur d'écriture :", err)
	}
}
