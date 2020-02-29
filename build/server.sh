#!/bin/bash

echo 'Building Server Component...'
cd cmd/itchy-bird-server
go build -o ../../dist/itchy-bird-server .
