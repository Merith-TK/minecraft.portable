package main

import (
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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
		log.Fatal("[MineCraftPortable]: ERROR, TechnicLauncher.jar Not found. Probably failed to download")
	}
	var java string
	var err error
	if !conf.Java.UsePortableJava {
		java, err = exec.LookPath("java")
		if err != nil {
			log.Fatalln("ERROR: NO JAVA INSTALLED, Using Portable Runtime")
			java = locateJava()
		}
	} else {
		java = locateJava()
	}
	err = nil
	java = filepath.ToSlash(java)
	javaPath := filepath.ToSlash(filepath.Dir(java))
	os.Setenv("PATH", javaPath)
	cmd := exec.Command(javaPath+"/javaw.exe", "-jar", dataDir+"/TechnicLauncher.jar")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Stdin = os.Stdin
	log.Println("[MineCraftPortable] Running TechnicLauncher.exe")
	log.Println("[MineCraftPortable] Launcher will start Shortly")
	err = cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

func technicVerify() {
	err := getjson("https://api.technicpack.net/launcher/version/stable4", &technicapi)
	if err != nil {
		log.Println("Failed to PARSE json\n", err)
		log.Println(technicapi)
		os.Exit(1)
	}
	download(dataDir+"/TechnicLauncher.jar", technicapi.URL.Jar)
}
