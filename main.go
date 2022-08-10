package main

import (
	"fmt"
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
		fmt.Println("Could not init config")
		log.Println(err)
	}
	for k, v := range conf.Environment {
		os.Setenv(k, v)
		fmt.Println("ENV:", k, "=", v)
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
			minecraftexe()
			return
		}
		if conf.Launcher == "technic.exe" {
			technicexe()
			return
		}
	}

}
