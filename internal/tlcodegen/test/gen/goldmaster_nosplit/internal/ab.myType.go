// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package internal

import (
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite

type AbMyType struct {
	X int32
}

func (AbMyType) TLName() string { return "ab.myType" }
func (AbMyType) TLTag() uint32  { return 0xe0e96c86 }

func (item *AbMyType) Reset() {
	item.X = 0
}

func (item *AbMyType) FillRandom(rg *basictl.RandGenerator) {
	item.X = basictl.RandomInt(rg)
}

func (item *AbMyType) Read(w []byte) (_ []byte, err error) {
	return basictl.IntRead(w, &item.X)
}

func (item *AbMyType) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *AbMyType) Write(w []byte) []byte {
	w = basictl.IntWrite(w, item.X)
	return w
}

func (item *AbMyType) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xe0e96c86); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *AbMyType) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *AbMyType) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xe0e96c86)
	return item.Write(w)
}

func (item AbMyType) String() string {
	return string(item.WriteJSON(nil))
}

func (item *AbMyType) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propXPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "x":
				if propXPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("ab.myType", "x")
				}
				if err := Json2ReadInt32(in, &item.X); err != nil {
					return err
				}
				propXPresented = true
			default:
				return ErrorInvalidJSONExcessElement("ab.myType", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propXPresented {
		item.X = 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *AbMyType) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w), nil
}

