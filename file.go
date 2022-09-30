package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
)

/*

	### This is for Misc File operations ###

*/

var (
	rel release
)

type release struct {
	Official   map[string]string `json:"official"`
	Unofficial map[string]string `json:"unofficial"`
}

func filecheck(filename string) {
	err := getjson("https://raw.githubusercontent.com/Merith-TK/minecraft.portable/master/assets/launchers.json", &rel)
	if err != nil {
		log.Println("Failed to PARSE json\n", err)
		log.Println(rel)
		os.Exit(1)
	}
	for k, v := range rel.Official {
		if filename == k {
			if _, err := os.Stat(dataDir + "/" + k); err != nil {
				download(dataDir+"/"+k, v)
			}
		}
	}
}

func getjson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	if r.StatusCode != 200 {
		return errors.New("received non 200 response code")
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

func download(fileName string, URL string) error {
	//Get the response bytes from the url
	log.Println(fileName, "\n", URL)
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		response, err := http.Get(URL)
		if err != nil {
			return err
		}
		defer response.Body.Close()

		if response.StatusCode != 200 {
			return errors.New("received non 200 response code")
		}
		//Create a empty file
		file, err := os.Create(fileName)
		if err != nil {
			return err
		}
		defer file.Close()

		//Write the bytes to the fiel
		_, err = io.Copy(file, response.Body)
		if err != nil {
			return err
		}
	}
	return nil
}
