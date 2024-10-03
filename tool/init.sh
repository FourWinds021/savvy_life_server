#!/bin/bash

export PATH=$PATH:/usr/local/mysql/bin

if [[ $# -lt 1 ]]; then
    echo "USAGE:$(basename "$0") [-u mysql-user] [-p mysql-passwd] [-f sql-file]"
    exit 1
fi

func() {
    echo "USAGE:$(basename "$0") [-u mysql-user] [-p mysql-passwd] [-f sql-file]"
    exit 0
}

while getopts u:p:f:h OPTION
do
    case $OPTION in
        u) user=$OPTARG;;
        p) passwd=$OPTARG;;
        f) file=$OPTARG;;
        h) func;;
        ?) func;;
    esac
done

exe_path=$(pwd)/$(dirname "$0")
cd "$exe_path" || exit

mysql -u "$user" -p"$passwd" "savvy_data" -e"source $file"
