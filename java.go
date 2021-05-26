package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

var (
	javaPaths = []string{
		"MinecraftData/runtime/jre-legacy/windows-x64/jre-legacy/bin/javaw.exe",
		"/PortableApps/CommonFiles/Java64/bin/javaw.exe",
		"/PortableApps/CommonFiles/Java/bin/javaw.exe",
	}
)

func javaexe(jarfile string) {
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
	if conf.Java.UseJava16 {
		java, err = filepath.Abs("MinecraftData/runtime/jre-runtime-alpha/windows-x64/jre-runtime-alpha/bin/javaw.exe")
		if err != nil {
			if _, err := os.Stat(java); err != nil {
				log.Fatalln("ERROR: Java 16 not found, did you run Minecraft Snapshot 21w19a or later?")
				time.Sleep(20 * time.Second)
				os.Exit(2)
			}
		}
	}
	if java == "" {
		log.Fatalln("ERROR: NO JAVA FOUND, Please run minecraft atleast once through the default launcher")
		time.Sleep(20 * time.Second)
	} else {
		fmt.Println(java)
	}
	cmd := exec.Command(java, "-jar", filepath.Base(jarfile), conf.Java.JavaArgs)
	cmd.Dir = filepath.Dir(jarfile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Stdin = os.Stdin
	fmt.Println("[MineCraftPortable] Running Launcher")
	fmt.Println("[MineCraftPortable] Java Launcher will start Shortly")
	cmd.Run()
}

func locateJava() string {
	for _, p := range javaPaths {
		java, _ := filepath.Abs(p)
		fmt.Println("LOCATE:", java)
		if _, err := os.Stat(java); err == nil {
			return java
		}
	}
	return ""
}
