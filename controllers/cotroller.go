package controllers

import (
	"fmt"
	"initial/api"
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
	PlayerUUID, PlayerSkin, err := api.GetPlayerUUID(playerName) // should be calling HypixelSkyblock APIs and hypixel API should call GetPlayerUUID

	fmt.Printf("DEBUG: Succès ! PlayerName = %s\n", playerName)
	fmt.Printf("%s\n%s\n %v\n", PlayerUUID, PlayerSkin, err)
	fmt.Fprintf(w, "Joueur trouvé : %s", playerName)

}
