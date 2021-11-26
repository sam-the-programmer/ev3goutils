@ECHO OFF
ECHO.

COLOR 04
ECHO Setting environment variables...
SET GOOS=linux
SET GOARCH=arm
SET GOARM=5

COLOR 06
ECHO Building application...

ECHO OS = %GOOS%
ECHO Architecture = %GOARCH%
ECHO ARM Version = %GOARM%

ECHO.
CALL go build main.go

COLOR 02
ECHO All Done!