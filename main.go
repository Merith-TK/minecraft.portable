package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	//"github.com/gen2brain/dlgs"
)

var (
	conf       config
	configfile = "MinecraftData" + "/config.portable.json"
)

type config struct {
	Launcher     string            `json:"launcher"`
	Java         bool              `json:"java"`
	JavaPortable bool              `json:"javaPortable"`
	Args         string            `json:"args"`
	Environment  map[string]string `json:"environment"`
}

func main() {
	if _, err := os.Stat("MinecraftData"); err != nil {
		os.Mkdir("MinecraftData", 0755)
	}

	for k, v := range conf.Environment {
		os.Setenv(k, v)
		fmt.Println("ENV:", k, "=", v)
	}
	launcher, java, portable := readjson()

	if portable {
		pwd, _ := os.Getwd()
		path := pwd + "\\" + portableJavaPath
		path = strings.ReplaceAll(path, "/", "\\")
		os.Setenv("PATH", path)
	}
	if !java {
		if launcher == "minecraft.exe" {
			minecraftexe()
		} else {
			os.Setenv("APPDATA", "MinecraftData/")
			os.Setenv("HOME", "MinecraftData")
			if strings.Contains(strings.ToLower(launcher), "technic") {
				fmt.Println("Launching Technic")
				technic(launcher, java)
			} else {
				unknownexe(launcher, conf.Args)
			}
		}
	} else {
		javaexe(launcher)
	}

}

func createConfig() (string, bool, bool) {
	log.Println("[MineCraftPortable]", "This application takes a bit to work on the first run, please be patient")
	file, _ := os.Create(configfile)
	defer file.Close()

	defaultConf := `{
	"launcher":"minecraft.exe",
	"java":false,
	"javaPortable":false,
	"args":"",
	"environment":{
		"APPDATA":"./",
		"HOME":"./"
	}
}`
	_, _ = io.WriteString(file, defaultConf)
	file.Sync()
	return readjson()
}

func readjson() (string, bool, bool) {
	str, err := ioutil.ReadFile(configfile)
	if err != nil {
		fmt.Println("Creating Config")
		return createConfig()
	}
	_ = json.Unmarshal([]byte(str), &conf)

	return conf.Launcher, conf.Java, conf.JavaPortable
}
