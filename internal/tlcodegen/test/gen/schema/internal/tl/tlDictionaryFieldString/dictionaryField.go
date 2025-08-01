// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlDictionaryFieldString

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type DictionaryFieldString struct {
	Key   string
	Value string
}

func (DictionaryFieldString) TLName() string { return "dictionaryField" }
func (DictionaryFieldString) TLTag() uint32  { return 0x239c1b62 }

func (item *DictionaryFieldString) Reset() {
	item.Key = ""
	item.Value = ""
}

func (item *DictionaryFieldString) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.StringRead(w, &item.Key); err != nil {
		return w, err
	}
	return basictl.StringRead(w, &item.Value)
}

func (item *DictionaryFieldString) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *DictionaryFieldString) Write(w []byte) []byte {
	w = basictl.StringWrite(w, item.Key)
	w = basictl.StringWrite(w, item.Value)
	return w
}

func (item *DictionaryFieldString) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x239c1b62); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *DictionaryFieldString) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *DictionaryFieldString) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x239c1b62)
	return item.Write(w)
}

func (item DictionaryFieldString) String() string {
	return string(item.WriteJSON(nil))
}

func (item *DictionaryFieldString) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propKeyPresented bool
	var propValuePresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "key":
				if propKeyPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("dictionaryField", "key")
				}
				if err := internal.Json2ReadString(in, &item.Key); err != nil {
					return err
				}
				propKeyPresented = true
			case "value":
				if propValuePresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("dictionaryField", "value")
				}
				if err := internal.Json2ReadString(in, &item.Value); err != nil {
					return err
				}
				propValuePresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("dictionaryField", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propKeyPresented {
		item.Key = ""
	}
	if !propValuePresented {
		item.Value = ""
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *DictionaryFieldString) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w), nil
}

func (item *DictionaryFieldString) WriteJSON(w []byte) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}
func (item *DictionaryFieldString) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) []byte {
	w = append(w, '{')
	backupIndexKey := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"key":`...)
	w = basictl.JSONWriteString(w, item.Key)
	if (len(item.Key) != 0) == false {
		w = w[:backupIndexKey]
	}
	backupIndexValue := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"value":`...)
	w = basictl.JSONWriteString(w, item.Value)
	if (len(item.Value) != 0) == false {
		w = w[:backupIndexValue]
	}
	return append(w, '}')
}

func (item *DictionaryFieldString) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *DictionaryFieldString) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("dictionaryField", err.Error())
	}
	return nil
}
