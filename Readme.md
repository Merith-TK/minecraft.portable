## TO USE
Simply download the Release.exe file and run it. 
It will download the latest `Minecraft.exe` launcher and run it, storing all data in `MinecraftData`.

## Advanced
If you wish to use a custom launcher, or even another program entirely, you'll need to edit the 
`toml` file in the same folder as the program, below is the default config, without the comments
that are inplace to describe each one, 

The Launcher will assume the program is inside the `MinecraftData` folder unless you specify a 
filepath like so `/Portable/Program/program.exe`, which is `DRIVE:\Portable\Program\program.exe`, 
the `MinecraftData` folder will be named to whatever you name your EXE

```toml
launcher = "minecraft.exe"
launcherArgs = ""

[java]
  javaArgs = ""
  useJava16 = false
  usePortableJava = false

[environment]
  APPDATA = "./"
  HOME = "./"
```
