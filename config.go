package main

import (
	"fmt"
	"os"
	"strings"

	toml "github.com/BurntSushi/toml"
)

var (
	conf          config
	configfile    = strings.TrimSuffix(os.Args[0], ".exe") + ".toml"
	dataDir       = strings.TrimSuffix(os.Args[0], ".exe") + ".data"
	defaultConfig = `
launcher = "minecraft.exe"
launcherArgs = ""

[java]
  javaArgs = ""
  useJava = false
  useJava17 = false
  usePortableJava = false

# Use only if you need to set custom variables, usually not needed
# useful for non-minecraft applications,
[environment]
  APPDATA = "./"
  HOME = "./"
`

// this is being ignored
)

type config struct {
	Launcher     string `toml:"launcher"`
	LauncherArgs string `toml:"launcherArgs"`
	Java         struct {
		JavaArgs        string `toml:"javaArgs"`
		UseJava         bool   `toml:"useJava"`
		UseJava17       bool   `toml:"useJava17"`
		UsePortableJava bool   `toml:"usePortableJava"`
	} `toml:"java"`
	Environment map[string]string `toml:"environment"`
}

func setupConfig() error {
	if _, err := os.Stat(configfile); os.IsNotExist(err) {
		fmt.Println("[MineCraftPortable] No config found, creating default config")
		fmt.Println("[MineCraftPortable] Default config:", configfile)
		fmt.Println("[MineCraftPortable] Default data:", dataDir)
		f, err := os.Create(configfile)
		if err != nil {
			return err
		}
		toml.NewEncoder(f).Encode(conf)
		// write config file
		f.Close()
	}
	fmt.Println("[MineCraftPortable] Java:", conf.Java.UseJava)
	fmt.Println("[MineCraftPortable] Java17:", conf.Java.UseJava17)
	fmt.Println("[MineCraftPortable] Portable Java:", conf.Java.UsePortableJava)

	return nil
}
