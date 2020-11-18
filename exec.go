package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/gen2brain/dlgs"
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

func unknownexe(execute string) {
	filecheck(execute)
	if _, err := os.Stat("MinecraftData/" + execute); err != nil {
		dlgs.Error("[MineCraftPortable]: ERROR", execute+" not found, did you edit config.portable.json?")
	}
	cmd := exec.Command("MinecraftData/" + execute)
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
	cmd := exec.Command(java, "-version")
	if err := cmd.Run(); err != nil {
		java = "MinecraftData/runtime/jre-x64/bin/java.exe"
	}
	cmd = exec.Command(java, "-jar", jarfile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Stdin = os.Stdin
	fmt.Println("[MineCraftPortable] Running Launcher")
	fmt.Println("[MineCraftPortable] MineCraft will start Shortly")
	cmd.Run()
}
