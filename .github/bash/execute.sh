#!/usr/bin/env bash

# Usage:
#  alias docker='bash execute.sh docker'
if [[ "$1" == "docker" ]]; then
  if [[ "$2" == "ps" || "$2" == "images" ]]; then
    "$@" | dockercolorize
  elif [[ "$2" == "compose" && "$3" == "ps" ]]; then
    "$@" | dockercolorize
  elif [[ "$2" == "stats" && "$3" == "--no-stream" ]]; then
    "$@" | dockercolorize
  else
    "$@"
  fi
else
  "$@"
fi
