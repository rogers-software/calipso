git add .
git commit -m "Ultimo Commit"
git push
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
go build -ldflags="-s -w" -o bootstrap main.go
rm main.zip
tar -a -cf main.zip main bootstrap
