#!/bin/bash

while true; do
    curl --unix-socket /tmp/log.sock http://127.0.0.1/alloc
done
