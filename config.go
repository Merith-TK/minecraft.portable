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
		UseJava17       bool   `toml:"useJava17"`
		UsePortableJava bool   `toml:"usePortableJava"`
	} `toml:"java"`
	Environment map[string]string `toml:"environment"`
}

func setupConfig() error {
	if _, err := os.Stat(configfile); os.IsNotExist(err) {
		fmt.Println("[MineCraftPortable] No config found, creating default config")
		f, err := os.Create(configfile)
		if err != nil {
			return err
		}
		toml.NewEncoder(f).Encode(conf)
		// write config file
		f.Close()
	} else {
		fmt.Println("[MineCraftPortable] Found config, loading config")
		_, err := toml.DecodeFile(configfile, &conf)
		if err != nil {
			return err
		}
	}

	return nil
}
