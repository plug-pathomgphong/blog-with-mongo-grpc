#!/bin/bash
# Go server
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative blog/blogpb/blog.proto


# JS server
protoc -I=blog/blogpb/ blog.proto --js_out=import_style=commonjs:blog/blogpb/web_client/
# JS client
protoc -I=blog/blogpb/ blog.proto --grpc-web_out=import_style=commonjs,mode=grpcwebtext:blog/blogpb/web_client/



# For web client (FE/BE same architecture)
protoc -I=blog/blogpb/ blog.proto --js_out=import_style=commonjs,binary:blog/jsclient/src/ --grpc-web_out=import_style=commonjs,mode=grpcwebtext:blog/jsclient/src/

# Example in web
# protoc --proto_path=todo --js_out=import_style=commonjs,binary:todo-client/src/ --grpc-web_out=import_style=commonjs,mode=grpcwebtext:todo-client/src/ todo/todo.proto