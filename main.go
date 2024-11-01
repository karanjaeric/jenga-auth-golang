package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Credential struct {
	MerchantCode   string `json:"merchantCode"`
	ConsumerSecret string `json:"consumerSecret"`
}

func main() {

	credentialPayload := Credential{MerchantCode: "8390533684", ConsumerSecret: "IiFb819ao6MR4CHAImHknz7n78gZ31"}
	jsonCredentialPayload, err := json.Marshal(credentialPayload)
	if err != nil {
		log.Fatalf("error creating json credential payload %v", err)
	}
	req, err := http.NewRequest(http.MethodPost, "https://uat.finserve.africa/authentication/api/v3/authenticate/merchant",
		bytes.NewBuffer(jsonCredentialPayload))
	if err != nil {
		log.Fatalf("error creating request %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Api-Key", "cWLXytP1CzYUj4D3wve6htKpeXjZikFNu6wJWnGBlEVJRTp1DASkQ++d8px8eOmm5xJhBXwRHb3cPLhbxkYLqg==")
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("error invoking jenga auth api %v", err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("error extracting request body")
	}
	fmt.Println("response body is ", string(body))

}
