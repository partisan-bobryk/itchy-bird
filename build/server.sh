#!/bin/bash

echo 'Building Server Component...'
cd cmd/itchy-bird-server
go build -o ../../dist/server/itchy-bird-server .

[ $? -eq 0 ] && echo "Done!"