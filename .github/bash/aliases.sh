#!/usr/bin/env bash

di() {
  docker images "$@" | dockercolorize
}

dps() {
  docker ps "$@" | dockercolorize
}

dcps() {
  docker compose ps "$@" | dockercolorize
}

dpsa() {
  docker ps -a "$@" | dockercolorize
}

dstats() {
  docker stats --no-stream "$@" | dockercolorize
}