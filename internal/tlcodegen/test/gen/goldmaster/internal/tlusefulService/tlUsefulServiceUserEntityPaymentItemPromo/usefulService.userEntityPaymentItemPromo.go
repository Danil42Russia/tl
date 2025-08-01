// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlUsefulServiceUserEntityPaymentItemPromo

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type UsefulServiceUserEntityPaymentItemPromo struct {
	Content string
}

func (UsefulServiceUserEntityPaymentItemPromo) TLName() string {
	return "usefulService.userEntityPaymentItemPromo"
}
func (UsefulServiceUserEntityPaymentItemPromo) TLTag() uint32 { return 0x24c7ec9f }

func (item *UsefulServiceUserEntityPaymentItemPromo) Reset() {
	item.Content = ""
}

func (item *UsefulServiceUserEntityPaymentItemPromo) FillRandom(rg *basictl.RandGenerator, nat_fields_mask uint32) {
	item.Content = basictl.RandomString(rg)
}

func (item *UsefulServiceUserEntityPaymentItemPromo) Read(w []byte, nat_fields_mask uint32) (_ []byte, err error) {
	return basictl.StringRead(w, &item.Content)
}

func (item *UsefulServiceUserEntityPaymentItemPromo) WriteGeneral(w []byte, nat_fields_mask uint32) (_ []byte, err error) {
	return item.Write(w, nat_fields_mask), nil
}

func (item *UsefulServiceUserEntityPaymentItemPromo) Write(w []byte, nat_fields_mask uint32) []byte {
	w = basictl.StringWrite(w, item.Content)
	return w
}

func (item *UsefulServiceUserEntityPaymentItemPromo) ReadBoxed(w []byte, nat_fields_mask uint32) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x24c7ec9f); err != nil {
		return w, err
	}
	return item.Read(w, nat_fields_mask)
}

func (item *UsefulServiceUserEntityPaymentItemPromo) WriteBoxedGeneral(w []byte, nat_fields_mask uint32) (_ []byte, err error) {
	return item.WriteBoxed(w, nat_fields_mask), nil
}

func (item *UsefulServiceUserEntityPaymentItemPromo) WriteBoxed(w []byte, nat_fields_mask uint32) []byte {
	w = basictl.NatWrite(w, 0x24c7ec9f)
	return item.Write(w, nat_fields_mask)
}

func (item *UsefulServiceUserEntityPaymentItemPromo) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, nat_fields_mask uint32) error {
	var propContentPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "content":
				if propContentPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("usefulService.userEntityPaymentItemPromo", "content")
				}
				if err := internal.Json2ReadString(in, &item.Content); err != nil {
					return err
				}
				propContentPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("usefulService.userEntityPaymentItemPromo", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propContentPresented {
		item.Content = ""
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *UsefulServiceUserEntityPaymentItemPromo) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte, nat_fields_mask uint32) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w, nat_fields_mask), nil
}

func (item *UsefulServiceUserEntityPaymentItemPromo) WriteJSON(w []byte, nat_fields_mask uint32) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w, nat_fields_mask)
}
func (item *UsefulServiceUserEntityPaymentItemPromo) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte, nat_fields_mask uint32) []byte {
	w = append(w, '{')
	backupIndexContent := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"content":`...)
	w = basictl.JSONWriteString(w, item.Content)
	if (len(item.Content) != 0) == false {
		w = w[:backupIndexContent]
	}
	return append(w, '}')
}

func (item *UsefulServiceUserEntityPaymentItemPromo) CalculateLayout(sizes []int, nat_fields_mask uint32) []int {
	sizePosition := len(sizes)
	sizes = append(sizes, 0)

	currentSize := 0
	lastUsedByte := 0

	// calculate layout for item.Content
	if len(item.Content) != 0 {

		if len(item.Content) != 0 {
			lastUsedByte = 1
			currentSize += len(item.Content)
			currentSize += basictl.TL2CalculateSize(len(item.Content))
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

func (item *UsefulServiceUserEntityPaymentItemPromo) InternalWriteTL2(w []byte, sizes []int, nat_fields_mask uint32) ([]byte, []int) {
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
	// write item.Content
	if len(item.Content) != 0 {
		serializedSize += len(item.Content)
		if len(item.Content) != 0 {
			serializedSize += basictl.TL2CalculateSize(len(item.Content))
			currentBlock |= (1 << 1)
			w = basictl.StringWriteTL2(w, item.Content)
		}
	}
	w[currentBlockPosition] = currentBlock
	return w, sizes
}

func (item *UsefulServiceUserEntityPaymentItemPromo) WriteTL2(w []byte, ctx *basictl.TL2WriteContext, nat_fields_mask uint32) []byte {
	var sizes []int
	if ctx != nil {
		sizes = ctx.SizeBuffer
	}
	sizes = item.CalculateLayout(sizes[:0], nat_fields_mask)
	w, _ = item.InternalWriteTL2(w, sizes, nat_fields_mask)
	if ctx != nil {
		ctx.SizeBuffer = sizes[:0]
	}
	return w
}

func (item *UsefulServiceUserEntityPaymentItemPromo) InternalReadTL2(r []byte, nat_fields_mask uint32) (_ []byte, err error) {
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

	// read item.Content
	if block&(1<<1) != 0 {
		if currentR, err = basictl.StringReadTL2(currentR, &item.Content); err != nil {
			return currentR, err
		}
	} else {
		item.Content = ""
	}

	return r, nil
}

func (item *UsefulServiceUserEntityPaymentItemPromo) ReadTL2(r []byte, ctx *basictl.TL2ReadContext, nat_fields_mask uint32) (_ []byte, err error) {
	return item.InternalReadTL2(r, nat_fields_mask)
}
