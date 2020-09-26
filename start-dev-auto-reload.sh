#!/usr/bin/env bash

[[ `systemctl status postgresql | awk '/Active/{print $2}'` == inactive ]] && sudo systemctl start postgresql

CompileDaemon -build="go build" -include="*.pug" -include="*.gohtml" -include="*.css" -recursive="true" -command="./ws"