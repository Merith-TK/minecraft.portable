package main

import (
	"fmt"
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

	if conf.Launcher == "" {
		fmt.Println("[MineCraftPortable] No launcher specified, exiting")
		return
	}
	// if conf.Launcher ends with .jar, assume it's a jar file
	if strings.HasSuffix(conf.Launcher, ".jar") {
		javaexe(conf.Launcher)
	} else {
		if conf.Launcher == "minecraft.exe" {
			log.Println("[Main] Minecraft Launcher")
			minecraftexe()
			return
		}
		if strings.ContainsAny(conf.Launcher, strings.ToLower("technic")) {
			log.Println("[Main] Technic Launcher")
			technicexe()
			return
		}
	}

}
