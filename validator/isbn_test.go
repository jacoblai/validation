// Copyright 2014 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package validator

import (
	"testing"

	"github.com/caixw/lib.go/assert"
)

func TestEraseMinus(t *testing.T) {
	a := assert.New(t)

	a.Equal(eraseMinus([]byte("abc-def-abc-")), []byte("abcdefabc"))
	a.Equal(eraseMinus([]byte("abc-def-abc")), []byte("abcdefabc"))
	a.Equal(eraseMinus([]byte("-_abc-def-abc-")), []byte("_abcdefabc"))
	a.Equal(eraseMinus([]byte("-abc-d_ef-abc-")), []byte("abcd_efabc"))
}

func TestIsISBN(t *testing.T) {
	a := assert.New(t)

	a.True(IsISBN("1-919876-03-0"))
	a.True(IsISBN("0-471-00084-1"))
	a.True(IsISBN("978-7-301-04815-3"))
	a.True(IsISBN("978-986-181-728-6"))
}
