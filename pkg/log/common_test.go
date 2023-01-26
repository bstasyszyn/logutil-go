/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package log

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCommonLogs(t *testing.T) {
	const module = "test_module"

	t.Run("CloseIteratorError", func(t *testing.T) {
		stdOut := newMockWriter()

		logger := New(module,
			WithStdOut(stdOut),
			WithFields(WithAddress("some.address")),
			WithEncoding(Console),
		)

		CloseIteratorError(logger, errors.New("iterator error"))

		require.Contains(t, stdOut.Buffer.String(), `Error closing iterator`)
		require.Contains(t, stdOut.Buffer.String(), `"address": "some.address"`)
		require.Contains(t, stdOut.Buffer.String(), `"error": "iterator error"`)
		require.Contains(t, stdOut.Buffer.String(), "log/common_test.go")
	})

	t.Run("CloseResponseBodyError", func(t *testing.T) {
		stdOut := newMockWriter()

		logger := New(module,
			WithStdOut(stdOut),
			WithFields(WithAddress("some.address")),
			WithEncoding(Console),
		)

		CloseResponseBodyError(logger, errors.New("response body error"))

		require.Contains(t, stdOut.Buffer.String(), `Error closing response body`)
		require.Contains(t, stdOut.Buffer.String(), `"address": "some.address"`)
		require.Contains(t, stdOut.Buffer.String(), `"error": "response body error"`)
		require.Contains(t, stdOut.Buffer.String(), "log/common_test.go")
	})

	t.Run("WriteResponseBodyError", func(t *testing.T) {
		stdErr := newMockWriter()

		logger := New(module,
			WithStdErr(stdErr),
			WithFields(WithAddress("some.address")),
			WithEncoding(Console),
		)

		WriteResponseBodyError(logger, errors.New("response body error"))

		require.Contains(t, stdErr.Buffer.String(), `Error writing response body`)
		require.Contains(t, stdErr.Buffer.String(), `"address": "some.address"`)
		require.Contains(t, stdErr.Buffer.String(), `"error": "response body error"`)
		require.Contains(t, stdErr.Buffer.String(), "log/common_test.go")
	})

	t.Run("ReadRequestBodyError", func(t *testing.T) {
		stdErr := newMockWriter()

		logger := New(module,
			WithStdErr(stdErr),
			WithFields(WithAddress("some.address")),
			WithEncoding(Console),
		)

		ReadRequestBodyError(logger, errors.New("request body error"))

		require.Contains(t, stdErr.Buffer.String(), `Error reading request body`)
		require.Contains(t, stdErr.Buffer.String(), `"address": "some.address"`)
		require.Contains(t, stdErr.Buffer.String(), `"error": "request body error"`)
		require.Contains(t, stdErr.Buffer.String(), "log/common_test.go")
	})

	t.Run("WroteResponse", func(t *testing.T) {
		SetLevel(module, DEBUG)

		stdOut := newMockWriter()

		logger := New(module,
			WithStdOut(stdOut),
			WithFields(WithAddress("some.address")),
			WithEncoding(Console),
		)

		WroteResponse(logger, []byte("some response"))

		require.Contains(t, stdOut.Buffer.String(), `Wrote response`)
		require.Contains(t, stdOut.Buffer.String(), `"address": "some.address"`)
		require.Contains(t, stdOut.Buffer.String(), `"response": "some response"`)
		require.Contains(t, stdOut.Buffer.String(), "log/common_test.go")
	})
}
