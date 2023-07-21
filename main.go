package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/skip2/go-qrcode"
)

func main() {
	document, err := os.Open("requirements.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer document.Close()

	values, err := ioutil.ReadAll(document)
	if err != nil {
		log.Fatal(err)
	}

	randomID := uuid.New()
	requirements := string(values)

	path := "/codes"
	filename := randomID.String() + ".png"

	err = qrcode.WriteFile(requirements, qrcode.Medium, 512, path + filename)
	if err != nil {
		log.Fatal(err)
	}
}
