// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package cycle_463e8ce0a74f5ec3e904ecdc85b1b857

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

var _AMyUnion = [2]internal.UnionElement{
	{TLTag: 0xa7662843, TLName: "a.uNionA", TLString: "a.uNionA#a7662843"},
	{TLTag: 0xdf61f632, TLName: "au.nionA", TLString: "au.nionA#df61f632"},
}

type AMyUnion struct {
	valueUNionA AUNionA
	valueNionA  AuNionA
	index       int
}

func (item AMyUnion) TLName() string { return _AMyUnion[item.index].TLName }
func (item AMyUnion) TLTag() uint32  { return _AMyUnion[item.index].TLTag }

func (item *AMyUnion) Reset() { item.ResetToUNionA() }
func (item *AMyUnion) FillRandom(rg *basictl.RandGenerator) {
	index := basictl.RandomUint(rg) % 2
	switch index {
	case 0:
		item.index = 0
		item.valueUNionA.FillRandom(rg)
	case 1:
		item.index = 1
		item.valueNionA.FillRandom(rg)
	default:
	}
}

func (item *AMyUnion) IsUNionA() bool { return item.index == 0 }

func (item *AMyUnion) AsUNionA() (*AUNionA, bool) {
	if item.index != 0 {
		return nil, false
	}
	return &item.valueUNionA, true
}
func (item *AMyUnion) ResetToUNionA() *AUNionA {
	item.index = 0
	item.valueUNionA.Reset()
	return &item.valueUNionA
}
func (item *AMyUnion) SetUNionA(value AUNionA) {
	item.index = 0
	item.valueUNionA = value
}

func (item *AMyUnion) IsNionA() bool { return item.index == 1 }

func (item *AMyUnion) AsNionA() (*AuNionA, bool) {
	if item.index != 1 {
		return nil, false
	}
	return &item.valueNionA, true
}
func (item *AMyUnion) ResetToNionA() *AuNionA {
	item.index = 1
	item.valueNionA.Reset()
	return &item.valueNionA
}
func (item *AMyUnion) SetNionA(value AuNionA) {
	item.index = 1
	item.valueNionA = value
}

func (item *AMyUnion) ReadBoxed(w []byte) (_ []byte, err error) {
	var tag uint32
	if w, err = basictl.NatRead(w, &tag); err != nil {
		return w, err
	}
	switch tag {
	case 0xa7662843:
		item.index = 0
		return item.valueUNionA.Read(w)
	case 0xdf61f632:
		item.index = 1
		return item.valueNionA.Read(w)
	default:
		return w, internal.ErrorInvalidUnionTag("a.MyUnion", tag)
	}
}

func (item *AMyUnion) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *AMyUnion) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, _AMyUnion[item.index].TLTag)
	switch item.index {
	case 0:
		w = item.valueUNionA.Write(w)
	case 1:
		w = item.valueNionA.Write(w)
	}
	return w
}

func (item *AMyUnion) CalculateLayout(sizes []int) []int {
	switch item.index {
	case 0:
		sizes = item.valueUNionA.CalculateLayout(sizes)
	case 1:
		sizes = item.valueNionA.CalculateLayout(sizes)
	}
	return sizes
}

func (item *AMyUnion) InternalWriteTL2(w []byte, sizes []int) ([]byte, []int) {
	switch item.index {
	case 0:
		w, sizes = item.valueUNionA.InternalWriteTL2(w, sizes)
	case 1:
		w, sizes = item.valueNionA.InternalWriteTL2(w, sizes)
	}
	return w, sizes
}

func (item *AMyUnion) InternalReadTL2(r []byte) (_ []byte, err error) {
	currentSize := 0
	if r, currentSize, err = basictl.TL2ParseSize(r); err != nil {
		return r, err
	}

	currentR := r[:currentSize]
	r = r[currentSize:]

	var block byte
	if currentSize == 0 {
		item.index = 0
	} else {
		if currentR, err = basictl.ByteReadTL2(currentR, &block); err != nil {
			return r, err
		}
		if (block & 1) != 0 {
			if currentR, item.index, err = basictl.TL2ParseSize(currentR); err != nil {
				return r, err
			}
		} else {
			item.index = 0
		}
	}
	switch item.index {
	case 0:
		if currentR, err = item.valueUNionA.InternalReadTL2(currentR, block); err != nil {
			return currentR, err
		}
	case 1:
		if currentR, err = item.valueNionA.InternalReadTL2(currentR, block); err != nil {
			return currentR, err
		}
	}
	return r, nil
}
func (item *AMyUnion) WriteTL2(w []byte, ctx *basictl.TL2WriteContext) []byte {
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

func (item *AMyUnion) ReadTL2(r []byte, ctx *basictl.TL2ReadContext) ([]byte, error) {
	return item.InternalReadTL2(r)
}

func (item *AMyUnion) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	_tag, _value, err := internal.Json2ReadUnion("a.MyUnion", in)
	if err != nil {
		return err
	}
	switch _tag {
	case "a.uNionA#a7662843", "a.uNionA", "#a7662843":
		if !legacyTypeNames && _tag == "a.uNionA#a7662843" {
			return internal.ErrorInvalidUnionLegacyTagJSON("a.MyUnion", "a.uNionA#a7662843")
		}
		item.index = 0
		var in2Pointer *basictl.JsonLexer
		if _value != nil {
			in2 := basictl.JsonLexer{Data: _value}
			in2Pointer = &in2
		}
		if err := item.valueUNionA.ReadJSON(legacyTypeNames, in2Pointer); err != nil {
			return err
		}
	case "au.nionA#df61f632", "au.nionA", "#df61f632":
		if !legacyTypeNames && _tag == "au.nionA#df61f632" {
			return internal.ErrorInvalidUnionLegacyTagJSON("a.MyUnion", "au.nionA#df61f632")
		}
		item.index = 1
		var in2Pointer *basictl.JsonLexer
		if _value != nil {
			in2 := basictl.JsonLexer{Data: _value}
			in2Pointer = &in2
		}
		if err := item.valueNionA.ReadJSON(legacyTypeNames, in2Pointer); err != nil {
			return err
		}
	default:
		return internal.ErrorInvalidUnionTagJSON("a.MyUnion", _tag)
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *AMyUnion) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) ([]byte, error) {
	return item.WriteJSONOpt(tctx, w), nil
}

