#!/bin/bash

clear
echo "building..."
rm backend-dev
go build -o backend-dev
echo "starting..."
./backend-dev