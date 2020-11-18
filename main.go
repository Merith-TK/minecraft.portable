package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/gen2brain/dlgs"
)

var (
	conf       Config
	configfile = "MinecraftData" + "/config.portable.json"
)

type Config struct {
	Launcher string `json:"launcher"`
	Java     bool   `json:"java"`
}

func main() {
	if _, err := os.Stat("MinecraftData"); err != nil {
		os.Mkdir("MinecraftData", 755)
	}

	os.Setenv("MinecraftData", "./")
	os.Setenv("HOME", "./")
	launcher, java := readjson()

	if java == false {
		if launcher == "minecraft.exe" {
			minecraftexe()
		} else {
			unknownexe(launcher)
		}
	} else {
		javaexe(launcher)
	}

}

func createConfig() {
	dlgs.Warning("[MineCraftPortable]", "This application takes a bit to work on the first run, please be patient")
	file, _ := os.Create(configfile)
	defer file.Close()

	_, _ = io.WriteString(file, `{"launcher":"minecraft.exe","java":false}`)
	file.Sync()
	main()
}

func readjson() (string, bool) {
	str, err := ioutil.ReadFile(configfile)
	if err != nil {
		fmt.Println("Creating Config")
		createConfig()
	}
	err = json.Unmarshal([]byte(str), &conf)

	return conf.Launcher, conf.Java
}
