#! /bin/bash

echo "**************************************************************"
echo "          Running main in browser $( firefox --version )"
echo "**************************************************************"
# ensure sourece-map tools gets installed
npm install 

# Launch firefox with development tools open
firefox --devtools  http://localhost:8080/ &

# compile go files (minified) to js 
echo "Compiled without minification"
~/go/bin/gopherjs build *.go
ls -lah main.js

echo "Compiled with minification"
~/go/bin/gopherjs build -m *.go
ls -lah main.js

# and serve them ...
echo "Close browser and press Ctrl-C to exit server"
~/go/bin/gopherjs serve . 
