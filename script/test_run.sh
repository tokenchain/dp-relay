#!/usr/bin/env bash
. ./_auth.sh
cd $COMPILE
make build
$TESTPATH --start-rest-server --port "0.0.0.0:1333" --root $CONFIG