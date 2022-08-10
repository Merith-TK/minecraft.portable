package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

/*
	### This is where I define how to setup technic locally
*/

type technicApi struct {
	Build string `json:"build"`
	URL   struct {
		Exe string `json:"exe"`
		Jar string `json:"jar"`
		Osx string `json:"osx"`
	} `json:"url"`
	Resources []struct {
		Filename string `json:"filename"`
		URL      string `json:"url"`
		Md5      string `json:"md5"`
	} `json:"resources"`
}

var (
	technicapi        technicApi
	technicConfigFile = dataDir + "/.technic/settings.json"
)

func technic() {
	if _, err := os.Stat(dataDir + "/.technic/settings.json"); err != nil {
		os.MkdirAll(dataDir+"/.technic", 0755)
		technicConfig()
	}
	technicexe()
}

func technicConfig() {
	file, _ := os.Create(technicConfigFile)
	defer file.Close()
	_, _ = io.WriteString(file, `{
		"memory": 0,
		"launchAction": "HIDE",
		"buildStream": "stable",
		"showConsole": false,
		"languageCode": "default",
		"clientId": "",
		"latestNewsArticle": 162,
		"launchToModpacks": false,
		"javaVersion": "default",
		"autoAcceptRequirements": true,
		"javaBitness": true,
		"launcherSettingsVersion": "1",
		"windowType": "DEFAULT",
		"windowWidth": 0,
		"windowHeight": 0,
		"enableStencilBuffer": true
}`)
	file.Sync()
}

func technicexe() {
	technicVerify()
	if _, err := os.Stat(dataDir + "/TechnicLauncher.jar"); err != nil {
		log.Fatal("[MineCraftPortable]: ERROR, TechnicLauncher.exe Not found. Probably failed to download")
	}

	// if technic/settings.json doesn't exist, create it
	if _, err := os.Stat(dataDir + "/technic/settings.json"); err != nil {
		os.MkdirAll(dataDir+"/technic", 0755)
		technicDefualtConfig := `{
	"memory": 0,
	"launchAction": "HIDE",
	"buildStream": "stable",
	"showConsole": true,
	"languageCode": "default",
	"clientId": "",
	"directory": "portable",
	"latestNewsArticle": 0,
	"launchToModpacks": false,
	"javaVersion": "default",
	"autoAcceptRequirements": false,
	"javaBitness": true,
	"launcherSettingsVersion": "2",
	"windowType": "DEFAULT",
	"windowWidth": 0,
	"windowHeight": 0,
	"enableStencilBuffer": true,
	"useMojangJava": false
}`
		file, _ := os.Create(dataDir + "/technic/settings.json")
		defer file.Close()
		_, _ = io.WriteString(file, technicDefualtConfig)
		file.Sync()
	}
	javaexe("TechnicLauncher.jar")
}

func technicVerify() {
	err := getjson("https://api.technicpack.net/launcher/version/stable4", &technicapi)
	if err != nil {
		fmt.Println("Failed to PARSE json\n", err)
		fmt.Println(technicapi)
		os.Exit(1)
	}
	download(dataDir+"/TechnicLauncher.jar", technicapi.URL.Jar)
}
