// +build tools

package tools

import (
	// Tools as dependencies: https://github.com/go-modules-by-example/index/blob/master/010_tools/README.md
	_ "github.com/golang/protobuf/protoc-gen-go"
	_ "github.com/pressly/goose/cmd/goose"
)
