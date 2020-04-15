#!/usr/bin/env bash

#cpath=`pwd`
#PROJECT_PATH=${cpath%src*}
#echo $PROJECT_PATH
#export GOPATH=$GOPATH:${PROJECT_PATH}

SOURCE_FILE_NAME=main
TARGET_FILE_NAME=reskd

build(){
    echo 'start build:' $GOOS $GOARCH
    tname=${TARGET_FILE_NAME}_${GOOS}_${GOARCH}${EXT}
    env GOOS=$GOOS GOARCH=$GOARCH \
    go build -o ${tname} \
    -v
    #${SOURCE_FILE_NAME}.go 这边不指定main.go了，会报一个警告，不加貌似就是一次性编译此文件夹里的所有文件
    chmod +x ${tname}
}

#macOS 64
GOOS=darwin
GOARCH=amd64
build

#linux 64
GOOS=linux
GOARCH=amd64
build

#windows 64
GOOS=windows
GOARCH=amd64
EXT='.exe'
build
#32位系统的
GOARCH=386
build