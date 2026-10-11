package main

import (
	"fmt"
	"os"
)

// Récupèrer les arguments avec os.Args et vérifie qu'il y en a bien deux.
// Lire le fichier d'entrée avec os.ReadFile et recopier le tel quel dans le fichier de sortie avec os.writeFile.
// Gèrer les erreurs proprement.
func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Usage: go run . <fichier_entree> <fichier_sortie>")
		os.Exit(1)
	}

	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Erreur de lecture:", err)
		os.Exit(1)
	}

	result := Process(string(input))

	if err := os.WriteFile(os.Args[2], []byte(result), 0o644); err != nil {
		fmt.Fprintln(os.Stderr, "Erreur d'écriture:", err)
		os.Exit(1)
	}
}
