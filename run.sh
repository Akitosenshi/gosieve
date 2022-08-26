#!/bin/bash

shell="$(getent passwd | grep $USER)"
shell=${shell##*:}

$shell -c 'go build -ldflags "-s -w" -o gosieve ./main.go && time ./gosieve'
