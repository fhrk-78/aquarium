$env:GOOS = "windows"
$env:GOARCH = "amd64"
go build -o windows/amd64/aqrbuild.exe
$env:GOOS = "windows"
$env:GOARCH = "arm64"
go build -o windows/arm64/aqrbuild.exe
$env:GOOS = "windows"
$env:GOARCH = "arm"
go build -o windows/arm/aqrbuild.exe
$env:GOOS = "linux"
$env:GOARCH = "amd64"
go build -o linux/amd64/aqrbuild
$env:GOOS = "linux"
$env:GOARCH = "arm64"
go build -o linux/arm64/aqrbuild
$env:GOOS = "linux"
$env:GOARCH = "arm"
go build -o linux/arm/aqrbuild
