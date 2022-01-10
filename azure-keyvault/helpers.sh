#!/usr/bin/env bash
#set -euo pipefail

random_string() {
        cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w ${1:-8} | head -n 1 | tr '[:upper:]' '[:lower:]'
}

install() {
  echo Hello World
}

upgrade() {
  echo World 2.0
}

uninstall() {
  echo Goodbye World
}

# Call the requested function and pass the arguments as-is
"$@"
