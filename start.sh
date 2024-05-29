#!/bin/sh
set -e

# echo "nameserver 8.8.8.8" >> /etc/resolv.conf

if [ ! -d /printer ]; then
  export GOPROXY="https://goproxy.io"
  cd /printer-src
  make clean
  make build
  mv build /printer
  cd ..
  rm -r /printer-src
fi

cd /printer
exec ./printer server