func (item *AMyUnion) WriteJSON(w []byte) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}
func (item *AMyUnion) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) []byte {
	switch item.index {
	case 0:
		if tctx.LegacyTypeNames {
			w = append(w, `{"type":"a.uNionA#a7662843"`...)
		} else {
			w = append(w, `{"type":"a.uNionA"`...)
		}
		w = append(w, `,"value":`...)
		w = item.valueUNionA.WriteJSONOpt(tctx, w)
		return append(w, '}')
	case 1:
		if tctx.LegacyTypeNames {
			w = append(w, `{"type":"au.nionA#df61f632"`...)
		} else {
			w = append(w, `{"type":"au.nionA"`...)
		}
		w = append(w, `,"value":`...)
		w = item.valueNionA.WriteJSONOpt(tctx, w)
		return append(w, '}')
	default: // Impossible due to panic above
		return w
	}
}

func (item AMyUnion) String() string {
	return string(item.WriteJSON(nil))
}

func (item *AMyUnion) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *AMyUnion) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("a.MyUnion", err.Error())
	}
	return nil
}

func (item AUNionA) AsUnion() AMyUnion {
	var ret AMyUnion
	ret.SetUNionA(item)
	return ret
}

type AUNionA struct {
	A int32
}

func (AUNionA) TLName() string { return "a.uNionA" }
func (AUNionA) TLTag() uint32  { return 0xa7662843 }

func (item *AUNionA) Reset() {
	item.A = 0
}

func (item *AUNionA) FillRandom(rg *basictl.RandGenerator) {
	item.A = basictl.RandomInt(rg)
}

func (item *AUNionA) Read(w []byte) (_ []byte, err error) {
	return basictl.IntRead(w, &item.A)
}

func (item *AUNionA) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *AUNionA) Write(w []byte) []byte {
	w = basictl.IntWrite(w, item.A)
	return w
}

func (item *AUNionA) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xa7662843); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *AUNionA) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *AUNionA) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xa7662843)
	return item.Write(w)
}

func (item AUNionA) String() string {
	return string(item.WriteJSON(nil))
}

func (item *AUNionA) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propAPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "a":
				if propAPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("a.uNionA", "a")
				}
				if err := internal.Json2ReadInt32(in, &item.A); err != nil {
					return err
				}
				propAPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("a.uNionA", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propAPresented {
		item.A = 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *AUNionA) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w), nil
}

