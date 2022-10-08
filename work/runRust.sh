#!/bin/bash

cd "Rust/"

echo $1 | ./Rust$2 &>Rust$2.out &

if [ $(ps -ef | grep Rust$2 | grep -v grep | awk -F ' ' '{print $2}') ]; then
	sleep 1

	if [ $(ps -ef | grep Rust$2 | grep -v grep | awk -F ' ' '{print $2}') ]; then
		kill -9 $(ps -ef | grep Rust$2 | grep -v grep | awk -F ' ' '{print $2}')

		if [ $? -ne 0 ]; then
			echo "killWrong" >>Rust$2.err
		fi
		echo -e "...\n\nOver Time" >>Rust$2.info
	fi
fi
