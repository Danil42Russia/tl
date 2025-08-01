// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package cycle_16847572a0831d4cd4c0c0fb513151f3

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

func (item Service5EmptyOutput) AsUnion() Service5Output {
	var ret Service5Output
	ret.SetEmpty()
	return ret
}

type Service5EmptyOutput struct {
}

func (Service5EmptyOutput) TLName() string { return "service5.emptyOutput" }
func (Service5EmptyOutput) TLTag() uint32  { return 0x11e46879 }

func (item *Service5EmptyOutput) Reset() {}

func (item *Service5EmptyOutput) Read(w []byte) (_ []byte, err error) { return w, nil }

func (item *Service5EmptyOutput) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *Service5EmptyOutput) Write(w []byte) []byte {
	return w
}

func (item *Service5EmptyOutput) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x11e46879); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *Service5EmptyOutput) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *Service5EmptyOutput) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x11e46879)
	return item.Write(w)
}

func (item Service5EmptyOutput) String() string {
	return string(item.WriteJSON(nil))
}

func (item *Service5EmptyOutput) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			return internal.ErrorInvalidJSON("service5.emptyOutput", "this object can't have properties")
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Service5EmptyOutput) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w), nil
}

func (item *Service5EmptyOutput) WriteJSON(w []byte) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}
func (item *Service5EmptyOutput) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) []byte {
	w = append(w, '{')
	return append(w, '}')
}

func (item *Service5EmptyOutput) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *Service5EmptyOutput) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("service5.emptyOutput", err.Error())
	}
	return nil
}

var _Service5Output = [2]internal.UnionElement{
	{TLTag: 0x11e46879, TLName: "service5.emptyOutput", TLString: "service5.emptyOutput#11e46879"},
	{TLTag: 0x179e9863, TLName: "service5.stringOutput", TLString: "service5.stringOutput#179e9863"},
}

type Service5Output struct {
	valueString Service5StringOutput
	index       int
}

func (item Service5Output) TLName() string { return _Service5Output[item.index].TLName }
func (item Service5Output) TLTag() uint32  { return _Service5Output[item.index].TLTag }

func (item *Service5Output) Reset() { item.index = 0 }

func (item *Service5Output) IsEmpty() bool { return item.index == 0 }

func (item *Service5Output) AsEmpty() (Service5EmptyOutput, bool) {
	var value Service5EmptyOutput
	return value, item.index == 0
}
func (item *Service5Output) ResetToEmpty() { item.index = 0 }
func (item *Service5Output) SetEmpty()     { item.index = 0 }

func (item *Service5Output) IsString() bool { return item.index == 1 }

func (item *Service5Output) AsString() (*Service5StringOutput, bool) {
	if item.index != 1 {
		return nil, false
	}
	return &item.valueString, true
}
func (item *Service5Output) ResetToString() *Service5StringOutput {
	item.index = 1
	item.valueString.Reset()
	return &item.valueString
}
func (item *Service5Output) SetString(value Service5StringOutput) {
	item.index = 1
	item.valueString = value
}

func (item *Service5Output) ReadBoxed(w []byte) (_ []byte, err error) {
	var tag uint32
	if w, err = basictl.NatRead(w, &tag); err != nil {
		return w, err
	}
	switch tag {
	case 0x11e46879:
		item.index = 0
		return w, nil
	case 0x179e9863:
		item.index = 1
		return item.valueString.Read(w)
	default:
		return w, internal.ErrorInvalidUnionTag("service5.Output", tag)
	}
}

func (item *Service5Output) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *Service5Output) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, _Service5Output[item.index].TLTag)
	switch item.index {
	case 0:
		return w
	case 1:
		w = item.valueString.Write(w)
	}
	return w
}

