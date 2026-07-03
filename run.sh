#!/bin/bash

echo "Building MediaGo WebUI..."
cd "$(dirname "$0")"
go build -o mediago-webui server/main.go

if [ $? -eq 0 ]; then
    echo "Build successful!"
    echo "Starting server..."
    ./mediago-webui
else
    echo "Build failed!"
    exit 1
fi
