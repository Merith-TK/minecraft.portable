package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	//"github.com/gen2brain/dlgs"
)

var (
	conf       config
	configfile = "MinecraftData" + "/config.portable.json"
)

type config struct {
	Launcher     string `json:"launcher"`
	Java         bool   `json:"java"`
	JavaPortable bool   `json:"javaPortable"`
	Args         string `json:"args"`
}

func main() {
	if _, err := os.Stat("MinecraftData"); err != nil {
		os.Mkdir("MinecraftData", 0755)
	}

	os.Setenv("APPDATA", "./")
	os.Setenv("HOME", "./")
	launcher, java, portable := readjson()

	if portable {
		path := os.Getenv("PATH")
		os.Setenv("PATH", portableJavaPath+path)
	}

	if !java {
		if launcher == "minecraft.exe" {
			minecraftexe()
		} else {
			os.Setenv("APPDATA", "MinecraftData/")
			os.Setenv("HOME", "MinecraftData")
			unknownexe(launcher, conf.Args)
		}
	} else {
		os.Setenv("APPDATA", "MinecraftData/")
		os.Setenv("HOME", "MinecraftData")
		javaexe(launcher)
	}

}

func createConfig() (string, bool, bool) {
	log.Println("[MineCraftPortable]", "This application takes a bit to work on the first run, please be patient")
	file, _ := os.Create(configfile)
	defer file.Close()

	_, _ = io.WriteString(file, `{"launcher":"minecraft.exe","java":false,"javaPortable":false,"args":""}`)
	file.Sync()
	return readjson()
}

func readjson() (string, bool, bool) {
	str, err := ioutil.ReadFile(configfile)
	if err != nil {
		fmt.Println("Creating Config")
		return createConfig()
	}
	err = json.Unmarshal([]byte(str), &conf)

	return conf.Launcher, conf.Java, conf.JavaPortable
}
