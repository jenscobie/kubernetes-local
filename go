#!/bin/bash

set -e

VBoxManage -v >/dev/null 2>&1 || { echo >&2 "VirtualBox is required. Please install the latest version."; exit 1; }
docker -v >/dev/null 2>&1 || { echo >&2 "Docker is required. Please install the latest version."; exit 1; }

function helptext {
    echo "Usage: ./go <command>"
    echo ""
    echo "Available commands are:"
    echo "    pods        List all pods in the Kubernetes cluster"
    echo "    start       Start a Kubernetes cluster locally"
    echo "    stop        Stop the local Kubernetes cluster"
}

function pods {
  kubectl get pods
}

function start {
  docker-compose up -d
}

function stop {
  docker-compose stop
}

[[ $@ ]] || { helptext; exit 1; }

case "$1" in
    help) helptext
    ;;
    pods) pods
    ;;
    start) start
    ;;
    stop) stop
    ;;
esac
