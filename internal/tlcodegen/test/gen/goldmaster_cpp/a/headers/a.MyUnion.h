// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "a/types/a.MyUnion.h"

namespace tlgen { namespace details { 

void AMyUnionReset(::tlgen::a::MyUnion& item) noexcept;

bool AMyUnionWriteJSON(std::ostream & s, const ::tlgen::a::MyUnion& item) noexcept;
bool AMyUnionReadBoxed(::tlgen::basictl::tl_istream & s, ::tlgen::a::MyUnion& item) noexcept;
bool AMyUnionWriteBoxed(::tlgen::basictl::tl_ostream & s, const ::tlgen::a::MyUnion& item) noexcept;

}} // namespace tlgen::details

