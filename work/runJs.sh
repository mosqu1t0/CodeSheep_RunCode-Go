#!/bin/bash

cd "Js/"

echo $1 | node Js$2.js 1> Js$2.out 2> Js$2.err &

if [ `ps -ef | grep Js$2 | grep -v grep | awk -F ' ' '{print $2}'` ]
then
    sleep 1

    if [ `ps -ef | grep Js$2 | grep -v grep | awk -F ' ' '{print $2}'` ]
    then
        kill -9 `ps -ef | grep Js$2 | grep -v grep | awk -F ' ' '{print $2}'`

        if [ $? -ne 0 ]; then
            echo "killWrong" >> Js$2.err
        fi
        echo -e "...\n\nOver Time" >> Js$2.info
    fi
fi


