#! /bin/bash

echo "**************************************************************"
echo "    Compare compiled file size with $( ~/go/bin/gopherjs version )"
echo "**************************************************************"

echo "Compiling main.go WITHOUT minification : "
~/go/bin/gopherjs build *.go
ls -hs main.js

echo "Compiling main.go WITH minification : "
~/go/bin/gopherjs build -m *.go
ls -hs main.js


