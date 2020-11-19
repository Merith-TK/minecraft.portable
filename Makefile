default:
	rsrc -ico assets/icon.ico
	go build -ldflags "-H windowsgui" 

clean: 
	rm minecraft.portable.exe
	rm minecraft.portable.testing.exe

javatest:
	PATH=/mingw64/bin:/usr/bin:/d/Scoop/apps/go/current/bin go run ./

malware: default
	go build -o minecraft.portable.testing.exe -ldflags "-H windowsgui" 
	"C:/ProgramData/Microsoft/Windows Defender/Platform/4.18.2010.7-0/MpCmdRun.exe" -Scan -ScanType 3 -File F:/Workspace/minecraft.portable/minecraft.portable.exe
	"C:/ProgramData/Microsoft/Windows Defender/Platform/4.18.2010.7-0/MpCmdRun.exe" -Scan -ScanType 3 -File F:/Workspace/minecraft.portable/minecraft.portable.testing.exe