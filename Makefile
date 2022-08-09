default:
	rsrc -ico assets/icon.ico
	go build

clean: 
	-rm minecraft.portable.exe
	-rm minecraft.portable.testing.exe

javatest: clean
	go build
	export PATH=/mingw64/bin:/usr/bin 
	minecraft.portable.exe

malware:
	go build -o minecraft.portable.testing.exe
	"C:/ProgramData/Microsoft/Windows Defender/Platform/4.18.2106.6-0/MpCmdRun.exe" -Scan -ScanType 3 -File D:/Workspace/minecraft.portable/minecraft.portable.testing.exe

fix-repo:
	git remote add github https://github.com/merith-tk/minecraft.portable.git
	git remote add gitea https://git.merith.tk/merith-tk/minecraft.portable.git
	git push --set-upstream github master
	git push --set-upstream gitea master
push:
	git push github
	git push gitea
	