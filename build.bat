echo off
set CGO_ENABLED=0

echo Setting Architecture: amd64
set GOARCH=amd64

echo Setting OS: windows
set GOOS=windows
echo Building: GO-CONNECTION-MONITOR for Windows
go build -o .\.dist\conn-monitor.exe ./cmd/conn_monitor
7z a -tzip .\.dist\go-monitor-connection--windows-amd64--%*.zip .\.dist\conn-monitor.exe readme.md


echo Setting OS: linux
set GOOS=linux

echo Building: GO-CONNECTION-MONITOR for Linux
go build -o .\.dist\conn-monitor ./cmd/conn_monitor
7z a -tzip .\.dist\go-monitor-connection--linux-amd64--%*.zip .\.dist\conn-monitor readme.md

copy .\.dist\conn-monitor.exe .\conn-monitor.exe