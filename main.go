package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	log.Println("[Main] Starting")
	if _, err := os.Stat(dataDir); err != nil {
		log.Println("[Main] Creating data directory")
		os.Mkdir(dataDir, 0755)
	}
	log.Println("[Main] Loading config")
	err := setupConfig()
	if err != nil {
		log.Println("[Main] [Error] Could not init config")
		log.Println(err)
		os.Exit(1)
	}

	if !conf.Java.UseJava {
		if conf.Launcher == "minecraft.exe" {
			minecraftexe()
		} else {
			os.Setenv("APPDATA", dataDir+"/")
			os.Setenv("HOME", dataDir+"/")
			if strings.Contains(strings.ToLower(conf.Launcher), "technic") {
				log.Println("[Technic] Launching Technic")
				technic()
			} else {
				unknownexe(conf.Launcher, conf.LauncherArgs)
			}
		}
	}
	if conf.Java.UseJava {
		log.Println("running java")
		javaexe(conf.Launcher)
	}
}
