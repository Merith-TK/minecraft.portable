default:
	go build -o MineCraft.portable.exe -ldflags "-H windowsgui" 
	rcedit MineCraft.portable.exe \
	--set-icon assets/icon.ico \
	--set-version-string "FileDescription" "MineCraft Portable" \
	--set-version-string "ProductName" "MineCraft Portable" \
	--set-version-string "LegalCopyright" "Merith.TK [PortableLauncher]"

	GOOS=linux go build -o MineCraft.portable.linux.amd64

clean: 
	rm MineCraft.portable.exe
	rm MineCraft.portable.linux.amd64

javatest:
	PATH=/mingw64/bin:/usr/bin:/d/Scoop/apps/go/current/bin go run ./
