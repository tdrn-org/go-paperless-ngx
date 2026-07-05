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

package paperlessngx_test

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
	paperlessngx "github.com/tdrn-org/go-paperless-ngx"
)

func TestClient(t *testing.T) {
	_ = newClient(t)
}

func newClient(t *testing.T) *paperlessngx.Client {
	apiURL, err := url.Parse("https://www.holger.mobi/home/paperless")
	require.NoError(t, err)
	client, err := paperlessngx.NewClient(apiURL, "")
	require.NoError(t, err)
	return client
}
