package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	if _, err := os.Stat(dataDir); err != nil {
		os.Mkdir(dataDir, 0755)
	}

	err := setupConfig()
	if err != nil {
		log.Println("Could not init config")
		log.Println(err)
	}

	if !conf.Java.UseJava {
		if conf.Launcher == "minecraft.exe" {
			minecraftexe()
		} else {
			os.Setenv("APPDATA", dataDir+"/")
			os.Setenv("HOME", dataDir+"/")
			if strings.Contains(strings.ToLower(conf.Launcher), "technic") {
				log.Println("Launching Technic")
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
