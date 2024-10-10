#!/bin/bash

cd ../
go build -o ./build ./
cd build
./auth
