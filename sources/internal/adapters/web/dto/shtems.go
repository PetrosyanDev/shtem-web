// Erik Petrosyan ©
package dto

import (
	"encoding/json"
	"log"
	"net/http"
	"shtem-web/sources/internal/core/domain"
)

type ShtemsResponse struct {
	Data []string `json:"data"`
}

func ShtemsData() *domain.Page {
	const (
		title       = "shtemaran.am • Learning Fast | The Best Way to Save Time"
		description = "Welcome to shtemaran.am"
	)

	req, err := http.NewRequest(http.MethodPost, "https://shtemaran.am/api/v1/questions/getShtems", nil)
	if err != nil {
		log.Println("Failed to create request")
		return nil
	}

	req.Header.Set("X-Shtem-Api-Key", "someKey")

	// Make the request using the http.Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Failed to make request")
		return nil
	}
	defer resp.Body.Close()

	var shtemsResponse ShtemsResponse
	err = json.NewDecoder(resp.Body).Decode(&shtemsResponse)
	if err != nil {
		log.Println("Failed to decode response")
		return nil
	}

	pb := newPageBuilder().
		AddHeader(title, description).
		AddTopMenuItem("HOME", "/", false).
		AddTopMenuItem("SHTEMS", "/shtems", true).
		AddTopMenuItem("ABOUT", "/about", false).
		AddShtemNames(shtemsResponse.Data)

	return pb.Page()

}
