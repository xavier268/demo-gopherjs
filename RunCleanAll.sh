#! /bin/bash

echo "**************************************************************"
echo "    Clean all caches and compiled files"
echo "**************************************************************"

rm -rdfv *.js *.map *.lock.json node_modules *.sum
go mod tidy

