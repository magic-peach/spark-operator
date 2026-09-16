/*
Copyright 2025 The Kubeflow authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package version

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func capturePrintVersion(t *testing.T, short bool) string {
	original := os.Stdout
	r, w, err := os.Pipe()
	assert.NoError(t, err)
	os.Stdout = w

	PrintVersion(short)

	assert.NoError(t, w.Close())
	os.Stdout = original
	out, err := io.ReadAll(r)
	assert.NoError(t, err)
	return string(out)
}

func TestPrintVersionShort(t *testing.T) {
	out := capturePrintVersion(t, true)
	assert.True(t, strings.HasPrefix(out, "Git Version:"))
	assert.NotContains(t, out, "Go Version:")
}

func TestPrintVersionFull(t *testing.T) {
	out := capturePrintVersion(t, false)
	assert.Contains(t, out, "Git Version:")
	assert.Contains(t, out, "Git Commit:")
	assert.Contains(t, out, "Git Tree State:")
	assert.Contains(t, out, "Build Date:")
	assert.Contains(t, out, "Go Version:")
	assert.Contains(t, out, "Compiler:")
	assert.Contains(t, out, "Platform:")
}
