#!/bin/bash
set -x
export GOPATH=~/Workspace/go
export PATH=$GOPATH/bin:$PATH
`gometalinter > a.out`
