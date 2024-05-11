#!/bin/bash
set -xe

APPDIR="/www"
DOCKNAME="goweb_1"
IMGNAME="goweb"
SVCNAME="weblogs.service"
SVCPATH="/logging/weblogs.service"
WEBPATH="/etc/website/"
LOGPATH="/etc/website/logs/"
LOGSHPATH="/www/logging/logs.sh"

if [ -d $APPDIR ]; then
    rm $APPDIR -r
fi

mkdir $APPDIR
cp ./ $APPDIR -r

if [ ! -d $WEBPATH ]; then
   mkdir $WEBPATH
   cp $LOGSHPATH $WEBPATH
fi

if [ ! -d $LOGPATH ]; then
    mkdir $LOGPATH
fi

if [ ! service --status-all | grep -Fq $SVCNAME ]; then    
  cp $SVCPATH /etc/systemd/system/
  systemctl enable $SVCNAME
fi


docker build --tag $IMGNAME $APPDIR

if [ $( docker ps -a -f name=$DOCKNAME| wc -l ) -eq 2 ]; then  
  docker rm -f $DOCKNAME
fi

docker run -d --name $DOCKNAME -p 80:8080 $IMGNAME:latest
systemctl restart $SVCNAME 