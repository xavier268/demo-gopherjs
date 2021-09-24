#! /bin/bash

echo "**************************************************************"
echo "          Running main in node $( node --version )"
echo "**************************************************************"

npm install

~/go/bin/gopherjs build *.go
node main.js
