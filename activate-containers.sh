#!/bin/sh

CONTAINER_NAME=plac-container
CONTAINER_IMAGE=ubuntu
CONTAINER_CID_PATH=./.containerid

docker run --rm --name=$CONTAINER_NAME --cidfile=$CONTAINER_CID_PATH -it $CONTAINER_IMAGE /bin/sh
