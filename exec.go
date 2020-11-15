package main

import (
	"fmt"
	"os"
	"os/exec"
)

/*

	### THIS IS WHERE MY COMMANDS TO LAUNCH THE GAME ARE STORED ###

*/

var (
	exeMinecraft = "https://launcher.mojang.com/download/Minecraft.exe"

	// More Reliable link to the jar...
	jarMinecraft = "https://cdn.merith.tk/_Releases/Other/OfficialMinecraft.jar"
)

func minecraftexe() {
	filecheck("MinecraftData/minecraft.exe", exeMinecraft)
	cmd := exec.Command("MinecraftData/minecraft.exe", "--workDir", ".minecraft")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Stdin = os.Stdin
	fmt.Println("[MineCraftPortable] Running Launcher")
	fmt.Println("[MineCraftPortable] MineCraft will start Shortly")
	cmd.Run()
}

func unknownexe(execute string) {
	cmd := exec.Command("MinecraftData/" + execute)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Stdin = os.Stdin
	fmt.Println("[MineCraftPortable] Running Launcher")
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
