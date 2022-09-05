#!/bin/bash


cd "Cpp/"

echo $1 | ./Cpp$2 &> Cpp$2.out & #进程实现真异步　哈哈

if [ `ps -ef | grep Cpp$2 | grep -v grep | awk -F ' ' '{print $2}'` ]
then
    sleep 1

    if [ `ps -ef | grep Cpp$2 | grep -v grep | awk -F ' ' '{print $2}'` ]
    then
        kill -9 `ps -ef | grep Cpp$2 | grep -v grep | awk -F ' ' '{print $2}'`

        if [ $? -ne 0 ]; then
            echo "killWrong" >> Cpp$2.err
        fi
        echo -e "...\n\nOver Time" >> Cpp$2.info
    fi
fi


