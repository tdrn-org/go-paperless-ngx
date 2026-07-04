//go:build tools
// +build tools

/*
 * Copyright (C) 2026 Holger de Carne
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"bytes"
	"go/format"
	"log"
	"os"

	"github.com/tdrn-org/go-paperless-ngx/api"
)

// Used via go:generate to perform build tasks.
func main() {
	switch os.Args[1] {
	case "genwrapper":
		genwrapper()
	}
}

// generate wrapper code
func genwrapper() {
	buffer := &bytes.Buffer{}
	err := api.GenerateClientWrapper(buffer)
	if err != nil {
		log.Fatal("generate failure", err)
	}
	codeBytes := buffer.Bytes()
	codeBytes, err = format.Source(codeBytes)
	if err != nil {
		log.Fatal("format failure", err)
	}
	wrapperFile := os.Args[2]
	err = os.WriteFile(wrapperFile, codeBytes, 0666)
	if err != nil {
		log.Fatal("write failure", err)
	}
}
