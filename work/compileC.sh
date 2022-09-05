#!/bin/bash

cd "C/"

gcc -w "C$1.c" -o "C$1" 2> ./"C$1.err"
