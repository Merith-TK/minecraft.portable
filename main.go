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

	if !conf.Java.UseJava {
		if conf.Launcher == "minecraft.exe" {
			minecraftexe()
		} else {
			os.Setenv("APPDATA", dataDir+"/")
			os.Setenv("HOME", dataDir+"/")
			if strings.Contains(strings.ToLower(conf.Launcher), "technic") {
				fmt.Println("Launching Technic")
				technic()
			} else {
				unknownexe(conf.Launcher, conf.LauncherArgs)
			}
		}
	}
	if conf.Java.UseJava {
		fmt.Println("running java")
		javaexe(conf.Launcher)
	}
}
