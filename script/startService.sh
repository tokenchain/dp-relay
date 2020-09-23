#!/usr/bin/env bash

. ./_auth.sh

netstat -an | awk '/^tcp/ {++sta[$NF]} END {for(key in sta) print key,"\t",sta[key]}'
nohup $DEMONPATH/$DEMON --port "0.0.0.0:1320" --root $CONFIG >> $LOGFILE &
#p2p.persistent_peers=192.168.1.3:46656,192.168.1.4:46656,192.168.1.5:46656\ --home