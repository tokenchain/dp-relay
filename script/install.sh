#!/usr/bin/env bash
. ./_auth.sh
echo "====================================="
echo "install files"
echo "====================================="
#unzip achive.zip -d /usr/bin
if [[ ! -f $DEMONPATH ]]; then
    mkdir -p $DEMONPATH
fi
cp linux/$DEMON $DEMONPATH/$DEMON
echo "====================================="
echo "install completed"
echo "====================================="
