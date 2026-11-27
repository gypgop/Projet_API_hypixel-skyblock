package api

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"initial/models"
	"net/http"
	"time"
)

var client = &http.Client{Timeout: 5 * time.Second}

var ErrPlayerNotFound = errors.New("Unknown player")

func GetPlayerUUID(PlayerName string) (models.PlayerUUId, []byte, error) {

	fmt.Printf("Entrée GetPlayerUUID\n")

	var baseURL = "https://api.mojang.com/users/profiles/minecraft/"

	endpoint := fmt.Sprintf("%s%s", baseURL, PlayerName)

	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		// fmt.Printf("erreur, %v", err)
		return models.PlayerUUId{}, []byte{}, ErrPlayerNotFound
	}

	fmt.Printf("%v\n", request)

	response, err := client.Do(request)

	if err != nil {
		// fmt.Printf("erreur, %v", err)
		return models.PlayerUUId{}, []byte{}, ErrPlayerNotFound
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return models.PlayerUUId{}, []byte{}, ErrPlayerNotFound

	}

	var PlayerUUID models.PlayerUUId
	err = json.NewDecoder(response.Body).Decode(&PlayerUUID)
	if err != nil {
		return models.PlayerUUId{}, []byte{}, err
	}

	PlayerSkin, err := GetPlayerSkin(PlayerUUID)

	return PlayerUUID, PlayerSkin, err

}

func GetPlayerSkin(PlayerUUID models.PlayerUUId) ([]byte, error) {

	fmt.Printf("Entrée GetPlayerSkin\n")

	var baseURL = "https://sessionserver.mojang.com/session/minecraft/profile/"

	endpoint := fmt.Sprintf("%s%s", baseURL, PlayerUUID.UUID)

	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		fmt.Printf("erreur, %v\n", err)
		return []byte{}, err
	}

	fmt.Printf("%v\n", request)

	response, err := client.Do(request)

	if err != nil {
		fmt.Printf("erreur, %v\n", err)
		return []byte{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return []byte{}, ErrPlayerNotFound

	}

	var PlayerModel models.PlayerProperties
	err = json.NewDecoder(response.Body).Decode(&PlayerModel)
	if err != nil {
		return []byte{}, err
	}

	fmt.Printf("%v\n", PlayerModel)

	if len(PlayerModel.Properties) == 0 {
		return []byte{}, errors.New("Mojang response has no player properties")
	}

	decoded, err := base64.StdEncoding.DecodeString(PlayerModel.Properties[0].Value)
	if err != nil {
		return []byte{}, err
	}

	fmt.Printf("Base64 Decoder : %s\n", decoded)

	return decoded, nil

}
