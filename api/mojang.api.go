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

func GetPlayerUUID(PlayerName string) (models.PlayerUUId, []byte, string, error) {

	fmt.Printf("Enter GetPlayerUUID\n")

	var baseURL = "https://api.mojang.com/users/profiles/minecraft/"

	endpoint := fmt.Sprintf("%s%s", baseURL, PlayerName)

	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		// fmt.Printf("erreur, %v", err)
		return models.PlayerUUId{}, []byte{}, "-1", ErrPlayerNotFound
	}

	fmt.Printf("\n%v\n\n", request)

	response, err := client.Do(request)

	if err != nil {
		// fmt.Printf("erreur, %v", err)
		return models.PlayerUUId{}, []byte{}, "-1", ErrPlayerNotFound
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return models.PlayerUUId{}, []byte{}, "-1", ErrPlayerNotFound

	}

	var PlayerUUID models.PlayerUUId
	err = json.NewDecoder(response.Body).Decode(&PlayerUUID)
	if err != nil {
		return models.PlayerUUId{}, []byte{}, "-1", err
	}

	PlayerSkin, name, err := GetPlayerSkin(PlayerUUID)

	fmt.Printf("playername returned from getplayerSkin = %s\n", name)
	return PlayerUUID, PlayerSkin, name, err

}

func GetPlayerSkin(PlayerUUID models.PlayerUUId) ([]byte, string, error) {

	fmt.Printf("Enter GetPlayerSkin\n")

	var baseURL = "https://sessionserver.mojang.com/session/minecraft/profile/"

	endpoint := fmt.Sprintf("%s%s", baseURL, PlayerUUID.UUID)

	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		fmt.Printf("erreur, %v\n", err)
		return []byte{}, "-1", err
	}

	fmt.Printf("\n%v\n\n", request)

	response, err := client.Do(request)

	if err != nil {
		fmt.Printf("erreur, %v\n", err)
		return []byte{}, "-1", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return []byte{}, "-1", ErrPlayerNotFound

	}

	var PlayerModel models.PlayerProperties
	err = json.NewDecoder(response.Body).Decode(&PlayerModel)
	if err != nil {
		return []byte{}, "-1", err
	}

	// fmt.Printf("%v\n", PlayerModel)

	if len(PlayerModel.Properties) == 0 {
		return []byte{}, "-1", errors.New("Mojang response has no player properties")
	}

	decoded, err := base64.StdEncoding.DecodeString(PlayerModel.Properties[0].Value)
	if err != nil {
		return []byte{}, "-1", err
	}

	var PlayerSkin models.PlayerInfo
	PlayerSkin, err = IsolatePlayerInfo(decoded)
	fmt.Printf("PlayerSkin = ", PlayerSkin)
	return []byte(PlayerModel.Properties[0].Value), PlayerModel.ProfileName, nil

}

func IsolatePlayerInfo(decoded []byte) (models.PlayerInfo, error) {

	fmt.Printf("entrée IsolatePlayerInfo with Decoder = %s\n", decoded)
	var PlayerData models.PlayerInfo
	err := json.Unmarshal(decoded, &PlayerData)
	fmt.Printf("\nIsolate: %s %s\n", PlayerData, err)
	return models.PlayerInfo{}, err
}
