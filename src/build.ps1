$env:GOOS = "windows"
$env:GOARCH = "amd64"
go build -o windows/amd64/aqr.exe
$env:GOOS = "windows"
$env:GOARCH = "arm64"
go build -o windows/arm64/aqr.exe
$env:GOOS = "windows"
$env:GOARCH = "arm"
go build -o windows/arm/aqr.exe
$env:GOOS = "linux"
$env:GOARCH = "amd64"
go build -o linux/amd64/aqr
$env:GOOS = "linux"
$env:GOARCH = "arm64"
go build -o linux/arm64/aqr
$env:GOOS = "linux"
$env:GOARCH = "arm"
go build -o linux/arm/aqr
