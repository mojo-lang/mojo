# Lang Package for Mojo

## This package has AST definition for the Mojo language

## Install

1. install protoc
2. install gogoproto
3. install mojo

## Build

mojo build go // from protobuf to go


*.pb.go
*.mojo.go // array interface impl, 自动添加便利的成员函数
*.codec.go // for json/yaml/bson/protobuf decode encode
*.go // manually add
*_test.go // manually test case
