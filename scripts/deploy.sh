#!/bin/sh  
cd thirumathikart-order-service &&
git checkout . && 
git fetch origin prod &&
git reset --hard origin/prod &&
./scripts/dependencies.sh &&
./scripts/update_submod.sh &&
./scripts/run.sh
