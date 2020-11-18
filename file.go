package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

/*

	### This is for Misc File operations ###

*/

var (
	rel Release
	URL = "https://cdn.merith.tk/_Releases/Other/minecraft.json"
)

type Release struct {
	Official   map[string]string `json:"official"`
	Unofficial map[string]string `json:"unofficial"`
}

func filecheck(filename string) {
	err := getjson(URL, &rel)
	if err != nil {
		fmt.Println("Failed to PARSE json\n", err)
		fmt.Println(rel)
		os.Exit(1)
	}
	for k, v := range rel.Official {
		if filename == k {
			if _, err := os.Stat("MinecraftData/" + k); err != nil {
				download("MinecraftData/"+k, v)
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
		return errors.New("Received non 200 response code")
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

func download(fileName string, URL string) error {
	//Get the response bytes from the url
	fmt.Println(fileName, "\n", URL)
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		response, err := http.Get(URL)
		if err != nil {
			return err
		}
		defer response.Body.Close()

		if response.StatusCode != 200 {
			return errors.New("Received non 200 response code")
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

func md5hash(filePath string) (string, error) {
	//Initialize variable returnMD5String now in case an error has to be returned
	var returnMD5String string

	//Open the passed argument and check for any error
	file, err := os.Open(filePath)
	if err != nil {
		return returnMD5String, err
	}

	//Tell the program to call the following function when the current function returns
	defer file.Close()

	//Open a new hash interface to write to
	hash := md5.New()

	//Copy the file in the hash interface and check for any error
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}

	//Get the 16 bytes hash
	hashInBytes := hash.Sum(nil)[:16]

	//Convert the bytes to a string
	returnMD5String = hex.EncodeToString(hashInBytes)

	return returnMD5String, nil

}

func fileWrite(filename string, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}
