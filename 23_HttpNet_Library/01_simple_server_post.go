package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/echo", mainPage)
	fmt.Println("Server Start with 8080 port.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type UserInfo struct {
	Name string
	Age  int
	Job  string
	Tm   time.Time
}

func mainPage(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "POST":
		w.WriteHeader(http.StatusOK)

		var userInfo UserInfo

		b, err := io.ReadAll(r.Body)

		if err != nil {
			log.Fatal(err)
		}

		json.Unmarshal(b, &userInfo)

		userInfo.Tm = time.Now()

		writeByte, _ := json.Marshal(userInfo)
		w.Write(writeByte)
	}
}
