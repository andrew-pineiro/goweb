# goweb

## Summary

**goweb** is a website engine written in [GoLang](https://golang.org/). Including HTML/JS rendering for websites and an API endpoint for other various tasks. Currently this is being actively used on [www.andrewpineiro.com](www.andrewpineiro.com).

## Requirements
1. This only currently works on Linux (and has only been tested on Debian)
2. Docker must be installed.
3. A `token.secret` file needs to exist in the root directory. This is to store the secret API token and is required for startup.


## Setup

Ideally, the only setup needed is running the `build.sh` script with root permissions. If this is not the case, please create an [issue](https://github.com/andrew-pineiro/goweb/issues).

### Steps
1. Checks for root application directory (/www); if found it is deleted.
2. Created /www directory.
3. Moves local **goweb** files to /www directory.
4. Checks for logging root path (/etc/website); creates if not found.
5. Checks for log file directory (/etc/website/logs); creates if not found.
6. Builds a new docker image with the local `Dockerfile`.
7. Checks for an existing **goweb** docker container is running; deletes if found. 
8. Runs a new docker container with the fresh image
9. Checks status of `weblogs.service` logging service; if exit code is 4 (not found) it attempts to create the service. Otherwise the service is started.
10. END


## Logging

The `/logging/logs.sh` script is executed by a system process which follows the docker containers logs. These logs are filled by the stdout of the **goweb** application.

For easier record maintenace, you may want to create a cron job that restarts the `weblogs.service` to generate new log files each day.

Example -
```bash
20 5    * * 0-6 root    systemctl restart weblogs.service 2>&1
```

### Logging Script

The script that starts the logging will also do some record maintenance (removes any logs older than 14 days) as well as renaming the existing log file to the date of the logs.

#### Steps
1. Rename `website.log` to include today's date
2. Purge files older than 14 days
3. Follow docker logs and append to `website.log` file 