// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto

import (
	"github.com/golang/protobuf/v2/internal/errors"
	"github.com/golang/protobuf/v2/reflect/protoreflect"
	"github.com/golang/protobuf/v2/runtime/protoiface"
)

// Message is the top-level interface that all messages must implement.
type Message = protoreflect.ProtoMessage

// errInternalNoFast indicates that fast-path operations are not available for a message.
var errInternalNoFast = errors.New("BUG: internal error (errInternalNoFast)")

func protoMethods(m Message) *protoiface.Methods {
	if x, ok := m.(protoiface.Methoder); ok {
		return x.XXX_Methods()
	}
	return nil
}

func checkRequiredFields(m protoreflect.Message, nerr *errors.NonFatal) {
	req := m.Type().RequiredNumbers()
	knownFields := m.KnownFields()
	for i, reqLen := 0, req.Len(); i < reqLen; i++ {
		num := req.Get(i)
		if !knownFields.Has(num) {
			nerr.AppendRequiredNotSet(string(m.Type().Fields().ByNumber(num).FullName()))
		}
	}
}
