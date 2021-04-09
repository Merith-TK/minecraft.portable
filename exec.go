package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

/*

	### THIS IS WHERE MY COMMANDS TO LAUNCH THE GAME ARE STORED ###

*/

var (
	portableJavaPath = "MinecraftData/runtime/jre-legacy/windows-x64/jre-legacy/bin/;"
	portableJava     = "MinecraftData/runtime/jre-legacy/windows-x64/jre-legacy/bin/java.exe"
)

func minecraftexe() {
	filecheck("minecraft.exe")
	cmd := exec.Command("MinecraftData/minecraft.exe", "--workDir", ".minecraft")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Stdin = os.Stdin
	fmt.Println("[MineCraftPortable] Running Launcher")
	fmt.Println("[MineCraftPortable] MineCraft will start Shortly")
	cmd.Run()
}

func unknownexe(execute string, args string) {
	filecheck(execute)
	if _, err := os.Stat("MinecraftData/" + execute); err != nil {
		log.Fatal("[MineCraftPortable]: ERROR", execute+" not found, did you edit config.portable.json?")
	}
	cmdargs := strings.Split(args, " ")
	cmd := exec.Command("MinecraftData/"+execute, cmdargs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Stdin = os.Stdin
	fmt.Println("[MineCraftPortable] Running " + execute)
	fmt.Println("[MineCraftPortable] Launcher will start Shortly")
	cmd.Run()
}

func javaexe(jarfile string) {
	jarfile = "MinecraftData/" + jarfile
	java := "java"
	if conf.JavaPortable {
		cmd := exec.Command(java, "-version")
		if err := cmd.Run(); err != nil {
			java = portableJava
		}
	}
	cmd := exec.Command(java, "-jar", filepath.Base(jarfile), conf.Args)
	cmd.Dir = filepath.Dir(jarfile)
	fmt.Println(cmd.Dir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Stdin = os.Stdin
	fmt.Println("[MineCraftPortable] Running Launcher")
	fmt.Println("[MineCraftPortable] MineCraft will start Shortly")
	cmd.Run()
}
