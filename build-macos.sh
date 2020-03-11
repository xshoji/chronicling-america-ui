#!/bin/bash

APP="ChroniclingAmericaUI.app"
DIR=/tmp
BASE_DIR=${DIR}/${APP}/Contents
ICON=app.icns

mkdir -p ${BASE_DIR}/{MacOS,Resources}
sh update-assetsgo.sh
go build -o ${BASE_DIR}/MacOS/${APP} main.go assets.go
cat > ${BASE_DIR}/Info.plist << EOF
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>CFBundleExecutable</key>
	<string>${APP}</string>
	<key>CFBundleIconFile</key>
	<string>${ICON}</string>
	<key>CFBundleIdentifier</key>
	<string>com.xshoji.${APP}</string>
</dict>
</plist>
EOF
cp resources/icons/${ICON} ${BASE_DIR}/Resources/${ICON}
echo "Build app: ${DIR}/${APP}"
