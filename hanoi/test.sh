#!/bin/bash

SCRIPTPATH="$( cd "$(dirname "$0")" ; pwd -P )"
cd $SCRIPTPATH

echo "run in1"
go run main.go < in1
echo "output should be:"
cat out1

echo ""

echo "run in2"
go run main.go < in2
echo "output should be:"
cat out2
