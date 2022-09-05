#!/bin/bash


cd "C/"

echo $1 | ./C$2 &> C$2.out &

if [ `ps -ef | grep C$2 | grep -v grep | awk -F ' ' '{print $2}'` ]
then
    sleep 1

    if [ `ps -ef | grep C$2 | grep -v grep | awk -F ' ' '{print $2}'` ]
    then
        kill -9 `ps -ef | grep C$2 | grep -v grep | awk -F ' ' '{print $2}'`

        if [ $? -ne 0 ]; then
            echo "killWrong" >> C$2.err
        fi
        echo -e "...\n\nOver Time" >> C$2.info
    fi
fi


