#!/bin/bash
set -xe

APPDIR="/www"
DOCKNAME="goweb_1"
IMGNAME="goweb"
SVCNAME="weblogs.service"

if [ -d $APPDIR ]; then
    rm $APPDIR -r
fi
mkdir $APPDIR
cp ./ $APPDIR -r

docker build --tag $IMGNAME $APPDIR

if [ $( docker ps -a -f name=$DOCKNAME| wc -l ) -eq 2 ]; then  
  docker rm -f $DOCKNAME
fi

docker run -d --name $DOCKNAME -p 80:8080 $IMGNAME:latest
systemctl restart $SVCNAME 