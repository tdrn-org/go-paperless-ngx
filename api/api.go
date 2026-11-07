//
// Copyright (C) 2026 Holger de Carne
//
// This software may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.

// Package api contains the generated objects and functions of the Paperless NGX API.
package api

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=oapi-codegen-paperless-ngx.yaml paperless-ngx.yaml
//go:generate go run ../cmd/build/build.go genwrapper ../paperlessngx.gen.go
//go:generate go run ../cmd/build/build.go genmock ../mock/mock.gen.go
//go:generate go run ../cmd/build/build.go genapitest ../paperlessngx_test.go

import (
	_ "embed"
)

// The Paperless NGX API spec yaml file used to generate the API code.
//
//go:embed paperless-ngx.yaml
var PaperlessNGXRestApiSpecYaml []byte

// The header key to use for authorization.
const AuthorizationHeader = "Authorization"
