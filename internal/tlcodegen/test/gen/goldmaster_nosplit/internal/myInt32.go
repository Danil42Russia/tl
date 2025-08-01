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

func BuiltinTuple3MyInt32Reset(vec *[3]MyInt32) {
	for i := range *vec {
		(*vec)[i].Reset()
	}
}

func BuiltinTuple3MyInt32FillRandom(rg *basictl.RandGenerator, vec *[3]MyInt32) {
	rg.IncreaseDepth()
	for i := range *vec {
		(*vec)[i].FillRandom(rg)
	}
	rg.DecreaseDepth()
}

func BuiltinTuple3MyInt32Read(w []byte, vec *[3]MyInt32) (_ []byte, err error) {
	for i := range *vec {
		if w, err = (*vec)[i].Read(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func BuiltinTuple3MyInt32Write(w []byte, vec *[3]MyInt32) []byte {
	for _, elem := range *vec {
		w = elem.Write(w)
	}
	return w
}

func BuiltinTuple3MyInt32CalculateLayout(sizes []int, vec *[3]MyInt32) []int {
	currentSize := 0
	sizePosition := len(sizes)
	sizes = append(sizes, 0)
	if 3 != 0 {
		currentSize += basictl.TL2CalculateSize(3)
	}

	for i := 0; i < 3; i++ {
		sizes = (*vec)[i].CalculateLayout(sizes)
		currentSize += 4
	}

	sizes[sizePosition] = currentSize
	return sizes
}

func BuiltinTuple3MyInt32InternalWriteTL2(w []byte, sizes []int, vec *[3]MyInt32) ([]byte, []int) {
	currentSize := sizes[0]
	sizes = sizes[1:]

	w = basictl.TL2WriteSize(w, currentSize)
	if 3 != 0 {
		w = basictl.TL2WriteSize(w, 3)
	}

	for i := 0; i < 3; i++ {
		w, sizes = (*vec)[i].InternalWriteTL2(w, sizes)
	}
	return w, sizes
}

func BuiltinTuple3MyInt32InternalReadTL2(r []byte, vec *[3]MyInt32) (_ []byte, err error) {
	currentSize := 0
	if r, currentSize, err = basictl.TL2ParseSize(r); err != nil {
		return r, err
	}
	if len(r) < currentSize {
		return r, basictl.TL2Error("not enough data: expected %d, got %d", currentSize, len(r))
	}

	currentR := r[:currentSize]
	r = r[currentSize:]

	elementCount := 0
	if currentSize != 0 {
		if currentR, elementCount, err = basictl.TL2ParseSize(currentR); err != nil {
			return r, err
		}
	}

	lastIndex := elementCount
	if lastIndex > 3 {
		lastIndex = 3
	}
	for i := 0; i < lastIndex; i++ {
		if currentR, err = (*vec)[i].InternalReadTL2(currentR); err != nil {
			return currentR, err
		}
	}

	// reset elements if received less elements
	for i := lastIndex; i < 3; i++ {
		(*vec)[i].Reset()
	}

	return r, nil
}

func BuiltinTuple3MyInt32ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, vec *[3]MyInt32) error {
	index := 0
	if in != nil {
		in.Delim('[')
		if !in.Ok() {
			return ErrorInvalidJSON("[3]MyInt32", "expected json array")
		}
		for ; !in.IsDelim(']'); index++ {
			if index == 3 {
				return ErrorWrongSequenceLength("[3]MyInt32", index+1, 3)
			}
			if err := (*vec)[index].ReadJSON(legacyTypeNames, in); err != nil {
				return err
			}
			in.WantComma()
		}
		in.Delim(']')
		if !in.Ok() {
			return ErrorInvalidJSON("[3]MyInt32", "expected json array's end")
		}
	}
	if index != 3 {
		return ErrorWrongSequenceLength("[3]MyInt32", index+1, 3)
	}
	return nil
}

func BuiltinTuple3MyInt32WriteJSON(w []byte, vec *[3]MyInt32) []byte {
	tctx := basictl.JSONWriteContext{}
	return BuiltinTuple3MyInt32WriteJSONOpt(&tctx, w, vec)
}
func BuiltinTuple3MyInt32WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte, vec *[3]MyInt32) []byte {
	w = append(w, '[')
	for _, elem := range *vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = elem.WriteJSONOpt(tctx, w)
	}
	return append(w, ']')
}

func BuiltinTuple3MyInt32BoxedReset(vec *[3]MyInt32) {
	for i := range *vec {
		(*vec)[i].Reset()
	}
}

func BuiltinTuple3MyInt32BoxedFillRandom(rg *basictl.RandGenerator, vec *[3]MyInt32) {
	rg.IncreaseDepth()
	for i := range *vec {
		(*vec)[i].FillRandom(rg)
	}
	rg.DecreaseDepth()
}

