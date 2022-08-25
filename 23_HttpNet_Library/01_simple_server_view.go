package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/view/", viewPage)
	fmt.Println("Server Start with 8080 port.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{\"Hello\": \"World\"}")
}

func viewPage(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[len("/view/"):]
	fmt.Println("--------- ", path)
	fileName := fmt.Sprintf("%s.%s", path, "html")
	writeFileToHttp(w, fileName)
}

func writeFileToHttp(w http.ResponseWriter, fileName string) {

	buff := make([]byte, 1024)

	sourceFile, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Fail to oepn source file/n", err)
	}

	defer sourceFile.Close()

	for {
		readSize, err := sourceFile.Read(buff)
		if err != nil && err != io.EOF {
			fmt.Errorf("Error occur when read file data.", err)
		}

		if readSize == 0 {
			break
		}

		writeSize, err := w.Write(buff[:readSize])
		if err != nil {
			fmt.Errorf("Error occur write data to dest file", err)
		}
		fmt.Printf("WriteFile Size is %d \n", writeSize)
	}
}
