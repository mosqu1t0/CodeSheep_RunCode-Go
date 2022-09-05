#!/bin/bash

cd "Go/"

echo $1 | ./Go$2.go 1> Go$2.out &

if [ `ps -ef | grep Go$2 | grep -v grep | awk -F ' ' '{print $2}'` ]
then
    sleep 1

    if [ `ps -ef | grep Go$2 | grep -v grep | awk -F ' ' '{print $2}'` ]
    then
        kill -9 `ps -ef | grep Go$2 | grep -v grep | awk -F ' ' '{print $2}'`

        if [ $? -ne 0 ]; then
            echo "killWrong" >> Go$2.err
        fi
        echo -e "...\n\nOver Time" >> Go$2.info
    fi
fi


