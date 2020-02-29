#!/bin/bash

pwd
go build -o dist/itchy-bird-server cmd/server.go

# Create a dir to hold all of the distributions available for download
mkdir downloads