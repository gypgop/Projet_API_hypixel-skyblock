package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"initial/models"
	"net/http"
	"time"
)

const baseURL = "https://api.mojang.com/users/profiles/minecraft/"

var client = &http.Client{Timeout: 5 * time.Second}

var ErrPlayerNotFound = errors.New("Unknown player")

func GetPlayerUUID(PlayerName string) (models.PlayerUUId, error) {

	endpoint := fmt.Sprintf("%s%s", baseURL, PlayerName)

	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		fmt.Printf("erreur, %v", err)
		return models.PlayerUUId{}, err
	}

	fmt.Printf("%v\n", request)

	response, err := client.Do(request)

	if err != nil {
		fmt.Printf("erreur, %v", err)
		return models.PlayerUUId{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return models.PlayerUUId{}, ErrPlayerNotFound

	}

	var PlayerUUID models.PlayerUUId
	err = json.NewDecoder(response.Body).Decode(&PlayerUUID)
	return PlayerUUID, err

}
