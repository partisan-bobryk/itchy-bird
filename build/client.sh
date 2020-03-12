#!/bin/bash

echo 'Building Client Component...'
cd cmd/itchy-bird-client
go build -o ../../dist/client/itchy-bird-client .

[ $? -eq 0 ] && echo "Done!"