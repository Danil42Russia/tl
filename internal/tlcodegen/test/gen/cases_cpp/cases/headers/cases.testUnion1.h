// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "cases/types/cases.testUnion1.h"

namespace tlgen { namespace details { 

void CasesTestUnion1Reset(::tlgen::cases::TestUnion1& item) noexcept;

bool CasesTestUnion1WriteJSON(std::ostream& s, const ::tlgen::cases::TestUnion1& item) noexcept;
bool CasesTestUnion1Read(::tlgen::basictl::tl_istream & s, ::tlgen::cases::TestUnion1& item) noexcept; 
bool CasesTestUnion1Write(::tlgen::basictl::tl_ostream & s, const ::tlgen::cases::TestUnion1& item) noexcept;
bool CasesTestUnion1ReadBoxed(::tlgen::basictl::tl_istream & s, ::tlgen::cases::TestUnion1& item);
bool CasesTestUnion1WriteBoxed(::tlgen::basictl::tl_ostream & s, const ::tlgen::cases::TestUnion1& item);

}} // namespace tlgen::details

