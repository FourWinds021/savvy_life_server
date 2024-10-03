#!/bin/bash

for log_file in $(find /usr/local/service/savvy_life/logs -type f -name "*.log*" -mtime +3 | xargs -n 1)
do
    rm -rf "$log_file"
done
