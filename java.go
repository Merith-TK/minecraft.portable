package main

import (
	"fmt"
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
		"/PortableApps/CommonFiles/Java64/bin/java.exe", // https://portableapps.com/apps/utilities/java_portable_64
	}
	java17Paths = []string{
		dataDir + "/runtime/java-runtime-beta/windows-x64/java-runtime-beta/bin/javaw.exe",
		"/PortableApps/CommonFiles/OpenJDKJRE64/bin/java.exe", // https://portableapps.com/apps/utilities/OpenJDK64
	}
)

func javaexe(jarfile string) {
	filecheck(jarfile)
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
	fmt.Println("[MineCraftPortable] Running Launcher")
	fmt.Println("[MineCraftPortable] Java Launcher will start Shortly")
	fmt.Println("[MineCraftPortable] Running "+jarfile, conf.Java.JavaArgs)
	fmt.Println("[MineCraftPortable] Launcher will start Shortly")
	cmd.Run()
}

func locateJava() string {
	log.Println("[Java]: Locating JavaRuntime")
	log.Println("[Java]: PortableRuntime", conf.Java.UsePortableJava)
	log.Println("[Java] Searching for Java")
	var javaPath string
	if conf.Java.UsePortableJava {
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
	} else {
		java, err := exec.LookPath("java")
		if err != nil {
			log.Fatalln("[Java]: NO JAVA INSTALLED")
		}
		javaPath = java
	}
	if conf.Java.UsePortableJava {
		javaFile := filepath.ToSlash(javaPath)
		javaBin := filepath.ToSlash(filepath.Dir(javaFile))
		fmt.Println("[MineCraftPortable] Java Path: " + javaBin)
		os.Setenv("PATH", javaBin)
	}
	return javaPath
}
