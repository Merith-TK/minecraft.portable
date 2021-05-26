package main

import (
	"log"
	"os"
	"strings"

	toml "github.com/pelletier/go-toml"
)

var (
	conf          config
	configfile    = strings.TrimSuffix(os.Args[0], "exe") + "toml"
	defaultConfig = `
# launcher
#	the file to launch within the MinecraftData folder
#	if the file is located in MinecraftData/launcher.exe,
#	input "launcher.exe
# launcherArgs
#	Arguments you need to run the launcher, only useful if
#	the launcher is not the official launcher from microsoft
launcher = "minecraft.exe"
launcherArgs = ""

# javaArgs
#	Arguments to pass to java
# useJava
#	use java, this will be removed in later
#	updates
# useJava16
#	This is purely to use for later snapshots, 
#	furture proofing, only usable with 
# usePortableJava
#	Use the portable java bundled with the minecraft launcher
#	This currently only supports java8, snapshots after '21w19a'
#	will not work in third party launchers for anything that is NOT
#	a snapshot
[java]
  javaArgs = ""
  useJava = false
  useJava16 = false
  usePortableJava = false

# Use only if you need to set custom variables, usually not needed
# useful for non-minecraft applications,
[environment]
  APPDATA = "./"
  HOME = "./"
`
)

type config struct {
	Launcher     string `toml:"launcher"`
	LauncherArgs string `toml:"launcherArgs"`
	Java         struct {
		JavaArgs        string `toml:"javaArgs"`
		UseJava         bool   `toml:"useJava"`
		UseJava16       bool   `toml:"useJava16"`
		UsePortableJava bool   `toml:"usePortableJava"`
	} `toml:"java"`
	Environment map[string]string `toml:"environment"`
}

func setupConfig() error {
	str, err := os.ReadFile(configfile)
	if err != nil {
		f, err := os.OpenFile(configfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		defer f.Close()
		if _, err := f.WriteString(defaultConfig); err != nil {
			log.Println(err)
		}
		str, _ = os.ReadFile(configfile)
		err = nil
	}
	_ = toml.Unmarshal([]byte(str), &conf)
	return err
}
