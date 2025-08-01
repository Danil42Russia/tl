// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlCasesBytesTestArray

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tl/tlBuiltinTupleString"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type CasesBytesTestArray struct {
	N   uint32
	Arr []string
}

func (CasesBytesTestArray) TLName() string { return "cases_bytes.testArray" }
func (CasesBytesTestArray) TLTag() uint32  { return 0x3762fb81 }

func (item *CasesBytesTestArray) Reset() {
	item.N = 0
	item.Arr = item.Arr[:0]
}

func (item *CasesBytesTestArray) FillRandom(rg *basictl.RandGenerator) {
	item.N = basictl.RandomUint(rg)
	item.N = rg.LimitValue(item.N)
	tlBuiltinTupleString.BuiltinTupleStringFillRandom(rg, &item.Arr, item.N)
}

func (item *CasesBytesTestArray) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.N); err != nil {
		return w, err
	}
	return tlBuiltinTupleString.BuiltinTupleStringRead(w, &item.Arr, item.N)
}

func (item *CasesBytesTestArray) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w)
}

func (item *CasesBytesTestArray) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.N)
	if w, err = tlBuiltinTupleString.BuiltinTupleStringWrite(w, item.Arr, item.N); err != nil {
		return w, err
	}
	return w, nil
}

func (item *CasesBytesTestArray) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x3762fb81); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *CasesBytesTestArray) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w)
}

func (item *CasesBytesTestArray) WriteBoxed(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, 0x3762fb81)
	return item.Write(w)
}

func (item CasesBytesTestArray) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *CasesBytesTestArray) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propNPresented bool
	var rawArr []byte

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "n":
				if propNPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("cases_bytes.testArray", "n")
				}
				if err := internal.Json2ReadUint32(in, &item.N); err != nil {
					return err
				}
				propNPresented = true
			case "arr":
				if rawArr != nil {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("cases_bytes.testArray", "arr")
				}
				rawArr = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			default:
				return internal.ErrorInvalidJSONExcessElement("cases_bytes.testArray", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propNPresented {
		item.N = 0
	}
	var inArrPointer *basictl.JsonLexer
	inArr := basictl.JsonLexer{Data: rawArr}
	if rawArr != nil {
		inArrPointer = &inArr
	}
	if err := tlBuiltinTupleString.BuiltinTupleStringReadJSON(legacyTypeNames, inArrPointer, &item.Arr, item.N); err != nil {
		return err
	}

	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *CasesBytesTestArray) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w)
}

func (item *CasesBytesTestArray) WriteJSON(w []byte) (_ []byte, err error) {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}
func (item *CasesBytesTestArray) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	backupIndexN := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"n":`...)
	w = basictl.JSONWriteUint32(w, item.N)
	if (item.N != 0) == false {
		w = w[:backupIndexN]
	}
	backupIndexArr := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"arr":`...)
	if w, err = tlBuiltinTupleString.BuiltinTupleStringWriteJSONOpt(tctx, w, item.Arr, item.N); err != nil {
		return w, err
	}
	if (len(item.Arr) != 0) == false {
		w = w[:backupIndexArr]
	}
	return append(w, '}'), nil
}

func (item *CasesBytesTestArray) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *CasesBytesTestArray) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("cases_bytes.testArray", err.Error())
	}
	return nil
}

func (item *CasesBytesTestArray) CalculateLayout(sizes []int) []int {
	sizePosition := len(sizes)
	sizes = append(sizes, 0)

	currentSize := 0
	lastUsedByte := 0

	// calculate layout for item.N
	if item.N != 0 {

		lastUsedByte = 1
		currentSize += 4
	}

	// calculate layout for item.Arr
	currentPosition := len(sizes)
	if len(item.Arr) != 0 {
		sizes = tlBuiltinTupleString.BuiltinTupleStringCalculateLayout(sizes, &item.Arr, item.N)
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

func (item *CasesBytesTestArray) InternalWriteTL2(w []byte, sizes []int) ([]byte, []int) {
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
	// write item.N
	if item.N != 0 {
		serializedSize += 4
		if 4 != 0 {
			currentBlock |= (1 << 1)
			w = basictl.NatWrite(w, item.N)
		}
	}
	// write item.Arr
	if len(item.Arr) != 0 {
		serializedSize += sizes[0]
		if sizes[0] != 0 {
			serializedSize += basictl.TL2CalculateSize(sizes[0])
			currentBlock |= (1 << 2)
			w, sizes = tlBuiltinTupleString.BuiltinTupleStringInternalWriteTL2(w, sizes, &item.Arr, item.N)
		} else {
			sizes = sizes[1:]
		}
	}
	w[currentBlockPosition] = currentBlock
	return w, sizes
}

func (item *CasesBytesTestArray) WriteTL2(w []byte, ctx *basictl.TL2WriteContext) []byte {
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

func (item *CasesBytesTestArray) InternalReadTL2(r []byte) (_ []byte, err error) {
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

	// read item.N
	if block&(1<<1) != 0 {
		if currentR, err = basictl.NatRead(currentR, &item.N); err != nil {
			return currentR, err
		}
	} else {
		item.N = 0
	}

	// read item.Arr
	if block&(1<<2) != 0 {
		if currentR, err = tlBuiltinTupleString.BuiltinTupleStringInternalReadTL2(currentR, &item.Arr, item.N); err != nil {
			return currentR, err
		}
	} else {
		item.Arr = item.Arr[:0]
	}

	return r, nil
}

func (item *CasesBytesTestArray) ReadTL2(r []byte, ctx *basictl.TL2ReadContext) (_ []byte, err error) {
	return item.InternalReadTL2(r)
}

type CasesBytesTestArrayBytes struct {
	N   uint32
	Arr [][]byte
}

func (CasesBytesTestArrayBytes) TLName() string { return "cases_bytes.testArray" }
func (CasesBytesTestArrayBytes) TLTag() uint32  { return 0x3762fb81 }

func (item *CasesBytesTestArrayBytes) Reset() {
	item.N = 0
	item.Arr = item.Arr[:0]
}

func (item *CasesBytesTestArrayBytes) FillRandom(rg *basictl.RandGenerator) {
	item.N = basictl.RandomUint(rg)
	item.N = rg.LimitValue(item.N)
	tlBuiltinTupleString.BuiltinTupleStringBytesFillRandom(rg, &item.Arr, item.N)
}

func (item *CasesBytesTestArrayBytes) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.N); err != nil {
		return w, err
	}
	return tlBuiltinTupleString.BuiltinTupleStringBytesRead(w, &item.Arr, item.N)
}

func (item *CasesBytesTestArrayBytes) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w)
}

