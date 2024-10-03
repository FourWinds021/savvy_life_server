#!/bin/bash

export PATH=/usr/local/bin:/bin:/usr/bin:/usr/local/sbin:/usr/sbin:/sbin

# shellcheck disable=SC2006
proc_pid=`ss -lnp | grep ':39252 ' | awk -F, '{print $2}'`
# shellcheck disable=SC2006
file_pid=`cat /usr/local/service/savvy_life/proc/savvy_life.pid`

if [ "$proc_pid" == "" ]
then
    # shellcheck disable=SC2006
    echo "`date`: proc_pid does not exist."
    cd /usr/local/service/savvy_life/admin && ./savvy_life.sh start
    exit 0
fi

if [ "$proc_pid" != "$file_pid" ]
then
    # shellcheck disable=SC2006
    echo "`date`: proc_pid is not the same as file_pid."
    echo "proc_pid = ${proc_pid}"
    echo "file_pid = ${file_pid}"
    # shellcheck disable=SC2009
    ps -ef | grep "$proc_pid" | grep -v grep | awk '{print $2}' | xargs kill -9
    cd /usr/local/service/savvy_life/admin && ./savvy_life.sh restart
else
    # shellcheck disable=SC2006
    echo "`date`: savvy_life service is running."
fi
