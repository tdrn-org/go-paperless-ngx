## Go module for Paperless NGX REST API access
[![GoDoc](https://godoc.org/github.com/tdrn-org/go-paperless-ngx?status.svg)](https://godoc.org/github.com/tdrn-org/go-paperless-ngx)
[![Build](https://github.com/tdrn-org/go-paperless-ngx/actions/workflows/build.yml/badge.svg)](https://github.com/tdrn-org/go-paperless-ngx/actions/workflows/build.yml)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=tdrn-org_go-paperless-ngx&metric=coverage)](https://sonarcloud.io/summary/new_code?id=tdrn-org_go-paperless-ngx)
[![Go Report Card](https://goreportcard.com/badge/github.com/tdrn-org/go-paperless-ngx)](https://goreportcard.com/report/github.com/tdrn-org/go-paperless-ngx)

### FRITZ! Smart-Home-API
The [FRITZ! Smart Home system](https://fritz.com/collections/smart-home/) provides a [REST API](https://fritz.com/pages/schnittstellen/) to access Smart Home capable devices to evaluate and control the connected devices. This library provides the Go wrapper for this API. It is based on the provided [OpenAPI specification](https://fritz.com/pages/schnittstellen/).

## License & Attribution
This project is licensed under the **Apache License 2.0** (see the [LICENSE](./LICENSE) file for details).

#### Third-Party Notices
This project uses the official OpenAPI specification from **[Paperless-ngx](https://github.com/paperless-ngx)** as a source for code generation.

* **Component:** Paperless-ngx OpenAPI Specification (`api/paperless-ngx.gen.yaml`)
* **Copyright:** Copyright (c) Paperless-ngx Contributors
* **License:** [GNU General Public License v3.0 (GPLv3)](https://github.com/blob/main/LICENSE)

The OpenAPI specification file remains under the terms of the GPLv3. The Go code generator itself and any newly generated output code are licensed under the Apache License 2.0.

### Disclaimer
This project is an independent software module and is not affiliated with, endorsed by, or sponsored by FRITZ! GmbH or any of its products. The term “FRITZ!” is a registered trademark of FRITZ! GmbH. All trademarks and registered trademarks mentioned herein are the property of their respective owners. This module is provided for educational and interoperability purposes only.