GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./builds/linux_amd64/acaps ./main.go
GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o ./builds/linux_arm64/acaps ./main.go

GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o ./builds/macos_amd64/acaps ./main.go
GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o ./builds/macos_arm64/acaps ./main.go

GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o ./builds/win_amd64/acaps.exe ./main.go
GOOS=windows GOARCH=arm64 go build -ldflags="-s -w" -o ./builds/win_arm64/acaps.exe ./main.go
