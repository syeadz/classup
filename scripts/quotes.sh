#!/bin/bash

# FILE=/etc/classup/quotes.csv
FILE=/home/syead/info/quotes.csv
if [ -f "$FILE" ]; then
    echo "$FILE exists."
    exit
fi

if [ "$EUID" -ne 0 ]
  then echo "error: not root (use sudo)" >&2 ; exit
fi

curl https://raw.githubusercontent.com/akhiltak/inspirational-quotes/master/Quotes.csv --create-dirs -o $FILE

if [ -f "$FILE" ]; then
    echo "Successfully downloaded quotes.csv at /etc/classup/"
    exit
else
    echo "Error: file wasn't created" >&2 ; exit
fi

