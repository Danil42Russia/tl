// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/replace18.h"

namespace tlgen { namespace details { 

void Replace18Reset(::tlgen::Replace18& item) noexcept;

bool Replace18WriteJSON(std::ostream& s, const ::tlgen::Replace18& item) noexcept;
bool Replace18Read(::tlgen::basictl::tl_istream & s, ::tlgen::Replace18& item) noexcept; 
bool Replace18Write(::tlgen::basictl::tl_ostream & s, const ::tlgen::Replace18& item) noexcept;
bool Replace18ReadBoxed(::tlgen::basictl::tl_istream & s, ::tlgen::Replace18& item);
bool Replace18WriteBoxed(::tlgen::basictl::tl_ostream & s, const ::tlgen::Replace18& item);

}} // namespace tlgen::details

