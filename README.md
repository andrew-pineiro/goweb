# goweb

[![GoBuild](https://github.com/andrew-pineiro/goweb/actions/workflows/go.yml/badge.svg)](https://github.com/andrew-pineiro/goweb/actions/workflows/go.yml)

## Summary

**goweb** is a website engine written in [GoLang](https://golang.org/). Including HTML/JS rendering for websites and an API endpoint for other various tasks. Currently this is being actively used on [www.andrewpineiro.com](www.andrewpineiro.com).

## Requirements
1. This only currently works on Linux (and has only been tested on Debian)
2. Docker must be installed. (`apt install docker.io`)
3. If you specify `-Api`, a `token.secret` file needs to exist in the root directory or supplied with `-AuthToken`. This is to store the secret API token and is required for startup.
4. Add your front-end code to the `www/` folder, otherwise nothing will load on your website.


## Setup

Ideally, the only setup needed is running the `build.sh` script with root permissions. If this is not the case, please create an [issue](https://github.com/andrew-pineiro/goweb/issues).

### Build Script Steps
1. Checks for `goweb` user; creates if doesn't exist.
2. Checks for application directory (/opt/goweb/www); if found it is deleted.
3. Creates /opt/goweb/www directory. Makes `goweb` account the owner.
4. Remaining part of the script runs as the `goweb` account.
5. Moves local **goweb** files to application directory.
6. Checks for logging root path (/opt/goweb/logs); creates if not found.
7. Builds a new **goweb** docker image with the local `Dockerfile`.
8. Checks for an existing **goweb** docker container is running; deletes if found. 
9. Runs a new docker container with the fresh image
10. Checks status of `weblogs.service` logging service; if exit code is 4 (not found) it attempts to create the service. Otherwise the service is started.


## Logging

The `./logging/logs.sh` script is executed by a system process which follows the docker containers logs. These logs are filled by the stdout of the **goweb** application.

For easier record maintenace, you may want to create a cron job that restarts the `weblogs.service` to generate new log files each day.

Example -
```bash
20 5    * * 0-6 goweb    systemctl restart weblogs.service 2>&1
```

### Logging Script

The script that starts the logging will also do some record maintenance (removes any logs older than 14 days) as well as renaming the existing log file to the date of the logs.

#### Steps
1. Rename existing `website.log` to include today's date
2. Purge files older than 14 days
3. Follow docker logs and append to `website.log` file 