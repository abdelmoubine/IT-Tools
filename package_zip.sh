#!/usr/bin/env bash
APP="ittool"
VER="1.0"
mkdir -p dist
zip -r "dist/${APP}_${VER}_portable_win64.zip" ittool_win64.exe resources ui assets LICENSE README.md config.json
echo "Created dist/${APP}_${VER}_portable_win64.zip"