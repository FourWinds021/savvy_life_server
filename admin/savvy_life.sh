#!/bin/bash

exe_path="$( cd "$( dirname "$0"  )" && pwd  )"
SAVVY_LIFE_HOME=$exe_path/../
SAVVY_LIFE_SBIN=$SAVVY_LIFE_HOME/bin/
SAVVY_LIFE_PID=$SAVVY_LIFE_HOME/proc/savvy_life.pid

start()
{
    if [ -f "$SAVVY_LIFE_PID" ]
    then
        pid=$(cat "$SAVVY_LIFE_PID")
        # shellcheck disable=SC2126
        # shellcheck disable=SC2009
        process_num=$(ps -ef | grep -w "$pid" | grep -v "grep" | grep "savvy_life" | wc -l)
        if [ "$process_num" -ge 1 ];
        then
            echo "service already running. pid=$pid"
            return
        fi  
    fi
    # shellcheck disable=SC2164
    cd "$SAVVY_LIFE_SBIN"
    # shellcheck disable=SC2261
    nohup ./savvy_life &> /dev/null 2>> ../logs/savvy_life_except.log &
    echo "savvy_life start"
}

stop()
{
    if [ ! -f "$SAVVY_LIFE_PID" ]
    then
        echo "service already exit"
        return
    fi
    # shellcheck disable=SC2006
    pid=`cat "$SAVVY_LIFE_PID"`
    # shellcheck disable=SC2006
    # shellcheck disable=SC2126
    # shellcheck disable=SC2009
    process_num=`ps -ef | grep -w "$pid" | grep -v "grep" | grep "savvy_life" | wc -l`
    if [ "$process_num" -eq 0 ];
    then
        echo "service already exit"
        return
    fi 
    # shellcheck disable=SC2046
    # shellcheck disable=SC2006
    kill -TERM `cat "$SAVVY_LIFE_PID"`
    ret=$?
    if [ $ret -eq 0 ]
    then
        echo "savvy_life stop"
    else
        echo "savvy_life stop failed"
    fi
    return
}

restart()
{
    stop
    start
    return
}

reload()
{
    if [ ! -f "$SAVVY_LIFE_PID" ]
    then
        echo "service already exit"
        return
    fi
    # shellcheck disable=SC2006
    pid=`cat "$SAVVY_LIFE_PID"`
    # shellcheck disable=SC2006
    # shellcheck disable=SC2126
    # shellcheck disable=SC2009
    process_num=`ps -ef | grep -w "$pid" | grep -v "grep" | grep "savvy_life" | wc -l`
    if [ "$process_num" -eq 0 ];
    then
        echo "service already exit"
        return
    fi 
    # shellcheck disable=SC2046
    # shellcheck disable=SC2006
    kill -USR2 `cat "$SAVVY_LIFE_PID"`
    return
}

case "$1" in
    start)
        start
        ;;
    stop)
        stop
        ;;
    restart)
        restart
        ;;
    reload)
        reload
        ;;
    *)
        echo $"Usage: $0 {start|stop|restart|reload}"
        exit 1
esac

exit 0
