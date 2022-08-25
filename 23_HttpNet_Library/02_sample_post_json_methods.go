package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	requestData := map[string]any{"Name": "Kido", "Age": 20, "Job": "Developer"}
	json_data, err := json.Marshal(requestData)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("http://localhost:8080/echo", "application/json", bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	var resMap map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&resMap)

	fmt.Println(resMap)

}
