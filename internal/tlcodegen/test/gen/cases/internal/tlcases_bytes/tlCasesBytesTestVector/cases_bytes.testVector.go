// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlCasesBytesTestVector

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tl/tlBuiltinVectorString"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type CasesBytesTestVector struct {
	Arr []string
}

func (CasesBytesTestVector) TLName() string { return "cases_bytes.testVector" }
func (CasesBytesTestVector) TLTag() uint32  { return 0x3647c8ae }

func (item *CasesBytesTestVector) Reset() {
	item.Arr = item.Arr[:0]
}

func (item *CasesBytesTestVector) FillRandom(rg *basictl.RandGenerator) {
	tlBuiltinVectorString.BuiltinVectorStringFillRandom(rg, &item.Arr)
}

func (item *CasesBytesTestVector) Read(w []byte) (_ []byte, err error) {
	return tlBuiltinVectorString.BuiltinVectorStringRead(w, &item.Arr)
}

func (item *CasesBytesTestVector) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *CasesBytesTestVector) Write(w []byte) []byte {
	w = tlBuiltinVectorString.BuiltinVectorStringWrite(w, item.Arr)
	return w
}

func (item *CasesBytesTestVector) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x3647c8ae); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *CasesBytesTestVector) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *CasesBytesTestVector) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x3647c8ae)
	return item.Write(w)
}

func (item CasesBytesTestVector) String() string {
	return string(item.WriteJSON(nil))
}

func (item *CasesBytesTestVector) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propArrPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "arr":
				if propArrPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("cases_bytes.testVector", "arr")
				}
				if err := tlBuiltinVectorString.BuiltinVectorStringReadJSON(legacyTypeNames, in, &item.Arr); err != nil {
					return err
				}
				propArrPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("cases_bytes.testVector", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propArrPresented {
		item.Arr = item.Arr[:0]
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *CasesBytesTestVector) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w), nil
}

func (item *CasesBytesTestVector) WriteJSON(w []byte) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}
func (item *CasesBytesTestVector) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) []byte {
	w = append(w, '{')
	backupIndexArr := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"arr":`...)
	w = tlBuiltinVectorString.BuiltinVectorStringWriteJSONOpt(tctx, w, item.Arr)
	if (len(item.Arr) != 0) == false {
		w = w[:backupIndexArr]
	}
	return append(w, '}')
}

func (item *CasesBytesTestVector) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *CasesBytesTestVector) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("cases_bytes.testVector", err.Error())
	}
	return nil
}

func (item *CasesBytesTestVector) CalculateLayout(sizes []int) []int {
	sizePosition := len(sizes)
	sizes = append(sizes, 0)

	currentSize := 0
	lastUsedByte := 0

	// calculate layout for item.Arr
	currentPosition := len(sizes)
	if len(item.Arr) != 0 {
		sizes = tlBuiltinVectorString.BuiltinVectorStringCalculateLayout(sizes, &item.Arr)
		if sizes[currentPosition] != 0 {
			lastUsedByte = 1
			currentSize += sizes[currentPosition]
			currentSize += basictl.TL2CalculateSize(sizes[currentPosition])
		} else {
			sizes = sizes[:currentPosition+1]
		}
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

func (item *CasesBytesTestVector) InternalWriteTL2(w []byte, sizes []int) ([]byte, []int) {
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
	// write item.Arr
	if len(item.Arr) != 0 {
		serializedSize += sizes[0]
		if sizes[0] != 0 {
			serializedSize += basictl.TL2CalculateSize(sizes[0])
			currentBlock |= (1 << 1)
			w, sizes = tlBuiltinVectorString.BuiltinVectorStringInternalWriteTL2(w, sizes, &item.Arr)
		} else {
			sizes = sizes[1:]
		}
	}
	w[currentBlockPosition] = currentBlock
	return w, sizes
}

func (item *CasesBytesTestVector) WriteTL2(w []byte, ctx *basictl.TL2WriteContext) []byte {
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

func (item *CasesBytesTestVector) InternalReadTL2(r []byte) (_ []byte, err error) {
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

	// read item.Arr
	if block&(1<<1) != 0 {
		if currentR, err = tlBuiltinVectorString.BuiltinVectorStringInternalReadTL2(currentR, &item.Arr); err != nil {
			return currentR, err
		}
	} else {
		item.Arr = item.Arr[:0]
	}

	return r, nil
}

func (item *CasesBytesTestVector) ReadTL2(r []byte, ctx *basictl.TL2ReadContext) (_ []byte, err error) {
	return item.InternalReadTL2(r)
}

type CasesBytesTestVectorBytes struct {
	Arr [][]byte
}

func (CasesBytesTestVectorBytes) TLName() string { return "cases_bytes.testVector" }
func (CasesBytesTestVectorBytes) TLTag() uint32  { return 0x3647c8ae }

func (item *CasesBytesTestVectorBytes) Reset() {
	item.Arr = item.Arr[:0]
}

func (item *CasesBytesTestVectorBytes) FillRandom(rg *basictl.RandGenerator) {
	tlBuiltinVectorString.BuiltinVectorStringBytesFillRandom(rg, &item.Arr)
}

func (item *CasesBytesTestVectorBytes) Read(w []byte) (_ []byte, err error) {
	return tlBuiltinVectorString.BuiltinVectorStringBytesRead(w, &item.Arr)
}

func (item *CasesBytesTestVectorBytes) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *CasesBytesTestVectorBytes) Write(w []byte) []byte {
	w = tlBuiltinVectorString.BuiltinVectorStringBytesWrite(w, item.Arr)
	return w
}

func (item *CasesBytesTestVectorBytes) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x3647c8ae); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *CasesBytesTestVectorBytes) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *CasesBytesTestVectorBytes) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x3647c8ae)
	return item.Write(w)
}

func (item CasesBytesTestVectorBytes) String() string {
	return string(item.WriteJSON(nil))
}

func (item *CasesBytesTestVectorBytes) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propArrPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "arr":
				if propArrPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("cases_bytes.testVector", "arr")
				}
				if err := tlBuiltinVectorString.BuiltinVectorStringBytesReadJSON(legacyTypeNames, in, &item.Arr); err != nil {
					return err
				}
				propArrPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("cases_bytes.testVector", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propArrPresented {
		item.Arr = item.Arr[:0]
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *CasesBytesTestVectorBytes) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w), nil
}

func (item *CasesBytesTestVectorBytes) WriteJSON(w []byte) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}
func (item *CasesBytesTestVectorBytes) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) []byte {
	w = append(w, '{')
	backupIndexArr := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"arr":`...)
	w = tlBuiltinVectorString.BuiltinVectorStringBytesWriteJSONOpt(tctx, w, item.Arr)
	if (len(item.Arr) != 0) == false {
		w = w[:backupIndexArr]
	}
	return append(w, '}')
}

