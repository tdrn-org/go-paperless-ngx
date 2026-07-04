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

package api

import (
	"fmt"
	"go/types"
	"io"
	"strings"

	"golang.org/x/tools/go/packages"
)

const generateTypeName string = "ClientWithResponsesInterface"

func GenerateClientWrapper(out io.StringWriter) error {
	config := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedCompiledGoFiles | packages.NeedImports | packages.NeedTypes | packages.NeedTypesSizes | packages.NeedSyntax | packages.NeedTypesInfo,
	}
	pkgs, err := packages.Load(config, ".")
	if err != nil {
		return err
	}
	for _, pkg := range pkgs {
		obj := pkg.Types.Scope().Lookup(generateTypeName)
		if obj == nil {
			continue
		}
		src, ok := obj.Type().Underlying().(*types.Interface)
		if !ok {
			return fmt.Errorf("type %s is not an interface", generateTypeName)
		}
		generator := &wrapperGenerator{
			out: out,
			src: src,
		}
		return generator.Run()
	}
	return fmt.Errorf("failed to lookup %s", generateTypeName)
}

type wrapperGenerator struct {
	out io.StringWriter
	src *types.Interface
}

func (g *wrapperGenerator) Run() error {
	g.writePreamble()
	for method := range g.src.Methods() {
		signature := method.Type().(*types.Signature)
		g.writeParams(method, signature)
		g.writeResults(signature)
		g.out.WriteString(" {\n")
		g.writeBody(method, signature)
		g.out.WriteString("}\n\n")
	}
	return nil
}

func (g *wrapperGenerator) writePreamble() {
	g.out.WriteString("// generated from api." + generateTypeName + "\n")
	g.out.WriteString("package paperlessngx\n")
	g.out.WriteString("import (\n")
	g.out.WriteString("\"context\"\n")
	g.out.WriteString("\"io\"\n")
	g.out.WriteString("\n")
	g.out.WriteString("\"github.com/tdrn-org/go-paperless-ngx/api\"\n")
	g.out.WriteString(")\n")
}

func (g *wrapperGenerator) writeParams(method *types.Func, signature *types.Signature) {
	params := signature.Params()
	g.out.WriteString("func (client *Client) " + g.stripResponseSuffix(method.Name()) + "(")
	for i := 0; i < params.Len()-1; i++ {
		if i > 0 {
			g.out.WriteString(", ")
		}
		param := params.At(i)
		g.out.WriteString(param.Name() + " " + g.stripApiPackage(param.Type().String()))
	}
	g.out.WriteString(")")
}

func (g *wrapperGenerator) writeResults(signature *types.Signature) {
	results := signature.Results()
	resultsLen := results.Len()
	if resultsLen == 0 {
		return
	}
	g.out.WriteString(" (")
	for i := 0; i < resultsLen; i++ {
		if i > 0 {
			g.out.WriteString(", ")
		}
		result := results.At(i)
		g.out.WriteString(result.Name() + " " + g.stripApiPackage(result.Type().String()))
	}
	g.out.WriteString(")")
}

func (g *wrapperGenerator) writeBody(method *types.Func, signature *types.Signature) {
	params := signature.Params()
	g.out.WriteString("response, err := client.apiClient." + method.Name() + "(")
	for i := 0; i < params.Len()-1; i++ {
		param := params.At(i)
		g.out.WriteString(param.Name())
		g.out.WriteString(", ")
	}
	g.out.WriteString("client.authenticateRequest())\n")
	g.out.WriteString("if err != nil {\n")
	g.out.WriteString("return nil,client.wrapSystemError(err)\n")
	g.out.WriteString("}\n")
	g.out.WriteString("err = client.checkAPIResponse(response.HTTPResponse)\n")
	g.out.WriteString("if err != nil {\n")
	g.out.WriteString("return nil,err\n")
	g.out.WriteString("}\n")
	g.out.WriteString("return response,err\n")
}

func (g *wrapperGenerator) stripResponseSuffix(s string) string {
	return strings.TrimSuffix(s, "WithResponse")
}
func (g *wrapperGenerator) stripApiPackage(s string) string {
	return strings.Replace(s, "github.com/tdrn-org/go-paperless-ngx/", "", 1)
}
