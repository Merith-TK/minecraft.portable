package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	toml "github.com/BurntSushi/toml"
)

var (
	conf       config
	configfile = strings.TrimSuffix(os.Args[0], ".exe") + ".toml"
	dataDir    = strings.TrimSuffix(os.Args[0], ".exe") + ".data"

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

func setDefaultConfig() {
	conf.Launcher = "minecraft.exe"
	conf.LauncherArgs = ""
	conf.Java.JavaArgs = ""
	conf.Java.UseJava = false
	conf.Java.UseJava17 = false
	conf.Java.UsePortableJava = false
	conf.Environment = map[string]string{
		"APPDATA":     "{data}",
		"USERPROFILE": "{data}",
		"HOME":        "{data}",
	}
}

func setupConfig() error {
	if _, err := os.Stat(configfile); os.IsNotExist(err) {
		log.Println("[MineCraftPortable] No config found, creating default config")
		log.Println("[MineCraftPortable] Default config:", configfile)
		log.Println("[MineCraftPortable] Default data:", dataDir)
		f, err := os.Create(configfile)
		if err != nil {
			return err
		}
		setDefaultConfig()
		toml.NewEncoder(f).Encode(conf)
		// write config file
		f.Close()
	} else {
		log.Println("[MineCraftPortable] Config found, loading config")
		log.Println("[MineCraftPortable] Config:", configfile)
		log.Println("[MineCraftPortable] Data:", dataDir)
		_, err := toml.DecodeFile(configfile, &conf)
		if err != nil {
			return err
		}
	}

	log.Println("[MineCraftPortable] Config:")
	setupEnvironment()
	log.Println(conf)
	return nil
}

func setupEnvironment() {
	// Variables for env replacement
	drivePath, _ := filepath.Abs("/")
	drivePath = filepath.ToSlash(drivePath)
	drivePath = strings.TrimSuffix(drivePath, "/")
	configEnvReplace := map[string]string{
		"{data}":  dataDir,
		"{drive}": drivePath,
	}

	// Replace Normal Config options
	for key, value := range configEnvReplace {
		if strings.Contains(conf.Launcher, key) {
			conf.Launcher = filepath.ToSlash(strings.ReplaceAll(conf.Launcher, key, value))
		}
		if strings.Contains(conf.LauncherArgs, key) {
			conf.LauncherArgs = filepath.ToSlash(strings.ReplaceAll(conf.LauncherArgs, key, value))
		}
	}

	// Replace Environment Variables
	for k, v := range conf.Environment {
		for key, value := range configEnvReplace {
			if strings.Contains(v, key) {
				v = strings.ReplaceAll(v, key, value)
				v = filepath.ToSlash(v)
			}
		}
		log.Println("[MineCraftPortable] Setting Environment Variable:", k, "=", v)
		os.Setenv(k, v)

	}
}
