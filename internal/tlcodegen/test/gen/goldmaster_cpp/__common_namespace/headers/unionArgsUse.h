// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/unionArgsUse.h"

namespace tlgen { namespace details { 

void UnionArgsUseReset(::tlgen::UnionArgsUse& item) noexcept;

bool UnionArgsUseWriteJSON(std::ostream& s, const ::tlgen::UnionArgsUse& item) noexcept;
bool UnionArgsUseRead(::tlgen::basictl::tl_istream & s, ::tlgen::UnionArgsUse& item) noexcept; 
bool UnionArgsUseWrite(::tlgen::basictl::tl_ostream & s, const ::tlgen::UnionArgsUse& item) noexcept;
bool UnionArgsUseReadBoxed(::tlgen::basictl::tl_istream & s, ::tlgen::UnionArgsUse& item);
bool UnionArgsUseWriteBoxed(::tlgen::basictl::tl_ostream & s, const ::tlgen::UnionArgsUse& item);

}} // namespace tlgen::details

