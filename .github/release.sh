#!/bin/sh

echo "OK!!"

apk add --no-cache curl jq git zip
apt-get update -qq \
    && apt-get install -y -q --no-install-recommends \
        curl \
        jq

jq
