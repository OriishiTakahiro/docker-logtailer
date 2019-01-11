#!/bin/sh

CID=`cat ./.containerid`

tail -f /var/lib/docker/containers/$CID/$CID-json.log
