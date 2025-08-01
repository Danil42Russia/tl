// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"


namespace tlgen { 
struct BoxedTuple {
  std::array<int32_t, 3> x{};

  // tl type info
  static constexpr uint32_t TL_TAG = 0x30c9d533;
  static constexpr std::string_view TL_NAME = "boxedTuple";

  uint32_t tl_tag() const { return 0x30c9d533; }
  std::string_view tl_name() const { return "boxedTuple"; }

  // basic serialization methods 
  bool write_json(std::ostream& s) const;

  bool read(::tlgen::basictl::tl_istream & s) noexcept;
  bool write(::tlgen::basictl::tl_ostream & s) const noexcept;

  void read(::tlgen::basictl::tl_throwable_istream & s);
  void write(::tlgen::basictl::tl_throwable_ostream & s) const;

  bool read_boxed(::tlgen::basictl::tl_istream & s) noexcept;
  bool write_boxed(::tlgen::basictl::tl_ostream & s) const noexcept;
  
  void read_boxed(::tlgen::basictl::tl_throwable_istream & s);
  void write_boxed(::tlgen::basictl::tl_throwable_ostream & s) const;

  // function methods and properties
  using ResultType = std::array<int32_t, 3>;

  bool read_result(::tlgen::basictl::tl_istream & s, std::array<int32_t, 3> & result) const noexcept;
  bool write_result(::tlgen::basictl::tl_ostream & s, const std::array<int32_t, 3> & result) const noexcept;

  void read_result(::tlgen::basictl::tl_throwable_istream & s, std::array<int32_t, 3> & result) const;
  void write_result(::tlgen::basictl::tl_throwable_ostream & s, const std::array<int32_t, 3> & result) const;

  friend std::ostream& operator<<(std::ostream& s, const BoxedTuple& rhs) {
    rhs.write_json(s);
    return s;
  }
};

} // namespace tlgen

