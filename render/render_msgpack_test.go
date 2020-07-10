// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package render

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vmihailenco/msgpack/v5"
)

// TODO unit tests
// test errors

func TestRenderMsgPack(t *testing.T) {
	w := httptest.NewRecorder()
	data := map[string]interface{}{
		"foo": "bar",
	}

	(MsgPack{data}).WriteContentType(w)
	assert.Equal(t, "application/msgpack; charset=utf-8", w.Header().Get("Content-Type"))

	err := (MsgPack{data}).Render(w)

	assert.NoError(t, err)

	buf := bytes.NewBuffer([]byte{})
	assert.NotNil(t, buf)
	err = msgpack.NewEncoder(buf).Encode(data)

	assert.NoError(t, err)
	assert.Equal(t, w.Body.String(), string(buf.Bytes()))
	assert.Equal(t, "application/msgpack; charset=utf-8", w.Header().Get("Content-Type"))
}
