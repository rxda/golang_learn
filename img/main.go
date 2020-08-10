package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	src := "https://sentry--demo-api.demo.querycap.com/bff-iep/v1/objects/aiMaterials/1195169003028299781?secure=true"
	res, err := http.Get(src)

	if err != nil {
		log.Fatalf("http.Get -> %v", err)
	}

	// We read all the bytes of the image
	// Types: data []byte
	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalf("ioutil.ReadAll -> %v", err)
	}

	// You have to manually close the body, check docs
	// This is required if you want to use things like
	// Keep-Alive and other HTTP sorcery.
	res.Body.Close()

	// You can now save it to disk or whatever...
	ioutil.WriteFile("google_logo.png", data, 0666)

	log.Println("I saved your image buddy!")

	croppedImg, err := cutter.Crop(img, cutter.Config{
		Width:  250,
		Height: 500,
	})
	ioutil.WriteFile("google_logo.png", data, 0666)
}
