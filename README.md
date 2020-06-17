# syslgen-proto

## Features 

Generate *.proto and *.pb.go source files from sysl source.

See demo/Authorisation.sysl which has been used as an example for generating proto files from Sysl.
The proto files are created in the gen/proto/authorisation directory.

## Dependencies

Install the following and ensure that they are available in the path.

1. arr.ai from https://github.com/arr-ai/arrai
2. sysl from https://sysl.io/docs/install/
3. protobuf from https://grpc.io/docs/protoc-installation/

## Installation

`go get -u -v github.com/anz-bank/syslgen-proto`

## Usage

1. Put the sysl file in the /demo directory.
2. Run run.sh.
3. Provide the name and package of the target Sysl Application on the prompt.
4. The proto files are created in the gen/proto/package directory.

## Introduction

protobufs are a great tool to model the individual apps; coupled with grpc it can really help you build apps and services which are performance sensitive. So you could define your services in protobuf and have them auto generated in any language.

Sysl however works the same but operates at a much higher level and also helps you model the interactions between various apps in the enterprise landscape. Each app in a Sysl file, could potentially be mapped to a single protobuf file.

Please refer to the demo/Authorisation.sysl example.

## Model workflow

```
+-----------------------+   
| generated protobuf    |
+-----------------------+           
     ^                                     
     |    syslgen-proto
     |                           
  +--+-----------+               
  | .sysl file   |                
  +--------------+                

```