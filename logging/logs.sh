#!/bin/bash
DATE=$(date +%F)
LOGPATH="/etc/website/logs"

if [ -f "$LOGPATH/website.log" ]; then
    mv $LOGPATH/website.log "$LOGPATH/$DATE-website.log"
fi

find $LOGPATH -mindepth 1 -mtime +14 -delete

docker logs -f --since 5m goweb_1 &> $LOGPATH/website.log