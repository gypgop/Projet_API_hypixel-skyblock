package controllers

import (
	"fmt"
	"net/http"
	"strings"
)

func SkyTrackPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DEBUG: Requête reçue sur /menu")

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		fmt.Println("DEBUG: Méthode non autorisée :", r.Method)
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	playerName := strings.TrimSpace(r.URL.Query().Get("PlayerName"))
	if playerName == "" {
		fmt.Println("DEBUG: Paramètre PlayerName manquant ou vide")
		http.Error(w, "Veuillez indiquer un PlayerName", http.StatusBadRequest)
		return
	}

	fmt.Printf("DEBUG: Succès ! PlayerName = %s\n", playerName)
	fmt.Fprintf(w, "Joueur trouvé : %s", playerName)

}
