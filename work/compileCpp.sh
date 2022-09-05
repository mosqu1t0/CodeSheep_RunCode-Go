#!/bin/bash

cd "Cpp/"

g++ -w "Cpp$1.cpp" -o "Cpp$1" 2> ./"Cpp$1.err"
