package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
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
