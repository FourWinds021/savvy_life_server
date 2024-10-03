#!/bin/bash

func() {
    echo "USAGE:$(basename "$0") [-v version]"
    exit 0
}

while getopts v:h OPTION
do
    case $OPTION in
        v) version=$OPTARG;;
        h) func;;
        ?) func;;
    esac
done

# apt update
# apt-get install -y zip

cd "/usr/local/service" || exit
mkdir "$version"
tar -zxvf "$version.tar.gz" -C "$version" --strip-components 1
./savvy_life/admin/savvy_life.sh stop
./savvy_life/admin/savvy_life.sh stop
cp -rf ./savvy_life/config/config.ini "$version/config/"
rm -rf savvy_life
ln -s "$version" savvy_life
# ./savvy_life/admin/savvy_life.sh start
