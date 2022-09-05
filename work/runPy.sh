#!/bin/bash

cd "Py/"

echo $1 | python Py$2.py 1> Py$2.out 2> Py$2.err &

if [ `ps -ef | grep Py$2 | grep -v grep | awk -F ' ' '{print $2}'` ]
then
    sleep 1

    if [ `ps -ef | grep Py$2 | grep -v grep | awk -F ' ' '{print $2}'` ]
    then
        kill -9 `ps -ef | grep Py$2 | grep -v grep | awk -F ' ' '{print $2}'`

        if [ $? -ne 0 ]; then
            echo "killWrong" >> Py$2.err
        fi
        echo -e "...\n\nOver Time" >> Py$2.info
    fi
fi


