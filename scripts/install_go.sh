#!/bin/bash

EXPORT="Now you need to add export PATH=%\$PATH:/usr/local/go/bin to .bashrc"
NAME="go1.20.linux-amd64.tar.gz"

if [ -x "$(command -v go)" ]; then
  echo 'GO is already installed'
  exit
fi

if [ -f "/usr/local/go/bin" ]; then
    echo $EXPORT
else
    wget -O /tmp/$NAME https://go.dev/dl/$NAME
    tar -C /usr/local -xzf /tmp/$NAME
    echo "GO installed successfully"
fi

echo $EXPORT