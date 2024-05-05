#!/bin/bash

# List all available charsets using iconv
charsets=$(iconv -l)

function buildCharset() {
  for ((i=128; i<256; i++)); do
    hex=$(printf %x $i)
    iconv -f "$1" -t UTF-8 <<< "$(printf "\x$hex")"
  done | grep -v "^\s*$"
}

function printCharset() {
  mapfile -t chars < <(buildCharset "$1")

  counter=0
  for c in "${chars[@]}"
  do
    echo -ne "$c "
    counter=$((counter + 1))

    if [ $counter == 16 ]; then echo ""; counter=0; fi
  done
}

printCharset "CP437"
