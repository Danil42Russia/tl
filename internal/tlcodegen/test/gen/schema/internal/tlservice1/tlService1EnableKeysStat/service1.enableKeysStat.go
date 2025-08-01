// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlService1EnableKeysStat

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlBool"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type Service1EnableKeysStat struct {
	Period int32
}

func (Service1EnableKeysStat) TLName() string { return "service1.enableKeysStat" }
func (Service1EnableKeysStat) TLTag() uint32  { return 0x29a7090e }

func (item *Service1EnableKeysStat) Reset() {
	item.Period = 0
}

func (item *Service1EnableKeysStat) Read(w []byte) (_ []byte, err error) {
	return basictl.IntRead(w, &item.Period)
}

func (item *Service1EnableKeysStat) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *Service1EnableKeysStat) Write(w []byte) []byte {
	w = basictl.IntWrite(w, item.Period)
	return w
}

func (item *Service1EnableKeysStat) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x29a7090e); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *Service1EnableKeysStat) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *Service1EnableKeysStat) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x29a7090e)
	return item.Write(w)
}

func (item *Service1EnableKeysStat) ReadResult(w []byte, ret *bool) (_ []byte, err error) {
	return tlBool.BoolReadBoxed(w, ret)
}

func (item *Service1EnableKeysStat) WriteResult(w []byte, ret bool) (_ []byte, err error) {
	w = tlBool.BoolWriteBoxed(w, ret)
	return w, nil
}

func (item *Service1EnableKeysStat) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *bool) error {
	if err := internal.Json2ReadBool(in, ret); err != nil {
		return err
	}
	return nil
}

func (item *Service1EnableKeysStat) WriteResultJSON(w []byte, ret bool) (_ []byte, err error) {
	tctx := basictl.JSONWriteContext{}
	return item.writeResultJSON(&tctx, w, ret)
}

func (item *Service1EnableKeysStat) writeResultJSON(tctx *basictl.JSONWriteContext, w []byte, ret bool) (_ []byte, err error) {
	w = basictl.JSONWriteBool(w, ret)
	return w, nil
}

func (item *Service1EnableKeysStat) ReadResultWriteResultJSON(tctx *basictl.JSONWriteContext, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret bool
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(tctx, w, ret)
	return r, w, err
}

func (item *Service1EnableKeysStat) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret bool
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item Service1EnableKeysStat) String() string {
	return string(item.WriteJSON(nil))
}

func (item *Service1EnableKeysStat) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propPeriodPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "period":
				if propPeriodPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service1.enableKeysStat", "period")
				}
				if err := internal.Json2ReadInt32(in, &item.Period); err != nil {
					return err
				}
				propPeriodPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("service1.enableKeysStat", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propPeriodPresented {
		item.Period = 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Service1EnableKeysStat) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w), nil
}

func (item *Service1EnableKeysStat) WriteJSON(w []byte) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}
func (item *Service1EnableKeysStat) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) []byte {
	w = append(w, '{')
	backupIndexPeriod := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"period":`...)
	w = basictl.JSONWriteInt32(w, item.Period)
	if (item.Period != 0) == false {
		w = w[:backupIndexPeriod]
	}
	return append(w, '}')
}

func (item *Service1EnableKeysStat) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *Service1EnableKeysStat) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("service1.enableKeysStat", err.Error())
	}
	return nil
}
