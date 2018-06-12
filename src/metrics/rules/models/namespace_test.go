// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewNamespace(t *testing.T) {
	view := &NamespaceView{
		Name:              "name",
		ForRuleSetVersion: 1,
	}
	res := NewNamespace(view)
	expected := Namespace{
		ID:                "name",
		ForRuleSetVersion: 1,
	}
	require.Equal(t, expected, res)
}

func TestNewNamespaces(t *testing.T) {
	view := &NamespacesView{
		Version: 1,
		Namespaces: []*NamespaceView{
			&NamespaceView{
				Name:              "name1",
				ForRuleSetVersion: 1,
			},
			&NamespaceView{
				Name:              "name2",
				ForRuleSetVersion: 1,
			},
		},
	}
	res := NewNamespaces(view)
	expected := Namespaces{
		Version: 1,
		Namespaces: []Namespace{
			{
				ID:                "name1",
				ForRuleSetVersion: 1,
			},
			{
				ID:                "name2",
				ForRuleSetVersion: 1,
			},
		},
	}
	require.Equal(t, expected, res)
}