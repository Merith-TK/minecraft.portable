## TO USE
Simply download the Release.exe file and run it. 
It will download the latest `Minecraft.exe` launcher and run it, storing all data in `MinecraftData`.

Dead simple


## Advanced
You can use a custom launcher, all you have to do is open `MinecraftData/config.portable.json` and change `"launcher":"minecraft.exe"` to where your custom launcher is located, assuming MineraftData is the root folder. so `MultiMC` would be setup like so
```
Minecraft.portable.exe
MinecraftData/
	.minecraft/
	game/
	runtime/
	/MultiMC
		/MultiMC.exe
```

your `config.portable.json` file would look like `{"launcher":"MultiMC/MultiMC.exe","java":false}`

When setting up MultiMC with the setup in this example, it is IMPORTANT to use this path to javaw.exe
`../runtime/jre-x64/bin/javaw.exe`, otherwise Java may not work

If you wanted to use a `jar` version of minecraft, like `minecraft.jar` (can be found [here](https://launcher.mojang.com/mc/launcher/jar/fa896bd4c79d4e9f0d18df43151b549f865a3db6/launcher.jar.lzma), you will need winrar or 7zip to open the `lzma` archive)