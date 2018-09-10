#!/bin/bash

SCRIPTPATH="$( cd "$(dirname "$0")" ; pwd -P )"
cd $SCRIPTPATH

echo "test case 1:"
colordiff <(go run main.go < in1) out1

# echo ""

# echo "test case 2:"
# colordiff <(go run main.go < in2) out2
