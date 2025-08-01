// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlBuiltinTuplePairBoxedIntLong

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/goldmaster/internal/tl/tlPairIntLong"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

func BuiltinTuplePairBoxedIntLongFillRandom(rg *basictl.RandGenerator, vec *[]tlPairIntLong.PairIntLong, nat_n uint32) {
	rg.IncreaseDepth()
	*vec = make([]tlPairIntLong.PairIntLong, nat_n)
	for i := range *vec {
		(*vec)[i].FillRandom(rg)
	}
	rg.DecreaseDepth()
}

func BuiltinTuplePairBoxedIntLongRead(w []byte, vec *[]tlPairIntLong.PairIntLong, nat_n uint32) (_ []byte, err error) {
	if uint32(cap(*vec)) < nat_n {
		*vec = make([]tlPairIntLong.PairIntLong, nat_n)
	} else {
		*vec = (*vec)[:nat_n]
	}
	for i := range *vec {
		if w, err = (*vec)[i].ReadBoxed(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func BuiltinTuplePairBoxedIntLongWrite(w []byte, vec []tlPairIntLong.PairIntLong, nat_n uint32) (_ []byte, err error) {
	if uint32(len(vec)) != nat_n {
		return w, internal.ErrorWrongSequenceLength("[]tlPairIntLong.PairIntLong", len(vec), nat_n)
	}
	for _, elem := range vec {
		w = elem.WriteBoxed(w)
	}
	return w, nil
}

func BuiltinTuplePairBoxedIntLongCalculateLayout(sizes []int, vec *[]tlPairIntLong.PairIntLong, nat_n uint32) []int {
	currentSize := 0
	sizePosition := len(sizes)
	sizes = append(sizes, 0)
	if nat_n != 0 {
		currentSize += basictl.TL2CalculateSize(int(nat_n))
	}

	lastIndex := uint32(len(*vec))
	if lastIndex > nat_n {
		lastIndex = nat_n
	}

	for i := uint32(0); i < lastIndex; i++ {
		currentPosition := len(sizes)
		sizes = (*vec)[i].CalculateLayout(sizes)
		currentSize += sizes[currentPosition]
		currentSize += basictl.TL2CalculateSize(sizes[currentPosition])
	}

	// append empty objects if not enough
	for i := lastIndex; i < nat_n; i++ {
		var elem tlPairIntLong.PairIntLong
		currentPosition := len(sizes)
		sizes = elem.CalculateLayout(sizes)
		currentSize += sizes[currentPosition]
		currentSize += basictl.TL2CalculateSize(sizes[currentPosition])
	}

	sizes[sizePosition] = currentSize
	return sizes
}

func BuiltinTuplePairBoxedIntLongInternalWriteTL2(w []byte, sizes []int, vec *[]tlPairIntLong.PairIntLong, nat_n uint32) ([]byte, []int) {
	currentSize := sizes[0]
	sizes = sizes[1:]

	w = basictl.TL2WriteSize(w, currentSize)
	if nat_n != 0 {
		w = basictl.TL2WriteSize(w, int(nat_n))
	}

	lastIndex := uint32(len(*vec))
	if lastIndex > nat_n {
		lastIndex = nat_n
	}

	for i := uint32(0); i < lastIndex; i++ {
		w, sizes = (*vec)[i].InternalWriteTL2(w, sizes)
	}

	// append empty objects if not enough
	for i := lastIndex; i < nat_n; i++ {
		var elem tlPairIntLong.PairIntLong
		w, sizes = elem.InternalWriteTL2(w, sizes)
	}
	return w, sizes
}

func BuiltinTuplePairBoxedIntLongInternalReadTL2(r []byte, vec *[]tlPairIntLong.PairIntLong, nat_n uint32) (_ []byte, err error) {
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

	if uint32(cap(*vec)) < nat_n {
		*vec = make([]tlPairIntLong.PairIntLong, nat_n)
	} else {
		*vec = (*vec)[:nat_n]
	}

	lastIndex := uint32(elementCount)
	if lastIndex > nat_n {
		lastIndex = nat_n
	}

	for i := uint32(0); i < lastIndex; i++ {
		if currentR, err = (*vec)[i].InternalReadTL2(currentR); err != nil {
			return currentR, err
		}
	}

	// reset elements if received less elements
	for i := lastIndex; i < nat_n; i++ {
		(*vec)[i].Reset()
	}

	return r, nil
}
func BuiltinTuplePairBoxedIntLongReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, vec *[]tlPairIntLong.PairIntLong, nat_n uint32) error {
	if uint32(cap(*vec)) < nat_n {
		*vec = make([]tlPairIntLong.PairIntLong, nat_n)
	} else {
		*vec = (*vec)[:nat_n]
	}
	index := 0
	if in != nil {
		in.Delim('[')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("[]tlPairIntLong.PairIntLong", "expected json array")
		}
		for ; !in.IsDelim(']'); index++ {
			if nat_n <= uint32(index) {
				return internal.ErrorInvalidJSON("[]tlPairIntLong.PairIntLong", "array is longer than expected")
			}
			if err := (*vec)[index].ReadJSON(legacyTypeNames, in); err != nil {
				return err
			}
			in.WantComma()
		}
		in.Delim(']')
		if !in.Ok() {
			return internal.ErrorInvalidJSON("[]tlPairIntLong.PairIntLong", "expected json array's end")
		}
	}
	if uint32(index) != nat_n {
		return internal.ErrorWrongSequenceLength("[]tlPairIntLong.PairIntLong", index, nat_n)
	}
	return nil
}

func BuiltinTuplePairBoxedIntLongWriteJSON(w []byte, vec []tlPairIntLong.PairIntLong, nat_n uint32) (_ []byte, err error) {
	tctx := basictl.JSONWriteContext{}
	return BuiltinTuplePairBoxedIntLongWriteJSONOpt(&tctx, w, vec, nat_n)
}
func BuiltinTuplePairBoxedIntLongWriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte, vec []tlPairIntLong.PairIntLong, nat_n uint32) (_ []byte, err error) {
	if uint32(len(vec)) != nat_n {
		return w, internal.ErrorWrongSequenceLength("[]tlPairIntLong.PairIntLong", len(vec), nat_n)
	}
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = elem.WriteJSONOpt(tctx, w)
	}
	return append(w, ']'), nil
}