func (item *CasesBytesTestVectorBytes) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *CasesBytesTestVectorBytes) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("cases_bytes.testVector", err.Error())
	}
	return nil
}

func (item *CasesBytesTestVectorBytes) CalculateLayout(sizes []int) []int {
	sizePosition := len(sizes)
	sizes = append(sizes, 0)

	currentSize := 0
	lastUsedByte := 0

	// calculate layout for item.Arr
	currentPosition := len(sizes)
	if len(item.Arr) != 0 {
		sizes = tlBuiltinVectorString.BuiltinVectorStringBytesCalculateLayout(sizes, &item.Arr)
		if sizes[currentPosition] != 0 {
			lastUsedByte = 1
			currentSize += sizes[currentPosition]
			currentSize += basictl.TL2CalculateSize(sizes[currentPosition])
		} else {
			sizes = sizes[:currentPosition+1]
		}
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

func (item *CasesBytesTestVectorBytes) InternalWriteTL2(w []byte, sizes []int) ([]byte, []int) {
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
	// write item.Arr
	if len(item.Arr) != 0 {
		serializedSize += sizes[0]
		if sizes[0] != 0 {
			serializedSize += basictl.TL2CalculateSize(sizes[0])
			currentBlock |= (1 << 1)
			w, sizes = tlBuiltinVectorString.BuiltinVectorStringBytesInternalWriteTL2(w, sizes, &item.Arr)
		} else {
			sizes = sizes[1:]
		}
	}
	w[currentBlockPosition] = currentBlock
	return w, sizes
}

func (item *CasesBytesTestVectorBytes) WriteTL2(w []byte, ctx *basictl.TL2WriteContext) []byte {
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

func (item *CasesBytesTestVectorBytes) InternalReadTL2(r []byte) (_ []byte, err error) {
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

	// read item.Arr
	if block&(1<<1) != 0 {
		if currentR, err = tlBuiltinVectorString.BuiltinVectorStringBytesInternalReadTL2(currentR, &item.Arr); err != nil {
			return currentR, err
		}
	} else {
		item.Arr = item.Arr[:0]
	}

	return r, nil
}

func (item *CasesBytesTestVectorBytes) ReadTL2(r []byte, ctx *basictl.TL2ReadContext) (_ []byte, err error) {
	return item.InternalReadTL2(r)
}
