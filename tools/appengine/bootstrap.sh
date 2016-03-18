#!/bin/bash

goapp=`which goapp 2>&1`
if [[ ! "${?}" = "0" ]]; then
  echo "Need to install `goapp` to execute this script."
  exit 127
fi
GOGET="${goapp} get"

$GOGET -u google.golang.org/appengine/...
$GOGET -u google.golang.org/cloud/...
$GOGET -u golang.org/x/oauth2/google

$GOGET -u github.com/gin-gonic/gin
$GOGET -u github.com/stretchr/testify/assert
$GOGET -u github.com/kaneshin/pigeon
$GOGET -u github.com/olahol/go-imageupload
