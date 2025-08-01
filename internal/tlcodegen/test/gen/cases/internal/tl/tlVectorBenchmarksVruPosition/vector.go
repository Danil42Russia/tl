// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlVectorBenchmarksVruPosition

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tl/tlBuiltinVectorBenchmarksVruPosition"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/cases/internal/tlbenchmarks/tlBenchmarksVruPosition"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type VectorBenchmarksVruPosition []tlBenchmarksVruPosition.BenchmarksVruPosition

func (VectorBenchmarksVruPosition) TLName() string { return "vector" }
func (VectorBenchmarksVruPosition) TLTag() uint32  { return 0x1cb5c415 }

func (item *VectorBenchmarksVruPosition) Reset() {
	ptr := (*[]tlBenchmarksVruPosition.BenchmarksVruPosition)(item)
	*ptr = (*ptr)[:0]
}

func (item *VectorBenchmarksVruPosition) FillRandom(rg *basictl.RandGenerator) {
	ptr := (*[]tlBenchmarksVruPosition.BenchmarksVruPosition)(item)
	tlBuiltinVectorBenchmarksVruPosition.BuiltinVectorBenchmarksVruPositionFillRandom(rg, ptr)
}

func (item *VectorBenchmarksVruPosition) Read(w []byte) (_ []byte, err error) {
	ptr := (*[]tlBenchmarksVruPosition.BenchmarksVruPosition)(item)
	return tlBuiltinVectorBenchmarksVruPosition.BuiltinVectorBenchmarksVruPositionRead(w, ptr)
}

func (item *VectorBenchmarksVruPosition) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *VectorBenchmarksVruPosition) Write(w []byte) []byte {
	ptr := (*[]tlBenchmarksVruPosition.BenchmarksVruPosition)(item)
	return tlBuiltinVectorBenchmarksVruPosition.BuiltinVectorBenchmarksVruPositionWrite(w, *ptr)
}

func (item *VectorBenchmarksVruPosition) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x1cb5c415); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *VectorBenchmarksVruPosition) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *VectorBenchmarksVruPosition) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x1cb5c415)
	return item.Write(w)
}

func (item VectorBenchmarksVruPosition) String() string {
	return string(item.WriteJSON(nil))
}
func (item *VectorBenchmarksVruPosition) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	ptr := (*[]tlBenchmarksVruPosition.BenchmarksVruPosition)(item)
	if err := tlBuiltinVectorBenchmarksVruPosition.BuiltinVectorBenchmarksVruPositionReadJSON(legacyTypeNames, in, ptr); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *VectorBenchmarksVruPosition) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w), nil
}

func (item *VectorBenchmarksVruPosition) WriteJSON(w []byte) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}

func (item *VectorBenchmarksVruPosition) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) []byte {
	ptr := (*[]tlBenchmarksVruPosition.BenchmarksVruPosition)(item)
	w = tlBuiltinVectorBenchmarksVruPosition.BuiltinVectorBenchmarksVruPositionWriteJSONOpt(tctx, w, *ptr)
	return w
}
func (item *VectorBenchmarksVruPosition) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *VectorBenchmarksVruPosition) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("vector", err.Error())
	}
	return nil
}

func (item *VectorBenchmarksVruPosition) CalculateLayout(sizes []int) []int {
	ptr := (*[]tlBenchmarksVruPosition.BenchmarksVruPosition)(item)
	sizes = tlBuiltinVectorBenchmarksVruPosition.BuiltinVectorBenchmarksVruPositionCalculateLayout(sizes, ptr)
	return sizes
}

func (item *VectorBenchmarksVruPosition) InternalWriteTL2(w []byte, sizes []int) ([]byte, []int) {
	ptr := (*[]tlBenchmarksVruPosition.BenchmarksVruPosition)(item)
	w, sizes = tlBuiltinVectorBenchmarksVruPosition.BuiltinVectorBenchmarksVruPositionInternalWriteTL2(w, sizes, ptr)
	return w, sizes
}

func (item *VectorBenchmarksVruPosition) WriteTL2(w []byte, ctx *basictl.TL2WriteContext) []byte {
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

func (item *VectorBenchmarksVruPosition) InternalReadTL2(r []byte) (_ []byte, err error) {
	ptr := (*[]tlBenchmarksVruPosition.BenchmarksVruPosition)(item)
	if r, err = tlBuiltinVectorBenchmarksVruPosition.BuiltinVectorBenchmarksVruPositionInternalReadTL2(r, ptr); err != nil {
		return r, err
	}
	return r, nil
}

func (item *VectorBenchmarksVruPosition) ReadTL2(r []byte, ctx *basictl.TL2ReadContext) (_ []byte, err error) {
	return item.InternalReadTL2(r)
}
