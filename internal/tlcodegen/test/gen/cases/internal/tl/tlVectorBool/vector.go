// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlVectorBool

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tl/tlBuiltinVectorBool"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type VectorBool []bool

func (VectorBool) TLName() string { return "vector" }
func (VectorBool) TLTag() uint32  { return 0x1cb5c415 }

func (item *VectorBool) Reset() {
	ptr := (*[]bool)(item)
	*ptr = (*ptr)[:0]
}

func (item *VectorBool) FillRandom(rg *basictl.RandGenerator) {
	ptr := (*[]bool)(item)
	tlBuiltinVectorBool.BuiltinVectorBoolFillRandom(rg, ptr)
}

func (item *VectorBool) Read(w []byte) (_ []byte, err error) {
	ptr := (*[]bool)(item)
	return tlBuiltinVectorBool.BuiltinVectorBoolRead(w, ptr)
}

func (item *VectorBool) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *VectorBool) Write(w []byte) []byte {
	ptr := (*[]bool)(item)
	return tlBuiltinVectorBool.BuiltinVectorBoolWrite(w, *ptr)
}

func (item *VectorBool) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x1cb5c415); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *VectorBool) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *VectorBool) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x1cb5c415)
	return item.Write(w)
}

func (item VectorBool) String() string {
	return string(item.WriteJSON(nil))
}
func (item *VectorBool) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	ptr := (*[]bool)(item)
	if err := tlBuiltinVectorBool.BuiltinVectorBoolReadJSON(legacyTypeNames, in, ptr); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *VectorBool) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w), nil
}

func (item *VectorBool) WriteJSON(w []byte) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}

func (item *VectorBool) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) []byte {
	ptr := (*[]bool)(item)
	w = tlBuiltinVectorBool.BuiltinVectorBoolWriteJSONOpt(tctx, w, *ptr)
	return w
}
func (item *VectorBool) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *VectorBool) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("vector", err.Error())
	}
	return nil
}

func (item *VectorBool) CalculateLayout(sizes []int) []int {
	ptr := (*[]bool)(item)
	sizes = tlBuiltinVectorBool.BuiltinVectorBoolCalculateLayout(sizes, ptr)
	return sizes
}

func (item *VectorBool) InternalWriteTL2(w []byte, sizes []int) ([]byte, []int) {
	ptr := (*[]bool)(item)
	w, sizes = tlBuiltinVectorBool.BuiltinVectorBoolInternalWriteTL2(w, sizes, ptr)
	return w, sizes
}

func (item *VectorBool) WriteTL2(w []byte, ctx *basictl.TL2WriteContext) []byte {
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

func (item *VectorBool) InternalReadTL2(r []byte) (_ []byte, err error) {
	ptr := (*[]bool)(item)
	if r, err = tlBuiltinVectorBool.BuiltinVectorBoolInternalReadTL2(r, ptr); err != nil {
		return r, err
	}
	return r, nil
}

func (item *VectorBool) ReadTL2(r []byte, ctx *basictl.TL2ReadContext) (_ []byte, err error) {
	return item.InternalReadTL2(r)
}
