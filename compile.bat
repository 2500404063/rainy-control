rsrc.exe -manifest nac.manifest -o nac.syso
go build -ldflags="-s -w -H=windowsgui"