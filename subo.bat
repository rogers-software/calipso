git add .
git commit -m "Ultimo Commit"
git push
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
go build -tags lambda.norpc -o main.go
rm main.zip
tar -a -cf main.zip main bootstrap
