#!/usr/bin/env bash
. ./_auth.sh
cd $COMPILE
$TESTPATH --sync-service --root $CONFIG