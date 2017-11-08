// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package render

import (
	"net/http"
)

type Render interface {
	Render(http.ResponseWriter) error
	WriteContentType(w http.ResponseWriter)
}

var (
	_ Render = JSON{}
	_ Render = IndentedJSON{}
	_ Render = XML{}
	_ Render = String{}
	_ Render = Redirect{}
	_ Render = Data{}
	_ Render = HTML{}
)

func writeContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	header["Content-Type"] = value
}
