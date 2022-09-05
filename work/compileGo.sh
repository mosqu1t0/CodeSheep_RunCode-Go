#!/bin/bash

cd "Go/"

go build "Go$1.go" 2> ./"Go$1.err"
