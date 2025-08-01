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

type AbTypeC struct {
	X CdTypeA
}

func (AbTypeC) TLName() string { return "ab.typeC" }
func (AbTypeC) TLTag() uint32  { return 0x69920d6e }

func (item *AbTypeC) Reset() {
	item.X.Reset()
}

func (item *AbTypeC) FillRandom(rg *basictl.RandGenerator) {
	item.X.FillRandom(rg)
}

func (item *AbTypeC) Read(w []byte) (_ []byte, err error) {
	return item.X.Read(w)
}

func (item *AbTypeC) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *AbTypeC) Write(w []byte) []byte {
	w = item.X.Write(w)
	return w
}

func (item *AbTypeC) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x69920d6e); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *AbTypeC) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *AbTypeC) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x69920d6e)
	return item.Write(w)
}

func (item AbTypeC) String() string {
	return string(item.WriteJSON(nil))
}

func (item *AbTypeC) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
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
					return ErrorInvalidJSONWithDuplicatingKeys("ab.typeC", "x")
				}
				if err := item.X.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propXPresented = true
			default:
				return ErrorInvalidJSONExcessElement("ab.typeC", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propXPresented {
		item.X.Reset()
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *AbTypeC) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w), nil
}

func (item *AbTypeC) WriteJSON(w []byte) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}
func (item *AbTypeC) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) []byte {
	w = append(w, '{')
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"x":`...)
	w = item.X.WriteJSONOpt(tctx, w)
	return append(w, '}')
}

func (item *AbTypeC) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *AbTypeC) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("ab.typeC", err.Error())
	}
	return nil
}

func (item *AbTypeC) CalculateLayout(sizes []int) []int {
	sizePosition := len(sizes)
	sizes = append(sizes, 0)

	currentSize := 0
	lastUsedByte := 0

	// calculate layout for item.X
	currentPosition := len(sizes)
	sizes = item.X.CalculateLayout(sizes)
	if sizes[currentPosition] != 0 {
		lastUsedByte = 1
		currentSize += sizes[currentPosition]
		currentSize += basictl.TL2CalculateSize(sizes[currentPosition])
	} else {
		sizes = sizes[:currentPosition+1]
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

func (item *AbTypeC) InternalWriteTL2(w []byte, sizes []int) ([]byte, []int) {
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
	serializedSize += sizes[0]
	if sizes[0] != 0 {
		serializedSize += basictl.TL2CalculateSize(sizes[0])
		currentBlock |= (1 << 1)
		w, sizes = item.X.InternalWriteTL2(w, sizes)
	} else {
		sizes = sizes[1:]
	}
	w[currentBlockPosition] = currentBlock
	return w, sizes
}

func (item *AbTypeC) WriteTL2(w []byte, ctx *basictl.TL2WriteContext) []byte {
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

func (item *AbTypeC) InternalReadTL2(r []byte) (_ []byte, err error) {
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
		if currentR, err = item.X.InternalReadTL2(currentR); err != nil {
			return currentR, err
		}
	} else {
		item.X.Reset()
	}

	return r, nil
}

func (item *AbTypeC) ReadTL2(r []byte, ctx *basictl.TL2ReadContext) (_ []byte, err error) {
	return item.InternalReadTL2(r)
}
