#!/bin/bash
./app &
#serve -s ./build -l 3000
nginx -g 'daemon off;'