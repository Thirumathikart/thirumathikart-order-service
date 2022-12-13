#!/bin/sh  
export PATH=$PATH:$(pwd)/bin &&
export PATH=$PATH:/go/bin/  &&
export PATH=$PATH:/usr/local/go/bin &&
git submodule init  &&
git submodule update --recursive    &&
mkdir generated &&
make proto
