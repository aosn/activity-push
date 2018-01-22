set GOOS=linux
go build -o main .
..\..\..\..\..\..\bin\build-lambda-zip.exe -o ..\..\main.zip main
del main
set GOOS=windows
