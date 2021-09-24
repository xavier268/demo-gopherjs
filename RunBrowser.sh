#! /bin/bash

echo "**************************************************************"
echo "          Running main in browser $( firefox --version )"
echo "**************************************************************"
# ensure sourece-map tools gets installed
npm install 

# Launch firefox with development tools open
firefox --devtools  http://localhost:8080/ &

# compile files and serve them ...
echo "=================================================="
echo "   Close browser and press Ctrl-C to exit server "
echo "=================================================="
echo ""
~/go/bin/gopherjs serve . 