func (item *CasesBytesTestArrayBytes) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.N)
	if w, err = tlBuiltinTupleString.BuiltinTupleStringBytesWrite(w, item.Arr, item.N); err != nil {
		return w, err
	}
	return w, nil
}

func (item *CasesBytesTestArrayBytes) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x3762fb81); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *CasesBytesTestArrayBytes) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w)
}

func (item *CasesBytesTestArrayBytes) WriteBoxed(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, 0x3762fb81)
	return item.Write(w)
}

func (item CasesBytesTestArrayBytes) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *CasesBytesTestArrayBytes) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propNPresented bool
	var rawArr []byte

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "n":
				if propNPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("cases_bytes.testArray", "n")
				}
				if err := internal.Json2ReadUint32(in, &item.N); err != nil {
					return err
				}
				propNPresented = true
			case "arr":
				if rawArr != nil {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("cases_bytes.testArray", "arr")
				}
				rawArr = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			default:
				return internal.ErrorInvalidJSONExcessElement("cases_bytes.testArray", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propNPresented {
		item.N = 0
	}
	var inArrPointer *basictl.JsonLexer
	inArr := basictl.JsonLexer{Data: rawArr}
	if rawArr != nil {
		inArrPointer = &inArr
	}
	if err := tlBuiltinTupleString.BuiltinTupleStringBytesReadJSON(legacyTypeNames, inArrPointer, &item.Arr, item.N); err != nil {
		return err
	}

	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *CasesBytesTestArrayBytes) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w)
}

func (item *CasesBytesTestArrayBytes) WriteJSON(w []byte) (_ []byte, err error) {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}
func (item *CasesBytesTestArrayBytes) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	backupIndexN := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"n":`...)
	w = basictl.JSONWriteUint32(w, item.N)
	if (item.N != 0) == false {
		w = w[:backupIndexN]
	}
	backupIndexArr := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"arr":`...)
	if w, err = tlBuiltinTupleString.BuiltinTupleStringBytesWriteJSONOpt(tctx, w, item.Arr, item.N); err != nil {
		return w, err
	}
	if (len(item.Arr) != 0) == false {
		w = w[:backupIndexArr]
	}
	return append(w, '}'), nil
}

func (item *CasesBytesTestArrayBytes) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *CasesBytesTestArrayBytes) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("cases_bytes.testArray", err.Error())
	}
	return nil
}

func (item *CasesBytesTestArrayBytes) CalculateLayout(sizes []int) []int {
	sizePosition := len(sizes)
	sizes = append(sizes, 0)

	currentSize := 0
	lastUsedByte := 0

	// calculate layout for item.N
	if item.N != 0 {

		lastUsedByte = 1
		currentSize += 4
	}

	// calculate layout for item.Arr
	currentPosition := len(sizes)
	if len(item.Arr) != 0 {
		sizes = tlBuiltinTupleString.BuiltinTupleStringBytesCalculateLayout(sizes, &item.Arr, item.N)
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

func (item *CasesBytesTestArrayBytes) InternalWriteTL2(w []byte, sizes []int) ([]byte, []int) {
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
	// write item.N
	if item.N != 0 {
		serializedSize += 4
		if 4 != 0 {
			currentBlock |= (1 << 1)
			w = basictl.NatWrite(w, item.N)
		}
	}
	// write item.Arr
	if len(item.Arr) != 0 {
		serializedSize += sizes[0]
		if sizes[0] != 0 {
			serializedSize += basictl.TL2CalculateSize(sizes[0])
			currentBlock |= (1 << 2)
			w, sizes = tlBuiltinTupleString.BuiltinTupleStringBytesInternalWriteTL2(w, sizes, &item.Arr, item.N)
		} else {
			sizes = sizes[1:]
		}
	}
	w[currentBlockPosition] = currentBlock
	return w, sizes
}

func (item *CasesBytesTestArrayBytes) WriteTL2(w []byte, ctx *basictl.TL2WriteContext) []byte {
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

func (item *CasesBytesTestArrayBytes) InternalReadTL2(r []byte) (_ []byte, err error) {
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

	// read item.N
	if block&(1<<1) != 0 {
		if currentR, err = basictl.NatRead(currentR, &item.N); err != nil {
			return currentR, err
		}
	} else {
		item.N = 0
	}

	// read item.Arr
	if block&(1<<2) != 0 {
		if currentR, err = tlBuiltinTupleString.BuiltinTupleStringBytesInternalReadTL2(currentR, &item.Arr, item.N); err != nil {
			return currentR, err
		}
	} else {
		item.Arr = item.Arr[:0]
	}

	return r, nil
}

func (item *CasesBytesTestArrayBytes) ReadTL2(r []byte, ctx *basictl.TL2ReadContext) (_ []byte, err error) {
	return item.InternalReadTL2(r)
}
