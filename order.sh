#!/bin/bash
killall order
go build
nohup ./order &
