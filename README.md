# Itchy Bird

Itchy Bird is a tool created for developers to easily handle distribution of their favorite binaries/files across many devices connected to the internet.

## About the project

This project is still under heavy development and is not ready for production environment. (Be Warned!)

There are two parts to this project, the server and the client cli.
The server handles the keeping track of distributions by keeping the binaries right next to the server or pulling in from dedicated object storage. Such as AWS, GCP, digital ocean.
The client cli is installed on a client's machine  

## Building from source
To better suit your OS target, it is recomended to build these applications from source code. This is a temporary approach until more distributions become available to the public.

To build the server component follow the command bellow from project root.
```
cd server
go get
bash bin/build.sh
```

The `go build` command will automatically build a binary for your current OS. Said binary will be located in the `server/dist` directory.

## Downloads
Currently there is only one distribution available for linux x84-64 arch located in `server/dist`.

## Running the server

1. Navigate to the location of the server component binary
2. Start the server `./itchy-bird-server`