// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlservice1

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/cycle_6ca945392bbf8b14f24e5653edc8b214"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlDictionaryFieldService1Value"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlDictionaryService1Value"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlService1KeysStatMaybe"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlService1ValueBoxedMaybe"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlTupleService1Value3"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlVectorDictionaryFieldService1Value"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlVectorService1Value"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tlservice1/tlService1Add"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tlservice1/tlService1AddOrGet"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tlservice1/tlService1AddOrIncr"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tlservice1/tlService1Append"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tlservice1/tlService1Cas"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tlservice1/tlService1Decr"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tlservice1/tlService1Delete"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tlservice1/tlService1DisableExpiration"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tlservice1/tlService1DisableKeysStat"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tlservice1/tlService1EnableExpiration"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tlservice1/tlService1EnableKeysStat"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tlservice1/tlService1Exists"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tlservice1/tlService1Get"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tlservice1/tlService1GetExpireTime"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tlservice1/tlService1GetKeysStat"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tlservice1/tlService1GetKeysStatPeriods"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tlservice1/tlService1GetWildcard"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tlservice1/tlService1GetWildcardDict"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tlservice1/tlService1GetWildcardList"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tlservice1/tlService1GetWildcardWithFlags"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tlservice1/tlService1Incr"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tlservice1/tlService1KeysStat"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tlservice1/tlService1Replace"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tlservice1/tlService1ReplaceOrIncr"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tlservice1/tlService1Set"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tlservice1/tlService1SetOrIncr"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tlservice1/tlService1Touch"
)

type (
	Add                                = tlService1Add.Service1Add
	AddOrGet                           = tlService1AddOrGet.Service1AddOrGet
	AddOrIncr                          = tlService1AddOrIncr.Service1AddOrIncr
	Append                             = tlService1Append.Service1Append
	Cas                                = tlService1Cas.Service1Cas
	Decr                               = tlService1Decr.Service1Decr
	Delete                             = tlService1Delete.Service1Delete
	DictionaryFieldService1Value       = tlDictionaryFieldService1Value.DictionaryFieldService1Value
	DictionaryService1Value            = tlDictionaryService1Value.DictionaryService1Value
	DisableExpiration                  = tlService1DisableExpiration.Service1DisableExpiration
	DisableKeysStat                    = tlService1DisableKeysStat.Service1DisableKeysStat
	EnableExpiration                   = tlService1EnableExpiration.Service1EnableExpiration
	EnableKeysStat                     = tlService1EnableKeysStat.Service1EnableKeysStat
	Exists                             = tlService1Exists.Service1Exists
	Get                                = tlService1Get.Service1Get
	GetExpireTime                      = tlService1GetExpireTime.Service1GetExpireTime
	GetKeysStat                        = tlService1GetKeysStat.Service1GetKeysStat
	GetKeysStatPeriods                 = tlService1GetKeysStatPeriods.Service1GetKeysStatPeriods
	GetWildcard                        = tlService1GetWildcard.Service1GetWildcard
	GetWildcardDict                    = tlService1GetWildcardDict.Service1GetWildcardDict
	GetWildcardList                    = tlService1GetWildcardList.Service1GetWildcardList
	GetWildcardWithFlags               = tlService1GetWildcardWithFlags.Service1GetWildcardWithFlags
	Incr                               = tlService1Incr.Service1Incr
	KeysStat                           = tlService1KeysStat.Service1KeysStat
	KeysStatMaybe                      = tlService1KeysStatMaybe.Service1KeysStatMaybe
	Longvalue                          = cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Longvalue
	LongvalueWithTime                  = cycle_6ca945392bbf8b14f24e5653edc8b214.Service1LongvalueWithTime
	NotFound                           = cycle_6ca945392bbf8b14f24e5653edc8b214.Service1NotFound
	Replace                            = tlService1Replace.Service1Replace
	ReplaceOrIncr                      = tlService1ReplaceOrIncr.Service1ReplaceOrIncr
	Set                                = tlService1Set.Service1Set
	SetOrIncr                          = tlService1SetOrIncr.Service1SetOrIncr
	Strvalue                           = cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Strvalue
	StrvalueWithTime                   = cycle_6ca945392bbf8b14f24e5653edc8b214.Service1StrvalueWithTime
	Touch                              = tlService1Touch.Service1Touch
	TupleService1Value3                = tlTupleService1Value3.TupleService1Value3
	Value                              = cycle_6ca945392bbf8b14f24e5653edc8b214.Service1Value
	ValueBoxedMaybe                    = tlService1ValueBoxedMaybe.Service1ValueBoxedMaybe
	VectorDictionaryFieldService1Value = tlVectorDictionaryFieldService1Value.VectorDictionaryFieldService1Value
	VectorService1Value                = tlVectorService1Value.VectorService1Value
)
