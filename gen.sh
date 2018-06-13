#!/bin/bash
protoc \
  --plugin=protoc-gen-ts=./web/node_modules/.bin/protoc-gen-ts \
  --plugin=protoc-gen-go=./api/bin/protoc-gen-go \
  -I ./proto \
  --js_out=import_style=commonjs,binary:./web/src/pb \
  --go_out=plugins=grpc:./api/src/github.com/keywee/api/pb \
  --ts_out=service=true:./web/src/pb \
  --python_out=./topicextractor \
  ./proto/keywee.proto
