#!/usr/bin/env bash
set -euo pipefail

name=$(cat /cnab/app/foo/name.txt)

install() {
  echo "Hello, installing $name with password: $1"
}

upgrade() {
  echo "Hello, upgrading $name"
}

uninstall() {
  echo "Goodbye, $name"
}

# Call the requested function and pass the arguments as-is
"$@"
