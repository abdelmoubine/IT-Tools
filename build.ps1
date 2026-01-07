$env:CGO_ENABLED = "1"
$env:GOOS = "windows"
$env:GOARCH = "amd64"
go build -o ittool_win64.exe main.go
$env:GOARCH = "386"
go build -o ittool_win32.exe main.go
Write-Host "Build complete."