#!/bin/bash

count=0

date

while true
do
    count=$count+1
    curl --silent $1 > /dev/null
    echo -n "."
    if [[ $(($count % 1000)) == 0 ]]; then
        date
        echo "\n$count"
    fi
done