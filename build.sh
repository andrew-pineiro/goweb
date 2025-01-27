#!/bin/bash
set -x

GOWEBPATH="/opt/goweb"
APPDIR="$GOWEBPATH/www"
LOGPATH="$GOWEBPATH/logs"
LOGSHPATH="$GOWEBPATH/www/logging/logs.sh"
SVCPATH="$GOWEBPATH/www/logging/weblogs.service"
SVCNAME="weblogs.service"
DOCKNAME="goweb_1"
IMGNAME="goweb"

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


docker build --tag $IMGNAME $APPDIR

if [ $( docker ps -a -f name=$DOCKNAME| wc -l ) -eq 2 ]; then  
  docker rm -f $DOCKNAME
fi

docker run -d --name $DOCKNAME -p 80:8080 $IMGNAME:latest

systemctl status $SVCNAME >/dev/null
EXITCODE=$?
if [ $EXITCODE -eq 4 ]; then    
  cp $SVCPATH /etc/systemd/system/
  systemctl enable $SVCNAME
elif [ ! $EXITCODE -eq 0 ]; then
  systemctl start $SVCNAME
else
  systemctl restart $SVCNAME 
fi