func (item *AUNionA) WriteJSON(w []byte) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}
func (item *AUNionA) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) []byte {
	w = append(w, '{')
	backupIndexA := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"a":`...)
	w = basictl.JSONWriteInt32(w, item.A)
	if (item.A != 0) == false {
		w = w[:backupIndexA]
	}
	return append(w, '}')
}

func (item *AUNionA) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *AUNionA) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("a.uNionA", err.Error())
	}
	return nil
}

func (item *AUNionA) CalculateLayout(sizes []int) []int {
	sizePosition := len(sizes)
	sizes = append(sizes, 0)

	currentSize := 0
	lastUsedByte := 0

	// calculate layout for item.A
	if item.A != 0 {

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

func (item *AUNionA) InternalWriteTL2(w []byte, sizes []int) ([]byte, []int) {
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
	// write item.A
	if item.A != 0 {
		serializedSize += 4
		if 4 != 0 {
			currentBlock |= (1 << 1)
			w = basictl.IntWrite(w, item.A)
		}
	}
	w[currentBlockPosition] = currentBlock
	return w, sizes
}

func (item *AUNionA) WriteTL2(w []byte, ctx *basictl.TL2WriteContext) []byte {
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

func (item *AUNionA) InternalReadTL2(r []byte, block byte) (_ []byte, err error) {
	currentR := r

	// read item.A
	if block&(1<<1) != 0 {
		if currentR, err = basictl.IntRead(currentR, &item.A); err != nil {
			return currentR, err
		}
	} else {
		item.A = 0
	}

	return r, nil
}

func (item *AUNionA) ReadTL2(r []byte, ctx *basictl.TL2ReadContext) (_ []byte, err error) {
	currentSize := 0
	if r, currentSize, err = basictl.TL2ParseSize(r); err != nil {
		return r, err
	}

	currentR := r[:currentSize]
	r = r[currentSize:]

	var block byte
	var index int
	if currentSize == 0 {
		index = 0
	} else {
		if currentR, err = basictl.ByteReadTL2(currentR, &block); err != nil {
			return r, err
		}
		if (block & 1) != 0 {
			if currentR, index, err = basictl.TL2ParseSize(currentR); err != nil {
				return r, err
			}
		} else {
			index = 0
		}
	}
	if index != 0 {
		return r, basictl.TL2Error("unexpected constructor number %d, instead of %d", index, 0)
	}
	_, err = item.InternalReadTL2(currentR, block)
	return r, err
}

func (item AuNionA) AsUnion() AMyUnion {
	var ret AMyUnion
	ret.SetNionA(item)
	return ret
}

type AuNionA struct {
	B int32
}

func (AuNionA) TLName() string { return "au.nionA" }
func (AuNionA) TLTag() uint32  { return 0xdf61f632 }

func (item *AuNionA) Reset() {
	item.B = 0
}

func (item *AuNionA) FillRandom(rg *basictl.RandGenerator) {
	item.B = basictl.RandomInt(rg)
}

func (item *AuNionA) Read(w []byte) (_ []byte, err error) {
	return basictl.IntRead(w, &item.B)
}

func (item *AuNionA) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *AuNionA) Write(w []byte) []byte {
	w = basictl.IntWrite(w, item.B)
	return w
}

func (item *AuNionA) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xdf61f632); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *AuNionA) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *AuNionA) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xdf61f632)
	return item.Write(w)
}

func (item AuNionA) String() string {
	return string(item.WriteJSON(nil))
}

func (item *AuNionA) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propBPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "b":
				if propBPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("au.nionA", "b")
				}
				if err := internal.Json2ReadInt32(in, &item.B); err != nil {
					return err
				}
				propBPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("au.nionA", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propBPresented {
		item.B = 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *AuNionA) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w), nil
}

func (item *AuNionA) WriteJSON(w []byte) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}
func (item *AuNionA) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) []byte {
	w = append(w, '{')
	backupIndexB := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"b":`...)
	w = basictl.JSONWriteInt32(w, item.B)
	if (item.B != 0) == false {
		w = w[:backupIndexB]
	}
	return append(w, '}')
}

func (item *AuNionA) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *AuNionA) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("au.nionA", err.Error())
	}
	return nil
}

func (item *AuNionA) CalculateLayout(sizes []int) []int {
	sizePosition := len(sizes)
	sizes = append(sizes, 0)

	currentSize := 0
	lastUsedByte := 0

	// add constructor No for union type in case of non first option
	lastUsedByte = 1
	currentSize += basictl.TL2CalculateSize(1)

	// calculate layout for item.B
	if item.B != 0 {

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

func (item *AuNionA) InternalWriteTL2(w []byte, sizes []int) ([]byte, []int) {
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

	// add constructor No for union type in case of non first option
	currentBlock |= (1 << 0)

	w = basictl.TL2WriteSize(w, 1)
	serializedSize += basictl.TL2CalculateSize(1)
	// write item.B
	if item.B != 0 {
		serializedSize += 4
		if 4 != 0 {
			currentBlock |= (1 << 1)
			w = basictl.IntWrite(w, item.B)
		}
	}
	w[currentBlockPosition] = currentBlock
	return w, sizes
}

func (item *AuNionA) WriteTL2(w []byte, ctx *basictl.TL2WriteContext) []byte {
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

func (item *AuNionA) InternalReadTL2(r []byte, block byte) (_ []byte, err error) {
	currentR := r

	// read item.B
	if block&(1<<1) != 0 {
		if currentR, err = basictl.IntRead(currentR, &item.B); err != nil {
			return currentR, err
		}
	} else {
		item.B = 0
	}

	return r, nil
}

func (item *AuNionA) ReadTL2(r []byte, ctx *basictl.TL2ReadContext) (_ []byte, err error) {
	currentSize := 0
	if r, currentSize, err = basictl.TL2ParseSize(r); err != nil {
		return r, err
	}

	currentR := r[:currentSize]
	r = r[currentSize:]

	var block byte
	var index int
	if currentSize == 0 {
		index = 0
	} else {
		if currentR, err = basictl.ByteReadTL2(currentR, &block); err != nil {
			return r, err
		}
		if (block & 1) != 0 {
			if currentR, index, err = basictl.TL2ParseSize(currentR); err != nil {
				return r, err
			}
		} else {
			index = 0
		}
	}
	if index != 1 {
		return r, basictl.TL2Error("unexpected constructor number %d, instead of %d", index, 1)
	}
	_, err = item.InternalReadTL2(currentR, block)
	return r, err
}
