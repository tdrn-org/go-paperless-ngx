//
// Copyright (C) 2026 Holger de Carne
//
// This software may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.

package mock_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-paperless-ngx/mock"
)

func TestStartStop(t *testing.T) {
	serverMock := mock.Start()
	defer serverMock.Stop(t.Context())

	apiURL := serverMock.APIURL()
	require.NotNil(t, apiURL)
}
