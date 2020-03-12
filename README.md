# Itchy Bird

Itchy Bird is a tool created for developers to easily handle distribution of their favorite binaries/files across many devices connected to the internet.

## About the project

This project is still under heavy development and is not ready for production environment. (Be Warned!)

There are two parts to this project, the server and the client cli.
The first is the server handles the keeping track of distributions by keeping the binary files in a centralized location. Either on prem or pulling in from dedicated object storage. Such as AWS, GCP, digital ocean.
The second is the client cli installed on a client's machine.

## Building from source
To better suit your OS target, it is recomended to build these applications from source code. This is a temporary approach until more distributions become available to the public.

To build the server component follow the command bellow from project root.
```
go get
bash build/server.sh
```

To build the client component follow the command bellow from project root.
```
go get
bash build/client.sh
```

The `go build` command will automatically build a binary for your current OS. Said binary will be located in the `dist/<component>` directory.


## Running the server

1. Navigate to the location of the server component binary
2. Start the server `./itchy-bird-server`


## Running the client

1. Navigate to the location of the client component binary
2. Run the client `./itchy-bird-client`

There are several commands available while running the client

| Command | Description |
| -- | -- |
| `--ls` | Get list of files in the local directory |
| `--ls-remote` | Get list of files from the remote directory |
| `--pull <file-name>` | Downloads requested file to a local directory | 
| `--update <file-name>` | Checks to see if there is a newer version avaialable and downloads requested file to a local directory (TODO) | 

#### Example
`./itchy-bird-client --pull file-name`


## How to use this package effectively

_This section will change as the application get developed_

The indended use for this application is to keep server component running somewhere on the cloud, while the client component can be installed on dedicated machines. (IoT,  dedicates servers, etc). The client component can be triggered manually, but it would be best to create a cron-job that will call `itchy-bird-client --update <file-name>` every so often.