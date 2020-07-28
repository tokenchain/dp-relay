# DP-Relay

Official Golang implementation of the Darkpool protocol.

Darkpool is a command function framework.

It needs to run under windows server version

# DP-Tech
For prerequisites and detailed build instructions.

Building `main.go` requires both a Go (version 1.13 or later) and a C compiler. You can install
them using your favourite package manager. Once the dependencies are installed, run main.go


## Running `main.go`

Going through all the possible command line flags is out of scope here (please consult our
but we've enumerated a few common parameter combos to get you up to speed quickly
on how you can run your own `main` instance.

- Base on Golang

# Contact

# Install Guide

Install mysql

Install BT platform

yum update

`yum install leveldb-devel`

`install core tendermint`

`wget https://dl.google.com/go/go1.13.linux-amd64.tar.gz`

sudo tar -C /usr/local -xzf go1.13.linux-amd64.tar.gz

# Use Guide
Please visit this documentation from: https://documenter.getpostman.com/view/2597586/T1Ds9Fs5

GET `http://localhost:1320/`
```json
{
    "endpoints": [
        "/hdwallet/create/mnemonic",
        "/",
        "/hdwallet/create/{at_index}/",
        "/hdwallet/recovery"
    ]
}
```

GET `http://localhost:1320/hdwallet/create/mnemonic`
```json
{"words":"annual job denial sleep misery guess apple april message jacket require afford swamp ticket erode stumble involve skate minute satoshi trick kit virtual boat"}
```

POST `http://localhost:1320/hdwallet/recovery`

REQUEST BODY:
```json
{
	"keywords": "annual job denial sleep misery guess apple april message jacket require afford swamp ticket erode stumble involve skate minute satoshi trick kit virtual boat",
	"names":["ann","peter","lason","ming","ming","mymy@gmail.com","lalalmoon"],
	"from_index":0
}
```


POST ` "/hdwallet/create/{at_index}/"`

REQUEST BODY:
```json
{
	"keywords": "annual job denial sleep misery guess apple april message jacket require afford swamp ticket erode stumble involve skate minute satoshi trick kit virtual boat"
}
```

RESPONSE SAMPLE:



# Build Guide

centos specific linux

`make centos`

The only native build

`make build`