#!/bin/bash

URL=https://gist.githubusercontent.com/JakubPetriska/060958fd744ca34f099e947cd080b540/raw/963b5a9355f04741239407320ac973a6096cd7b6/quotes.csv
FILE=~/.classup/quotes.csv
if [ -f "$FILE" ]; then
    echo "$FILE already exists."
    exit
fi

if [ "$EUID" -ne 0 ]
  then echo "error: not root (use sudo)" >&2 ; exit
fi

curl $URL --create-dirs -o $FILE

if [ -f "$FILE" ]; then
    echo "Successfully downloaded quotes.csv at $FILE"
    exit
else
    echo "Error: file wasn't created" >&2 ; exit
fi

