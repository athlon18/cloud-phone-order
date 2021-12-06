#!/bin/bash
git pull
killall order
go build
nohup ./order & tail -f nohup.out
