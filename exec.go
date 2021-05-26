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
	if strings.HasPrefix(execute, "/") || strings.HasPrefix(execute, "./") || strings.HasPrefix(execute, "../") {
		execute, _ = filepath.Abs(execute)
	} else {
		execute, _ = filepath.Abs("MinecraftData/" + execute)
	}
	if _, err := os.Stat(execute); err != nil {
		log.Fatal("[MineCraftPortable]: ERROR", execute, "not found")
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
