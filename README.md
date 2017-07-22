# Golang GRPC-TLS Service to Dynamically Control & Configure Beats (A node.js GRPC client is provided.)

This Repo Consists of a Golang GRPC-TLS Service and proto file dependencies to Dynamically Control & Configure Beats. Also a node.js GRPC client is included which can be used to talk to the Golang Service.

## Installation:
1. Install Go latest Vesrion.
2. Install npm and node latest Version.
3. Install protoc compiler latest Version.

## Introduction and a Look into Development:

GRPC - Google RPC Protocol is a Open-Source RPC framework which uses protocol buffers binary serialization format.Automatically generate idiomatic client and server stubs for your service in a variety of languages and platforms.Bi-directional streaming and fully integrated pluggable authentication with http/2 based transport.(For Detailed Information & Tutorial on GRPC you can visit https://grpc.io/)

To start with building a GRPC Service one should download and install protobuf service. Here I have used protobuf-3.2.0.

The first step will be to create a .proto file. ".proto" file is the one where you define the structure of the Data and required variabels and functions to build Server and Clients.(You can find the .proto file created for this repo in iot directory).

On Creating a .proto file depending on the Language in which we want to build our Server and Clients we need to generate the proto functions library from the .proto file we have developed using the protobuf service("protoc" compiler).

In this Repo I have built the Service in Golang. Below are the steps to generate your Go library with the functions and structure and you have defined in the .proto file.

```
protoc --go_out=plugins=grpc:. ./iot.proto
```

It Generates a ".go" file with the required dependencies to build Servers and Clients.

And Here I have built a Client in node.js. For node.js Programs you can directly require the ".proto" file and start using all the data structures defined in that ".proto" file.(No need of using protoc to generate a ".js" file for node.js programs).

## To Generate Private and Public Keys for TLS:

If you are running the Service with TLS Layer Enabled(-tls=true), then you are to follow the below steps to generate the required private & self-signed public keys for TLS.

Key considerations for algorithm "RSA" = 2048-bit
```
openssl genrsa -out server.key 2048
```

Generation of self-signed(x509) public key (PEM-encodings .pem|.crt) based on the private (.key):
```
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
```
Here we have generated a private key using RSA 2048-bit Encryption and Self-Signed(x509) public key using the generated private key.

If you are running the Service with TLS Disabled(-tls=false), you can Ignore this Section.

## To Run the Service:

To Start the Golang GRPC Service use the below command

```
go run beatgrpcservice.go -port 7771 -tls true -cert_file ./server.crt -key_file ./server.key
```
```
Flags:
-port - Port on which the Service should Run.
-tls  - Enable or Disable TLS on the Service (It should be a boolean value true/false).
-cert_file - Path to the Self-Signed Public Key File.
-key_file - Path to the encrypted Private Key File.
```

If Go is not installed on your machine, An Executable Binary File is also provided in the Repo and you can Run the Service using the below command:

```
./beatgrpcservice -port 7771 -tls true -cert_file ./server.crt -key_file ./server.key
```

## To Run the Client:

To start the node.js GRPC client to talk to the GRPC server execute the Below Command.(All the packages used in the Program are to be installed using npm before Starting this program):

```
node nodegrpcclient.js --beat=filebeat --action=status --cert_file=./server.crt --config_file=./config.yml --addr=localhost:7771
```
```
Flags:
-beat - Specify the type of beat.(filebeat/metricbeat).
-action - Choose the Action to be Performed.(start/stop/status/pause/resume/restart).
-cert_file - Path to the Self-Signed Public Key File.
-config_file - Path to the File to Read the Full YAML Configurations from.
-addr - Server Address & Port
```

## Author:

   NAGA SAI AAKARSHIT BATCHU (aakarshit.batchu@gmail.com)

