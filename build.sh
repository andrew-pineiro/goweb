#!/bin/bash
set -xe

APPDIR = "/www"
DOCKNAME = "goweb_1"
IMGNAME = "goweb"

if [ -d $APPDIR]; then
    rm $APPDIR -r
fi
mkdir $APPDIR
cp "./*" $APPDIR -r

docker build --tag $IMGNAME $APPDIR

if [ ! `$(docker ps -a -q -f name=$DOCKNAME)` ]; then
    if [ `$(docker ps -aq -f status=exited -f name=$DOCKNAME)` ]; then
        docker rm $DOCKNAME
    fi
    docker run --name $DOCKNAME -p 80:8080 $IMGNAME:latest
fi
