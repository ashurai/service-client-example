#!/bin/bash
cd /go/src/app || exit 1
apt-get update && apt-get install make
make intigration-test