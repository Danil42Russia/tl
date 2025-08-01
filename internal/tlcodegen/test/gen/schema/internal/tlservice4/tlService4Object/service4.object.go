// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlService4Object

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlBuiltinVectorInt"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type Service4Object struct {
	Type     int32
	JointId  []int32
	ObjectId []int32
}

func (Service4Object) TLName() string { return "service4.object" }
func (Service4Object) TLTag() uint32  { return 0xa6eeca4f }

func (item *Service4Object) Reset() {
	item.Type = 0
	item.JointId = item.JointId[:0]
	item.ObjectId = item.ObjectId[:0]
}

func (item *Service4Object) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.IntRead(w, &item.Type); err != nil {
		return w, err
	}
	if w, err = tlBuiltinVectorInt.BuiltinVectorIntRead(w, &item.JointId); err != nil {
		return w, err
	}
	return tlBuiltinVectorInt.BuiltinVectorIntRead(w, &item.ObjectId)
}

func (item *Service4Object) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *Service4Object) Write(w []byte) []byte {
	w = basictl.IntWrite(w, item.Type)
	w = tlBuiltinVectorInt.BuiltinVectorIntWrite(w, item.JointId)
	w = tlBuiltinVectorInt.BuiltinVectorIntWrite(w, item.ObjectId)
	return w
}

func (item *Service4Object) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xa6eeca4f); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *Service4Object) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *Service4Object) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xa6eeca4f)
	return item.Write(w)
}

func (item Service4Object) String() string {
	return string(item.WriteJSON(nil))
}

func (item *Service4Object) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propTypePresented bool
	var propJointIdPresented bool
	var propObjectIdPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "type":
				if propTypePresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service4.object", "type")
				}
				if err := internal.Json2ReadInt32(in, &item.Type); err != nil {
					return err
				}
				propTypePresented = true
			case "joint_id":
				if propJointIdPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service4.object", "joint_id")
				}
				if err := tlBuiltinVectorInt.BuiltinVectorIntReadJSON(legacyTypeNames, in, &item.JointId); err != nil {
					return err
				}
				propJointIdPresented = true
			case "object_id":
				if propObjectIdPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("service4.object", "object_id")
				}
				if err := tlBuiltinVectorInt.BuiltinVectorIntReadJSON(legacyTypeNames, in, &item.ObjectId); err != nil {
					return err
				}
				propObjectIdPresented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("service4.object", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propTypePresented {
		item.Type = 0
	}
	if !propJointIdPresented {
		item.JointId = item.JointId[:0]
	}
	if !propObjectIdPresented {
		item.ObjectId = item.ObjectId[:0]
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *Service4Object) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w), nil
}

func (item *Service4Object) WriteJSON(w []byte) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}
func (item *Service4Object) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) []byte {
	w = append(w, '{')
	backupIndexType := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"type":`...)
	w = basictl.JSONWriteInt32(w, item.Type)
	if (item.Type != 0) == false {
		w = w[:backupIndexType]
	}
	backupIndexJointId := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"joint_id":`...)
	w = tlBuiltinVectorInt.BuiltinVectorIntWriteJSONOpt(tctx, w, item.JointId)
	if (len(item.JointId) != 0) == false {
		w = w[:backupIndexJointId]
	}
	backupIndexObjectId := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"object_id":`...)
	w = tlBuiltinVectorInt.BuiltinVectorIntWriteJSONOpt(tctx, w, item.ObjectId)
	if (len(item.ObjectId) != 0) == false {
		w = w[:backupIndexObjectId]
	}
	return append(w, '}')
}

func (item *Service4Object) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *Service4Object) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("service4.object", err.Error())
	}
	return nil
}
