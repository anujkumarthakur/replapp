package logic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"replapp/model"
)

const (
	LIMIT = "10"
)

func GetAllMarvalCharacter(name string, offset string) (*model.SearchCharacterResponse, error) {
	baseUrl := os.Getenv("BASE_URL")
	apiKey := os.Getenv("API_KEY")
	characterPath := os.Getenv("CHARACTERSPATH")
	response, err := http.Get(baseUrl + characterPath + "?name=" + name + "&limit=" + LIMIT + "&offset=" + offset + "&apikey=" + apiKey)

	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var responseObject model.SearchCharacterResponse
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return nil, err
	}
	return &responseObject, nil
}
