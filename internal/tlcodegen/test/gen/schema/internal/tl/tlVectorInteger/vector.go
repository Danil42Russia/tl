// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlVectorInteger

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlBuiltinVectorInteger"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlInteger"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type VectorInteger []tlInteger.Integer

func (VectorInteger) TLName() string { return "vector" }
func (VectorInteger) TLTag() uint32  { return 0x1cb5c415 }

func (item *VectorInteger) Reset() {
	ptr := (*[]tlInteger.Integer)(item)
	*ptr = (*ptr)[:0]
}

func (item *VectorInteger) Read(w []byte) (_ []byte, err error) {
	ptr := (*[]tlInteger.Integer)(item)
	return tlBuiltinVectorInteger.BuiltinVectorIntegerRead(w, ptr)
}

func (item *VectorInteger) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *VectorInteger) Write(w []byte) []byte {
	ptr := (*[]tlInteger.Integer)(item)
	return tlBuiltinVectorInteger.BuiltinVectorIntegerWrite(w, *ptr)
}

func (item *VectorInteger) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x1cb5c415); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *VectorInteger) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *VectorInteger) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x1cb5c415)
	return item.Write(w)
}

func (item VectorInteger) String() string {
	return string(item.WriteJSON(nil))
}
func (item *VectorInteger) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	ptr := (*[]tlInteger.Integer)(item)
	if err := tlBuiltinVectorInteger.BuiltinVectorIntegerReadJSON(legacyTypeNames, in, ptr); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *VectorInteger) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w), nil
}

func (item *VectorInteger) WriteJSON(w []byte) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}

func (item *VectorInteger) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) []byte {
	ptr := (*[]tlInteger.Integer)(item)
	w = tlBuiltinVectorInteger.BuiltinVectorIntegerWriteJSONOpt(tctx, w, *ptr)
	return w
}
func (item *VectorInteger) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *VectorInteger) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("vector", err.Error())
	}
	return nil
}
