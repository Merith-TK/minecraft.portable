package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

var (
	conf       config
	configfile = "MinecraftData" + "/config.portable.json"
)

type config struct {
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
	fmt.Println(str)
	err = json.Unmarshal([]byte(str), &config)

	return config.Launcher, config.Java
}