func (item *Service5Output) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	_tag, _value, err := internal.Json2ReadUnion("service5.Output", in)
	if err != nil {
		return err
	}
	switch _tag {
	case "service5.emptyOutput#11e46879", "service5.emptyOutput", "#11e46879":
		if !legacyTypeNames && _tag == "service5.emptyOutput#11e46879" {
			return internal.ErrorInvalidUnionLegacyTagJSON("service5.Output", "service5.emptyOutput#11e46879")
		}
		item.index = 0
	case "service5.stringOutput#179e9863", "service5.stringOutput", "#179e9863":
		if !legacyTypeNames && _tag == "service5.stringOutput#179e9863" {
			return internal.ErrorInvalidUnionLegacyTagJSON("service5.Output", "service5.stringOutput#179e9863")
		}
		item.index = 1
		var in2Pointer *basictl.JsonLexer
		if _value != nil {
			in2 := basictl.JsonLexer{Data: _value}
			in2Pointer = &in2
		}
		if err := item.valueString.ReadJSON(legacyTypeNames, in2Pointer); err != nil {
			return err
		}
	default:
		return internal.ErrorInvalidUnionTagJSON("service5.Output", _tag)
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Service5Output) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) ([]byte, error) {
	return item.WriteJSONOpt(tctx, w), nil
}

func (item *Service5Output) WriteJSON(w []byte) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}
func (item *Service5Output) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) []byte {
	switch item.index {
	case 0:
		if tctx.LegacyTypeNames {
			w = append(w, `{"type":"service5.emptyOutput#11e46879"`...)
		} else {
			w = append(w, `{"type":"service5.emptyOutput"`...)
		}
		return append(w, '}')
	case 1:
		if tctx.LegacyTypeNames {
			w = append(w, `{"type":"service5.stringOutput#179e9863"`...)
		} else {
			w = append(w, `{"type":"service5.stringOutput"`...)
		}
		w = append(w, `,"value":`...)
		w = item.valueString.WriteJSONOpt(tctx, w)
		return append(w, '}')
	default: // Impossible due to panic above
		return w
	}
}

func (item Service5Output) String() string {
	return string(item.WriteJSON(nil))
}

func (item *Service5Output) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *Service5Output) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("service5.Output", err.Error())
	}
	return nil
}

func (item Service5StringOutput) AsUnion() Service5Output {
	var ret Service5Output
	ret.SetString(item)
	return ret
}

type Service5StringOutput struct {
	HttpCode int32
	Response string
}

func (Service5StringOutput) TLName() string { return "service5.stringOutput" }
func (Service5StringOutput) TLTag() uint32  { return 0x179e9863 }

func (item *Service5StringOutput) Reset() {
	item.HttpCode = 0
	item.Response = ""
}

func (item *Service5StringOutput) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.IntRead(w, &item.HttpCode); err != nil {
		return w, err
	}
	return basictl.StringRead(w, &item.Response)
}

func (item *Service5StringOutput) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *Service5StringOutput) Write(w []byte) []byte {
	w = basictl.IntWrite(w, item.HttpCode)
	w = basictl.StringWrite(w, item.Response)
	return w
}

func (item *Service5StringOutput) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x179e9863); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *Service5StringOutput) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *Service5StringOutput) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x179e9863)
	return item.Write(w)
}

func (item Service5StringOutput) String() string {
	return string(item.WriteJSON(nil))
}

func (item *Service5StringOutput) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propHttpCodePresented bool
	var propResponsePresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "http_code":
				if propHttpCodePresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service5.stringOutput", "http_code")
				}
				if err := internal.Json2ReadInt32(in, &item.HttpCode); err != nil {
					return err
				}
				propHttpCodePresented = true
			case "response":
				if propResponsePresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service5.stringOutput", "response")
				}
				if err := internal.Json2ReadString(in, &item.Response); err != nil {
					return err
				}
				propResponsePresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("service5.stringOutput", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propHttpCodePresented {
		item.HttpCode = 0
	}
	if !propResponsePresented {
		item.Response = ""
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Service5StringOutput) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w), nil
}

func (item *Service5StringOutput) WriteJSON(w []byte) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}
func (item *Service5StringOutput) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) []byte {
	w = append(w, '{')
	backupIndexHttpCode := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"http_code":`...)
	w = basictl.JSONWriteInt32(w, item.HttpCode)
	if (item.HttpCode != 0) == false {
		w = w[:backupIndexHttpCode]
	}
	backupIndexResponse := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"response":`...)
	w = basictl.JSONWriteString(w, item.Response)
	if (len(item.Response) != 0) == false {
		w = w[:backupIndexResponse]
	}
	return append(w, '}')
}

func (item *Service5StringOutput) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *Service5StringOutput) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("service5.stringOutput", err.Error())
	}
	return nil
}
