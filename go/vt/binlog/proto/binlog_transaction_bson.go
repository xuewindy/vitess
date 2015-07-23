// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto

// DO NOT EDIT.
// FILE GENERATED BY BSONGEN.

import (
	"bytes"

	"github.com/youtube/vitess/go/bson"
	"github.com/youtube/vitess/go/bytes2"
)

// MarshalBson bson-encodes BinlogTransaction.
func (binlogTransaction *BinlogTransaction) MarshalBson(buf *bytes2.ChunkedWriter, key string) {
	bson.EncodeOptionalPrefix(buf, bson.Object, key)
	lenWriter := bson.NewLenWriter(buf)

	// []Statement
	{
		bson.EncodePrefix(buf, bson.Array, "Statements")
		lenWriter := bson.NewLenWriter(buf)
		for _i, _v1 := range binlogTransaction.Statements {
			_v1.MarshalBson(buf, bson.Itoa(_i))
		}
		lenWriter.Close()
	}
	bson.EncodeInt64(buf, "Timestamp", binlogTransaction.Timestamp)
	bson.EncodeString(buf, "TransactionID", binlogTransaction.TransactionID)

	lenWriter.Close()
}

// UnmarshalBson bson-decodes into BinlogTransaction.
func (binlogTransaction *BinlogTransaction) UnmarshalBson(buf *bytes.Buffer, kind byte) {
	switch kind {
	case bson.EOO, bson.Object:
		// valid
	case bson.Null:
		return
	default:
		panic(bson.NewBsonError("unexpected kind %v for BinlogTransaction", kind))
	}
	bson.Next(buf, 4)

	for kind := bson.NextByte(buf); kind != bson.EOO; kind = bson.NextByte(buf) {
		switch bson.ReadCString(buf) {
		case "Statements":
			// []Statement
			if kind != bson.Null {
				if kind != bson.Array {
					panic(bson.NewBsonError("unexpected kind %v for binlogTransaction.Statements", kind))
				}
				bson.Next(buf, 4)
				binlogTransaction.Statements = make([]Statement, 0, 8)
				for kind := bson.NextByte(buf); kind != bson.EOO; kind = bson.NextByte(buf) {
					bson.SkipIndex(buf)
					var _v1 Statement
					_v1.UnmarshalBson(buf, kind)
					binlogTransaction.Statements = append(binlogTransaction.Statements, _v1)
				}
			}
		case "Timestamp":
			binlogTransaction.Timestamp = bson.DecodeInt64(buf, kind)
		case "TransactionID":
			binlogTransaction.TransactionID = bson.DecodeString(buf, kind)
		default:
			bson.Skip(buf, kind)
		}
	}
}
