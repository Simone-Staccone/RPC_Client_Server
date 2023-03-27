package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func (x Ret) AddResource(resource *Resource, reply *Reply) error {
	var oldOverlay OverlayNetwork
	f, err := os.Open("data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	bytes, _ := ioutil.ReadFile("./data.json")

	err = json.Unmarshal(bytes, &(oldOverlay))

	if err != nil {
		log.Fatal(err)
	}

	overlay := oldOverlay
	overlay.Network[resource.Id].Value = resource.Value

	e, _ := json.Marshal(overlay)

	err = ioutil.WriteFile("./data.json", e, 0)
	if err != nil {
		log.Fatal(err)
	}

	reply.RET = 1
	return nil
}

func (x Ret) LookUpResource(resource *Resource, reply *Reply) error {
	var overlay OverlayNetwork
	f, err := os.Open("data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	bytes, _ := ioutil.ReadFile("./data.json")

	err = json.Unmarshal(bytes, &(overlay))

	if err != nil {
		log.Fatal(err)
	}

	reply.RET = -1

	var i int
	for i = 0; i < len(overlay.Network); i++ {
		if strings.Compare(overlay.Network[i].Value, resource.Value) == 0 {
			reply.RET = overlay.Network[i].Id
			break
		}
	}
	return nil
}
