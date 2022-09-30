package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

var (
	java8Paths = []string{
		dataDir + "/runtime/jre-legacy/windows-x64/jre-legacy/bin/javaw.exe",
		"/PortableApps/CommonFiles/Java64/bin/javaw.exe", // https://portableapps.com/apps/utilities/java_portable_64
	}
	java17Paths = []string{
		"/PortableApps/CommonFiles/OpenJDKJRE64/bin/javaw.exe", // https://portableapps.com/apps/utilities/OpenJDK64
	}
)

func javaexe(jarfile string) {
	var java string
	var err error

	if strings.HasPrefix(jarfile, "/") || strings.HasPrefix(jarfile, "./") {
		jarfile, _ = filepath.Abs(jarfile)
	} else {
		jarfile, _ = filepath.Abs(dataDir + "/" + jarfile)
	}

	if !conf.Java.UsePortableJava {
		java, err = exec.LookPath("java")
		if err != nil {
			log.Fatalln("ERROR: NO JAVA INSTALLED, Using Portable Runtime")
			java = locateJava()
		}
	} else {
		java = locateJava()
	}
	if java == "" {
		log.Fatalln("ERROR: NO JAVA FOUND, Please run minecraft atleast once through the default launcher")
		time.Sleep(20 * time.Second)
	} else {
		log.Println(java)
	}
	cmd := exec.Command(java, "-jar", filepath.Base(jarfile), conf.Java.JavaArgs)
	cmd.Dir = filepath.Dir(jarfile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Stdin = os.Stdin
	log.Println("[MineCraftPortable] Running Launcher")
	log.Println("[MineCraftPortable] Java Launcher will start Shortly")
	//cmd.Run()

	log.Println("[MineCraftPortable] DID NOT RUN, IN DEV MODE")
}

func locateJava() string {
	var javaPath string
	log.Println("[MineCraftPortable] Searching for Java")
	log.Println("[MineCraftPortable] Use Java17:", conf.Java.UseJava17)
	if conf.Java.UseJava17 {
		for _, path := range java17Paths {
			log.Println("LOCATE:", path)
			if _, err := os.Stat(path); err == nil {
				javaPath = path
				break
			}
		}
	} else {
		for _, path := range java8Paths {
			java, _ := filepath.Abs(path)
			if _, err := os.Stat(java); err == nil {
				javaPath = java
				break
			}
		}
	}
	return javaPath
}
