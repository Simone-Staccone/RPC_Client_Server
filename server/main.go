package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func initializeServer() {
	_, err := os.Stat("./data.json")

	if err != nil {
		os.Create("./data.json")
		overlay := new(OverlayNetwork)
		for i := 0; i < 256; i++ {
			overlay.Network[i].Id = i

			//_AddNode(res, rep)
		}

		f, err := os.Open("./data.json")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		e, _ := json.Marshal(overlay)

		ioutil.WriteFile("./data.json", e, 0)
	}

}

func main() {
	initializeServer()
	Server()
}
