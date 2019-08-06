#!/usr/bin/env bash
protoc -I . greet.proto --go_out=plugins=grpc:.