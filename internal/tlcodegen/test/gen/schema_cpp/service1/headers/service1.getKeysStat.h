// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "service1/functions/service1.getKeysStat.h"
#include "service1/types/service1.keysStat.h"

namespace tlgen { namespace details { 

void Service1GetKeysStatReset(::tlgen::service1::GetKeysStat& item) noexcept;

bool Service1GetKeysStatWriteJSON(std::ostream& s, const ::tlgen::service1::GetKeysStat& item) noexcept;
bool Service1GetKeysStatRead(::tlgen::basictl::tl_istream & s, ::tlgen::service1::GetKeysStat& item) noexcept; 
bool Service1GetKeysStatWrite(::tlgen::basictl::tl_ostream & s, const ::tlgen::service1::GetKeysStat& item) noexcept;
bool Service1GetKeysStatReadBoxed(::tlgen::basictl::tl_istream & s, ::tlgen::service1::GetKeysStat& item);
bool Service1GetKeysStatWriteBoxed(::tlgen::basictl::tl_ostream & s, const ::tlgen::service1::GetKeysStat& item);

bool Service1GetKeysStatReadResult(::tlgen::basictl::tl_istream & s, const ::tlgen::service1::GetKeysStat& item, std::optional<::tlgen::service1::KeysStat>& result);
bool Service1GetKeysStatWriteResult(::tlgen::basictl::tl_ostream & s, const ::tlgen::service1::GetKeysStat& item, const std::optional<::tlgen::service1::KeysStat>& result);
    
}} // namespace tlgen::details

