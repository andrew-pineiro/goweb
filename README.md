# goweb

[![GoBuild](https://github.com/andrew-pineiro/goweb/actions/workflows/go.yml/badge.svg)](https://github.com/andrew-pineiro/goweb/actions/workflows/go.yml)

## Summary

**goweb** is a website engine written in [GoLang](https://golang.org/). Including HTML/JS rendering for websites and an API endpoint for other various tasks. Currently this is being actively used on [www.andrewpineiro.com](www.andrewpineiro.com).

## Requirements
1. Execute the binary files depending on your operating system.
2. If you specify `-Api`, a `token.secret` file needs to exist in the root directory or supplied with `-AuthToken`. This is to store the secret API token and is required for startup.
3. Add your front-end code to the `www/` folder, otherwise nothing will load on your website.

## Build
**goweb** comes with a native build file, `build.go` in the root of the project. To start building **goweb**, type `go build build.go` in the root of the project. Run the build executable created, it supports the following flags:
```
  -arch string
        Architecture to build application under (default "amd64")
  -config string
        Configuration to build the project under (Debug/Release) (default "debug")
  -dir string
        Directory with main function for build (default ".")
  -name string
        Name of the project (default "main")
  -os string
        Operating systme to build application under (default "linux")
  -output string
        Directory to output  build files (default "./bin")
  -publish
        Enable publish mode in build
```

## Docker Setup

Ideally, the only setup needed is running the `docker-build.sh` script with root permissions. If this is not the case, please create an [issue](https://github.com/andrew-pineiro/goweb/issues).

This requires `docker.io` is installed on your machine.

**NOTE**: this has only been tested and setup for Debian Linux. 

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