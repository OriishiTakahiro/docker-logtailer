#!/bin/sh

CONTAINER_NAME=plac-container
CONTAINER_IMAGE=ubuntu
CONTAINER_ID_PATH=./.containerid

if [ $CONTAINER_ID_PATH ] ; then
  rm $CONTAINER_ID_PATH
fi

docker run --rm --name=$CONTAINER_NAME --cidfile=$CONTAINER_ID_PATH -it $CONTAINER_IMAGE /bin/sh