func BuiltinTuple3MyInt32BoxedRead(w []byte, vec *[3]MyInt32) (_ []byte, err error) {
	for i := range *vec {
		if w, err = (*vec)[i].ReadBoxed(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func BuiltinTuple3MyInt32BoxedWrite(w []byte, vec *[3]MyInt32) []byte {
	for _, elem := range *vec {
		w = elem.WriteBoxed(w)
	}
	return w
}

func BuiltinTuple3MyInt32BoxedCalculateLayout(sizes []int, vec *[3]MyInt32) []int {
	currentSize := 0
	sizePosition := len(sizes)
	sizes = append(sizes, 0)
	if 3 != 0 {
		currentSize += basictl.TL2CalculateSize(3)
	}

	for i := 0; i < 3; i++ {
		sizes = (*vec)[i].CalculateLayout(sizes)
		currentSize += 4
	}

	sizes[sizePosition] = currentSize
	return sizes
}

func BuiltinTuple3MyInt32BoxedInternalWriteTL2(w []byte, sizes []int, vec *[3]MyInt32) ([]byte, []int) {
	currentSize := sizes[0]
	sizes = sizes[1:]

	w = basictl.TL2WriteSize(w, currentSize)
	if 3 != 0 {
		w = basictl.TL2WriteSize(w, 3)
	}

	for i := 0; i < 3; i++ {
		w, sizes = (*vec)[i].InternalWriteTL2(w, sizes)
	}
	return w, sizes
}

func BuiltinTuple3MyInt32BoxedInternalReadTL2(r []byte, vec *[3]MyInt32) (_ []byte, err error) {
	currentSize := 0
	if r, currentSize, err = basictl.TL2ParseSize(r); err != nil {
		return r, err
	}
	if len(r) < currentSize {
		return r, basictl.TL2Error("not enough data: expected %d, got %d", currentSize, len(r))
	}

	currentR := r[:currentSize]
	r = r[currentSize:]

	elementCount := 0
	if currentSize != 0 {
		if currentR, elementCount, err = basictl.TL2ParseSize(currentR); err != nil {
			return r, err
		}
	}

	lastIndex := elementCount
	if lastIndex > 3 {
		lastIndex = 3
	}
	for i := 0; i < lastIndex; i++ {
		if currentR, err = (*vec)[i].InternalReadTL2(currentR); err != nil {
			return currentR, err
		}
	}

	// reset elements if received less elements
	for i := lastIndex; i < 3; i++ {
		(*vec)[i].Reset()
	}

	return r, nil
}

func BuiltinTuple3MyInt32BoxedReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, vec *[3]MyInt32) error {
	index := 0
	if in != nil {
		in.Delim('[')
		if !in.Ok() {
			return ErrorInvalidJSON("[3]MyInt32", "expected json array")
		}
		for ; !in.IsDelim(']'); index++ {
			if index == 3 {
				return ErrorWrongSequenceLength("[3]MyInt32", index+1, 3)
			}
			if err := (*vec)[index].ReadJSON(legacyTypeNames, in); err != nil {
				return err
			}
			in.WantComma()
		}
		in.Delim(']')
		if !in.Ok() {
			return ErrorInvalidJSON("[3]MyInt32", "expected json array's end")
		}
	}
	if index != 3 {
		return ErrorWrongSequenceLength("[3]MyInt32", index+1, 3)
	}
	return nil
}

func BuiltinTuple3MyInt32BoxedWriteJSON(w []byte, vec *[3]MyInt32) []byte {
	tctx := basictl.JSONWriteContext{}
	return BuiltinTuple3MyInt32BoxedWriteJSONOpt(&tctx, w, vec)
}
func BuiltinTuple3MyInt32BoxedWriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte, vec *[3]MyInt32) []byte {
	w = append(w, '[')
	for _, elem := range *vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = elem.WriteJSONOpt(tctx, w)
	}
	return append(w, ']')
}

type MyInt32 Int32

func (MyInt32) TLName() string { return "myInt32" }
func (MyInt32) TLTag() uint32  { return 0xba59e151 }

func (item *MyInt32) Reset() {
	ptr := (*Int32)(item)
	ptr.Reset()
}

func (item *MyInt32) FillRandom(rg *basictl.RandGenerator) {
	ptr := (*Int32)(item)
	ptr.FillRandom(rg)
}

func (item *MyInt32) Read(w []byte) (_ []byte, err error) {
	ptr := (*Int32)(item)
	return ptr.Read(w)
}

func (item *MyInt32) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *MyInt32) Write(w []byte) []byte {
	ptr := (*Int32)(item)
	return ptr.Write(w)
}

func (item *MyInt32) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xba59e151); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *MyInt32) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *MyInt32) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xba59e151)
	return item.Write(w)
}

func (item MyInt32) String() string {
	return string(item.WriteJSON(nil))
}
func (item *MyInt32) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	ptr := (*Int32)(item)
	if err := ptr.ReadJSON(legacyTypeNames, in); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *MyInt32) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w), nil
}

func (item *MyInt32) WriteJSON(w []byte) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}

func (item *MyInt32) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) []byte {
	ptr := (*Int32)(item)
	w = ptr.WriteJSONOpt(tctx, w)
	return w
}
func (item *MyInt32) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *MyInt32) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("myInt32", err.Error())
	}
	return nil
}

func (item *MyInt32) CalculateLayout(sizes []int) []int {
	ptr := (*Int32)(item)
	sizes = (*ptr).CalculateLayout(sizes)
	return sizes
}

func (item *MyInt32) InternalWriteTL2(w []byte, sizes []int) ([]byte, []int) {
	ptr := (*Int32)(item)
	w, sizes = ptr.InternalWriteTL2(w, sizes)
	return w, sizes
}

func (item *MyInt32) WriteTL2(w []byte, ctx *basictl.TL2WriteContext) []byte {
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

func (item *MyInt32) InternalReadTL2(r []byte) (_ []byte, err error) {
	ptr := (*Int32)(item)
	if r, err = ptr.InternalReadTL2(r); err != nil {
		return r, err
	}
	return r, nil
}

func (item *MyInt32) ReadTL2(r []byte, ctx *basictl.TL2ReadContext) (_ []byte, err error) {
	return item.InternalReadTL2(r)
}
