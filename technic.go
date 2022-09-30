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
	technicapi technicApi
)

func technic() {
	if _, err := os.Stat(dataDir + "/technic/settings.json"); err != nil {
		os.MkdirAll(dataDir+"/technic", 0755)

		// Create default config file for portable mode
		file, _ := os.Create(dataDir + "/technic/settings.json")
		defer file.Close()
		_, _ = io.WriteString(file, `{"directory": "portable"}`)
		file.Sync()
	}
	technicexe()
}

func technicexe() {
	technicVerify()
	if _, err := os.Stat(dataDir + "/TechnicLauncher.jar"); err != nil {
		log.Fatal("[MineCraftPortable]: ERROR, TechnicLauncher.jar Not found. Probably failed to download")
	}
	var java string
	var err error
	java = locateJava()
	err = nil
	java = filepath.ToSlash(java)
	javaPath := filepath.ToSlash(filepath.Dir(java))
	os.Setenv("PATH", javaPath)
	cmd := exec.Command(javaPath+"/javaw.exe", "-jar", dataDir+"/TechnicLauncher.jar")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Stdin = os.Stdin
	log.Println("[Technic] Running TechnicLauncher")
	log.Println("[Technic] Launcher will start Shortly")
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
