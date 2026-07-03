@echo off
echo Building MediaGo WebUI...
cd /d "%~dp0"
go build -o mediago-webui.exe server/main.go
if %ERRORLEVEL% EQU 0 (
    echo Build successful!
    echo Starting server...
    mediago-webui.exe
) else (
    echo Build failed!
    pause
)
