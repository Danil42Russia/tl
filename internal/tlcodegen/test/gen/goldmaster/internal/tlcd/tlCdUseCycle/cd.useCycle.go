// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlCdUseCycle

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/cycle_e10cb78db8a2766007111b86ce9e11d9"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlAColorBoxedMaybe"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type CdUseCycle struct {
	A cycle_e10cb78db8a2766007111b86ce9e11d9.Cyc2MyCycle
	B tlAColorBoxedMaybe.AColorBoxedMaybe
}

func (CdUseCycle) TLName() string { return "cd.useCycle" }
func (CdUseCycle) TLTag() uint32  { return 0x6ed67ca0 }

func (item *CdUseCycle) Reset() {
	item.A.Reset()
	item.B.Reset()
}

func (item *CdUseCycle) FillRandom(rg *basictl.RandGenerator) {
	item.A.FillRandom(rg)
	item.B.FillRandom(rg)
}

func (item *CdUseCycle) Read(w []byte) (_ []byte, err error) {
	if w, err = item.A.Read(w); err != nil {
		return w, err
	}
	return item.B.ReadBoxed(w)
}

func (item *CdUseCycle) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *CdUseCycle) Write(w []byte) []byte {
	w = item.A.Write(w)
	w = item.B.WriteBoxed(w)
	return w
}

func (item *CdUseCycle) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x6ed67ca0); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *CdUseCycle) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *CdUseCycle) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x6ed67ca0)
	return item.Write(w)
}

func (item CdUseCycle) String() string {
	return string(item.WriteJSON(nil))
}

func (item *CdUseCycle) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propAPresented bool
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
			case "a":
				if propAPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("cd.useCycle", "a")
				}
				if err := item.A.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propAPresented = true
			case "b":
				if propBPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("cd.useCycle", "b")
				}
				if err := item.B.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propBPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("cd.useCycle", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propAPresented {
		item.A.Reset()
	}
	if !propBPresented {
		item.B.Reset()
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *CdUseCycle) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w), nil
}

func (item *CdUseCycle) WriteJSON(w []byte) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}
func (item *CdUseCycle) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) []byte {
	w = append(w, '{')
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"a":`...)
	w = item.A.WriteJSONOpt(tctx, w)
	backupIndexB := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"b":`...)
	w = item.B.WriteJSONOpt(tctx, w)
	if (item.B.Ok) == false {
		w = w[:backupIndexB]
	}
	return append(w, '}')
}

func (item *CdUseCycle) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *CdUseCycle) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("cd.useCycle", err.Error())
	}
	return nil
}

func (item *CdUseCycle) CalculateLayout(sizes []int) []int {
	sizePosition := len(sizes)
	sizes = append(sizes, 0)

	currentSize := 0
	lastUsedByte := 0

	// calculate layout for item.A
	currentPosition := len(sizes)
	sizes = item.A.CalculateLayout(sizes)
	if sizes[currentPosition] != 0 {
		lastUsedByte = 1
		currentSize += sizes[currentPosition]
		currentSize += basictl.TL2CalculateSize(sizes[currentPosition])
	} else {
		sizes = sizes[:currentPosition+1]
	}

	// calculate layout for item.B
	currentPosition = len(sizes)
	if item.B.Ok {
		sizes = item.B.CalculateLayout(sizes)
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

func (item *CdUseCycle) InternalWriteTL2(w []byte, sizes []int) ([]byte, []int) {
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
	serializedSize += sizes[0]
	if sizes[0] != 0 {
		serializedSize += basictl.TL2CalculateSize(sizes[0])
		currentBlock |= (1 << 1)
		w, sizes = item.A.InternalWriteTL2(w, sizes)
	} else {
		sizes = sizes[1:]
	}
	// write item.B
	if item.B.Ok {
		serializedSize += sizes[0]
		if sizes[0] != 0 {
			serializedSize += basictl.TL2CalculateSize(sizes[0])
			currentBlock |= (1 << 2)
			w, sizes = item.B.InternalWriteTL2(w, sizes)
		} else {
			sizes = sizes[1:]
		}
	}
	w[currentBlockPosition] = currentBlock
	return w, sizes
}

func (item *CdUseCycle) WriteTL2(w []byte, ctx *basictl.TL2WriteContext) []byte {
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

func (item *CdUseCycle) InternalReadTL2(r []byte) (_ []byte, err error) {
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

	// read item.A
	if block&(1<<1) != 0 {
		if currentR, err = item.A.InternalReadTL2(currentR); err != nil {
			return currentR, err
		}
	} else {
		item.A.Reset()
	}

	// read item.B
	if block&(1<<2) != 0 {
		if currentR, err = item.B.InternalReadTL2(currentR); err != nil {
			return currentR, err
		}
	} else {
		item.B.Reset()
	}

	return r, nil
}

func (item *CdUseCycle) ReadTL2(r []byte, ctx *basictl.TL2ReadContext) (_ []byte, err error) {
	return item.InternalReadTL2(r)
}
