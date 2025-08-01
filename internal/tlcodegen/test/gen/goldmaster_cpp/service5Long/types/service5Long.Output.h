// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "service5Long/types/service5Long.stringOutput.h"
#include "service5Long/types/service5Long.emptyOutput.h"


namespace tlgen { namespace service5Long { 
struct Output {
  std::variant<::tlgen::service5Long::EmptyOutput, ::tlgen::service5Long::StringOutput> value;

  bool is_empty() const { return value.index() == 0; }
  bool is_string() const { return value.index() == 1; }

  void set_empty() { value.emplace<0>(); }

  std::string_view tl_name() const;
  uint32_t tl_tag() const;

  bool write_json(std::ostream& s)const;

  bool read_boxed(::tlgen::basictl::tl_istream & s) noexcept;
  bool write_boxed(::tlgen::basictl::tl_ostream & s)const noexcept;
  
  void read_boxed(::tlgen::basictl::tl_throwable_istream & s);
  void write_boxed(::tlgen::basictl::tl_throwable_ostream & s)const;
};

}} // namespace tlgen::service5Long

