#!/usr/bin/env bash

. ./_auth.sh
kill -9 $(ps -ef|grep $DEMON|gawk '$0 !~/grep/ {print $2}' |tr -s '\n' ' ')