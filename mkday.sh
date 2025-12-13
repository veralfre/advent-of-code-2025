#!/bin/zsh


# Script to create a new day directory with boilerplate code
if [ $# -ne 1 ]; then
    echo "Usage: $0 <day-number>"
    exit 1
fi

DAY_NUM=$1
DAY_DIR="cmds/day$DAY_NUM"
if [ -d "$DAY_DIR" ]; then
    echo "Directory $DAY_DIR already exists!"
    exit 1
fi
mkdir -p "$DAY_DIR"
touch "$DAY_DIR/input.txt"
touch "$DAY_DIR/main.go"