func (item *AbMyType) WriteJSON(w []byte) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}
func (item *AbMyType) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) []byte {
	w = append(w, '{')
	backupIndexX := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"x":`...)
	w = basictl.JSONWriteInt32(w, item.X)
	if (item.X != 0) == false {
		w = w[:backupIndexX]
	}
	return append(w, '}')
}

func (item *AbMyType) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *AbMyType) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("ab.myType", err.Error())
	}
	return nil
}

func (item *AbMyType) CalculateLayout(sizes []int) []int {
	sizePosition := len(sizes)
	sizes = append(sizes, 0)

	currentSize := 0
	lastUsedByte := 0

	// calculate layout for item.X
	if item.X != 0 {

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

func (item *AbMyType) InternalWriteTL2(w []byte, sizes []int) ([]byte, []int) {
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
	// write item.X
	if item.X != 0 {
		serializedSize += 4
		if 4 != 0 {
			currentBlock |= (1 << 1)
			w = basictl.IntWrite(w, item.X)
		}
	}
	w[currentBlockPosition] = currentBlock
	return w, sizes
}

func (item *AbMyType) WriteTL2(w []byte, ctx *basictl.TL2WriteContext) []byte {
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

func (item *AbMyType) InternalReadTL2(r []byte) (_ []byte, err error) {
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

	// read item.X
	if block&(1<<1) != 0 {
		if currentR, err = basictl.IntRead(currentR, &item.X); err != nil {
			return currentR, err
		}
	} else {
		item.X = 0
	}

	return r, nil
}

func (item *AbMyType) ReadTL2(r []byte, ctx *basictl.TL2ReadContext) (_ []byte, err error) {
	return item.InternalReadTL2(r)
}

type AbMyTypeBoxedMaybe struct {
	Value AbMyType // not deterministic if !Ok
	Ok    bool
}

func (item *AbMyTypeBoxedMaybe) Reset() {
	item.Ok = false
}
func (item *AbMyTypeBoxedMaybe) FillRandom(rg *basictl.RandGenerator) {
	if basictl.RandomUint(rg)&1 == 1 {
		item.Ok = true
		item.Value.FillRandom(rg)
	} else {
		item.Ok = false
	}
}

func (item *AbMyTypeBoxedMaybe) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.ReadBool(w, &item.Ok, 0x27930a7b, 0x3f9c8ef8); err != nil {
		return w, err
	}
	if item.Ok {
		return item.Value.ReadBoxed(w)
	}
	return w, nil
}

func (item *AbMyTypeBoxedMaybe) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *AbMyTypeBoxedMaybe) WriteBoxed(w []byte) []byte {
	if item.Ok {
		w = basictl.NatWrite(w, 0x3f9c8ef8)
		return item.Value.WriteBoxed(w)
	}
	return basictl.NatWrite(w, 0x27930a7b)
}

func (item *AbMyTypeBoxedMaybe) CalculateLayout(sizes []int) []int {
	sizePosition := len(sizes)
	sizes = append(sizes, 0)
	if item.Ok {
		sizes[sizePosition] += 1
		sizes[sizePosition] += basictl.TL2CalculateSize(1)
		currentPosition := len(sizes)
		sizes = item.Value.CalculateLayout(sizes)
		if sizes[currentPosition] != 0 {
			sizes[sizePosition] += sizes[currentPosition]
			sizes[sizePosition] += basictl.TL2CalculateSize(sizes[currentPosition])
		}
	}
	return sizes
}

func (item *AbMyTypeBoxedMaybe) InternalWriteTL2(w []byte, sizes []int) ([]byte, []int) {
	currentSize := sizes[0]
	sizes = sizes[1:]

	w = basictl.TL2WriteSize(w, currentSize)
	if currentSize == 0 {
		return w, sizes
	}

	if item.Ok {
		currentPosition := len(w)
		w = append(w, 1)
		w = basictl.TL2WriteSize(w, 1)
		if sizes[0] != 0 {
			w[currentPosition] |= (1 << 1)
			w, sizes = item.Value.InternalWriteTL2(w, sizes)
		} else {
			sizes = sizes[1:]
		}
	}
	return w, sizes
}

func (item *AbMyTypeBoxedMaybe) InternalReadTL2(r []byte) (_ []byte, err error) {
	saveR := r
	currentSize := 0
	if r, currentSize, err = basictl.TL2ParseSize(r); err != nil {
		return r, err
	}
	shift := currentSize + basictl.TL2CalculateSize(currentSize)

	if currentSize == 0 {
		item.Ok = false
	} else {
		var block byte
		if r, err = basictl.ByteReadTL2(r, &block); err != nil {
			return r, err
		}
		if block&1 == 0 {
			return r, basictl.TL2Error("must have constructor bytes")
		}
		var index int
		if r, index, err = basictl.TL2ParseSize(r); err != nil {
			return r, err
		}
		if index != 1 {
			return r, basictl.TL2Error("expected 1")
		}
		item.Ok = true
		if block&(1<<1) != 0 {
			if r, err = item.Value.InternalReadTL2(r); err != nil {
				return r, err
			}
		} else {
			item.Value.Reset()
		}
	}
	if len(saveR) < len(r)+shift {
		r = saveR[shift:]
	}
	return r, nil
}

func (item *AbMyTypeBoxedMaybe) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	_ok, _jvalue, err := Json2ReadMaybe("Maybe", in)
	if err != nil {
		return err
	}
	item.Ok = _ok
	if _ok {
		var in2Pointer *basictl.JsonLexer
		if _jvalue != nil {
			in2 := basictl.JsonLexer{Data: _jvalue}
			in2Pointer = &in2
		}
		if err := item.Value.ReadJSON(legacyTypeNames, in2Pointer); err != nil {
			return err
		}
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *AbMyTypeBoxedMaybe) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w), nil
}

func (item *AbMyTypeBoxedMaybe) WriteJSON(w []byte) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}
func (item *AbMyTypeBoxedMaybe) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) []byte {
	if !item.Ok {
		return append(w, "{}"...)
	}
	w = append(w, `{"ok":true`...)
	w = append(w, `,"value":`...)
	w = item.Value.WriteJSONOpt(tctx, w)
	return append(w, '}')
}

func (item AbMyTypeBoxedMaybe) String() string {
	return string(item.WriteJSON(nil))
}

type AbMyTypeMaybe struct {
	Value AbMyType // not deterministic if !Ok
	Ok    bool
}

func (item *AbMyTypeMaybe) Reset() {
	item.Ok = false
}
func (item *AbMyTypeMaybe) FillRandom(rg *basictl.RandGenerator) {
	if basictl.RandomUint(rg)&1 == 1 {
		item.Ok = true
		item.Value.FillRandom(rg)
	} else {
		item.Ok = false
	}
}

func (item *AbMyTypeMaybe) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.ReadBool(w, &item.Ok, 0x27930a7b, 0x3f9c8ef8); err != nil {
		return w, err
	}
	if item.Ok {
		return item.Value.Read(w)
	}
	return w, nil
}

func (item *AbMyTypeMaybe) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *AbMyTypeMaybe) WriteBoxed(w []byte) []byte {
	if item.Ok {
		w = basictl.NatWrite(w, 0x3f9c8ef8)
		return item.Value.Write(w)
	}
	return basictl.NatWrite(w, 0x27930a7b)
}

func (item *AbMyTypeMaybe) CalculateLayout(sizes []int) []int {
	sizePosition := len(sizes)
	sizes = append(sizes, 0)
	if item.Ok {
		sizes[sizePosition] += 1
		sizes[sizePosition] += basictl.TL2CalculateSize(1)
		currentPosition := len(sizes)
		sizes = item.Value.CalculateLayout(sizes)
		if sizes[currentPosition] != 0 {
			sizes[sizePosition] += sizes[currentPosition]
			sizes[sizePosition] += basictl.TL2CalculateSize(sizes[currentPosition])
		}
	}
	return sizes
}

func (item *AbMyTypeMaybe) InternalWriteTL2(w []byte, sizes []int) ([]byte, []int) {
	currentSize := sizes[0]
	sizes = sizes[1:]

	w = basictl.TL2WriteSize(w, currentSize)
	if currentSize == 0 {
		return w, sizes
	}

	if item.Ok {
		currentPosition := len(w)
		w = append(w, 1)
		w = basictl.TL2WriteSize(w, 1)
		if sizes[0] != 0 {
			w[currentPosition] |= (1 << 1)
			w, sizes = item.Value.InternalWriteTL2(w, sizes)
		} else {
			sizes = sizes[1:]
		}
	}
	return w, sizes
}

func (item *AbMyTypeMaybe) InternalReadTL2(r []byte) (_ []byte, err error) {
	saveR := r
	currentSize := 0
	if r, currentSize, err = basictl.TL2ParseSize(r); err != nil {
		return r, err
	}
	shift := currentSize + basictl.TL2CalculateSize(currentSize)

	if currentSize == 0 {
		item.Ok = false
	} else {
		var block byte
		if r, err = basictl.ByteReadTL2(r, &block); err != nil {
			return r, err
		}
		if block&1 == 0 {
			return r, basictl.TL2Error("must have constructor bytes")
		}
		var index int
		if r, index, err = basictl.TL2ParseSize(r); err != nil {
			return r, err
		}
		if index != 1 {
			return r, basictl.TL2Error("expected 1")
		}
		item.Ok = true
		if block&(1<<1) != 0 {
			if r, err = item.Value.InternalReadTL2(r); err != nil {
				return r, err
			}
		} else {
			item.Value.Reset()
		}
	}
	if len(saveR) < len(r)+shift {
		r = saveR[shift:]
	}
	return r, nil
}

func (item *AbMyTypeMaybe) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	_ok, _jvalue, err := Json2ReadMaybe("Maybe", in)
	if err != nil {
		return err
	}
	item.Ok = _ok
	if _ok {
		var in2Pointer *basictl.JsonLexer
		if _jvalue != nil {
			in2 := basictl.JsonLexer{Data: _jvalue}
			in2Pointer = &in2
		}
		if err := item.Value.ReadJSON(legacyTypeNames, in2Pointer); err != nil {
			return err
		}
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *AbMyTypeMaybe) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w), nil
}

func (item *AbMyTypeMaybe) WriteJSON(w []byte) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}
func (item *AbMyTypeMaybe) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) []byte {
	if !item.Ok {
		return append(w, "{}"...)
	}
	w = append(w, `{"ok":true`...)
	w = append(w, `,"value":`...)
	w = item.Value.WriteJSONOpt(tctx, w)
	return append(w, '}')
}

func (item AbMyTypeMaybe) String() string {
	return string(item.WriteJSON(nil))
}
