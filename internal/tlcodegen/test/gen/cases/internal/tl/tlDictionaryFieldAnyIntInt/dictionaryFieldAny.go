// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlDictionaryFieldAnyIntInt

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type DictionaryFieldAnyIntInt struct {
	Key   int32
	Value int32
}

func (DictionaryFieldAnyIntInt) TLName() string { return "dictionaryFieldAny" }
func (DictionaryFieldAnyIntInt) TLTag() uint32  { return 0x2c43a65b }

func (item *DictionaryFieldAnyIntInt) Reset() {
	item.Key = 0
	item.Value = 0
}

func (item *DictionaryFieldAnyIntInt) FillRandom(rg *basictl.RandGenerator) {
	item.Key = basictl.RandomInt(rg)
	item.Value = basictl.RandomInt(rg)
}

func (item *DictionaryFieldAnyIntInt) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.IntRead(w, &item.Key); err != nil {
		return w, err
	}
	return basictl.IntRead(w, &item.Value)
}

func (item *DictionaryFieldAnyIntInt) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *DictionaryFieldAnyIntInt) Write(w []byte) []byte {
	w = basictl.IntWrite(w, item.Key)
	w = basictl.IntWrite(w, item.Value)
	return w
}

func (item *DictionaryFieldAnyIntInt) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x2c43a65b); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *DictionaryFieldAnyIntInt) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *DictionaryFieldAnyIntInt) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x2c43a65b)
	return item.Write(w)
}

func (item DictionaryFieldAnyIntInt) String() string {
	return string(item.WriteJSON(nil))
}

func (item *DictionaryFieldAnyIntInt) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propKeyPresented bool
	var propValuePresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "key":
				if propKeyPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("dictionaryFieldAny", "key")
				}
				if err := internal.Json2ReadInt32(in, &item.Key); err != nil {
					return err
				}
				propKeyPresented = true
			case "value":
				if propValuePresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("dictionaryFieldAny", "value")
				}
				if err := internal.Json2ReadInt32(in, &item.Value); err != nil {
					return err
				}
				propValuePresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("dictionaryFieldAny", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propKeyPresented {
		item.Key = 0
	}
	if !propValuePresented {
		item.Value = 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *DictionaryFieldAnyIntInt) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w), nil
}

func (item *DictionaryFieldAnyIntInt) WriteJSON(w []byte) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}
func (item *DictionaryFieldAnyIntInt) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) []byte {
	w = append(w, '{')
	backupIndexKey := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"key":`...)
	w = basictl.JSONWriteInt32(w, item.Key)
	if (item.Key != 0) == false {
		w = w[:backupIndexKey]
	}
	backupIndexValue := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"value":`...)
	w = basictl.JSONWriteInt32(w, item.Value)
	if (item.Value != 0) == false {
		w = w[:backupIndexValue]
	}
	return append(w, '}')
}

func (item *DictionaryFieldAnyIntInt) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *DictionaryFieldAnyIntInt) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("dictionaryFieldAny", err.Error())
	}
	return nil
}

func (item *DictionaryFieldAnyIntInt) CalculateLayout(sizes []int) []int {
	sizePosition := len(sizes)
	sizes = append(sizes, 0)

	currentSize := 0
	lastUsedByte := 0

	// calculate layout for item.Key
	if item.Key != 0 {

		lastUsedByte = 1
		currentSize += 4
	}

	// calculate layout for item.Value
	if item.Value != 0 {

		lastUsedByte = 1
		currentSize += 4
	}

	// append byte for each section until last mentioned field
	if lastUsedByte != 0 {
		currentSize += lastUsedByte
	} else {
		// remove unused values
		sizes = sizes[:sizePosition+1]
	}
	sizes[sizePosition] = currentSize
	return sizes
}

func (item *DictionaryFieldAnyIntInt) InternalWriteTL2(w []byte, sizes []int) ([]byte, []int) {
	currentSize := sizes[0]
	sizes = sizes[1:]

	serializedSize := 0

	w = basictl.TL2WriteSize(w, currentSize)
	if currentSize == 0 {
		return w, sizes
	}

	var currentBlock byte
	currentBlockPosition := len(w)
	w = append(w, 0)
	serializedSize += 1
	// write item.Key
	if item.Key != 0 {
		serializedSize += 4
		if 4 != 0 {
			currentBlock |= (1 << 1)
			w = basictl.IntWrite(w, item.Key)
		}
	}
	// write item.Value
	if item.Value != 0 {
		serializedSize += 4
		if 4 != 0 {
			currentBlock |= (1 << 2)
			w = basictl.IntWrite(w, item.Value)
		}
	}
	w[currentBlockPosition] = currentBlock
	return w, sizes
}

func (item *DictionaryFieldAnyIntInt) WriteTL2(w []byte, ctx *basictl.TL2WriteContext) []byte {
	var sizes []int
	if ctx != nil {
		sizes = ctx.SizeBuffer
	}
	sizes = item.CalculateLayout(sizes[:0])
	w, _ = item.InternalWriteTL2(w, sizes)
	if ctx != nil {
		ctx.SizeBuffer = sizes[:0]
	}
	return w
}

func (item *DictionaryFieldAnyIntInt) InternalReadTL2(r []byte) (_ []byte, err error) {
	currentSize := 0
	if r, currentSize, err = basictl.TL2ParseSize(r); err != nil {
		return r, err
	}
	if len(r) < currentSize {
		return r, basictl.TL2Error("not enough data: expected %d, got %d", currentSize, len(r))
	}

	currentR := r[:currentSize]
	r = r[currentSize:]

	if currentSize == 0 {
		item.Reset()
		return r, nil
	}
	var block byte
	if currentR, err = basictl.ByteReadTL2(currentR, &block); err != nil {
		return currentR, err
	}
	// read No of constructor
	if block&1 != 0 {
		var index int
		if currentR, err = basictl.TL2ReadSize(currentR, &index); err != nil {
			return currentR, err
		}
		if index != 0 {
			// unknown cases for current type
			item.Reset()
			return r, nil
		}
	}

	// read item.Key
	if block&(1<<1) != 0 {
		if currentR, err = basictl.IntRead(currentR, &item.Key); err != nil {
			return currentR, err
		}
	} else {
		item.Key = 0
	}

	// read item.Value
	if block&(1<<2) != 0 {
		if currentR, err = basictl.IntRead(currentR, &item.Value); err != nil {
			return currentR, err
		}
	} else {
		item.Value = 0
	}

	return r, nil
}

func (item *DictionaryFieldAnyIntInt) ReadTL2(r []byte, ctx *basictl.TL2ReadContext) (_ []byte, err error) {
	return item.InternalReadTL2(r)
}
