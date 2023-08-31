package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// Replace with your actual token and secret
	token := os.Getenv("BRI_CLIENT_ID")
	secret := os.Getenv("BRI_CLIENT_SECRET")

	account := "888801000157508"
	url := "https://bri-partner-sandbox.deno.dev/v2/inquiry/" + account

	timestamp := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"BRI-Signature": "",
		"BRI-Timestamp": timestamp,
	}

	stringToSign := fmt.Sprintf("GET\n%s\n/v2/inquiry/%s\n", timestamp, account)
	mac := hmac.New(sha256.New, []byte(secret))
	_, err := mac.Write([]byte(stringToSign))
	if err != nil {
		log.Fatal("Error while writing HMAC: ", err)
	}
	signature := hex.EncodeToString(mac.Sum(nil))
	headers["BRI-Signature"] = signature

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error creating request: ", err)
	}
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	response, err := client.Do(req)
	if err != nil {
		log.Fatal("Error sending request: ", err)
	}
	defer response.Body.Close()

	var responseJSON map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&responseJSON)
	if err != nil {
		log.Fatal("Error decoding JSON response: ", err)
	}

	fmt.Println("Credentials from Envars")
	fmt.Println("========================")
	fmt.Println("Client ID:", token)
	fmt.Println("Client Secret:", secret)

	fmt.Println("Response from API")
	fmt.Println("=================")
	fmt.Println("Response status code:", response.StatusCode)
	fmt.Println("Response content:")
	prettyJSON, _ := json.MarshalIndent(responseJSON, "", "    ")
	fmt.Println(string(prettyJSON))
}
