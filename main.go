package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if _, err := os.Stat("MinecraftData"); err != nil {
		os.Mkdir("MinecraftData", 0755)
	}
	fmt.Println(configfile)
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
			os.Setenv("APPDATA", "MinecraftData/")
			os.Setenv("HOME", "MinecraftData")
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
