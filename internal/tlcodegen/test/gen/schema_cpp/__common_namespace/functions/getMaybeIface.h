// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "service1/types/service1.Value.h"


namespace tlgen { 
struct GetMaybeIface {
  ::tlgen::service1::Value x;

  // tl type info
  static constexpr uint32_t TL_TAG = 0x6b055ae4;
  static constexpr std::string_view TL_NAME = "getMaybeIface";

  uint32_t tl_tag() const { return 0x6b055ae4; }
  std::string_view tl_name() const { return "getMaybeIface"; }

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
  using ResultType = std::optional<::tlgen::service1::Value>;

  bool read_result(::tlgen::basictl::tl_istream & s, std::optional<::tlgen::service1::Value> & result) const noexcept;
  bool write_result(::tlgen::basictl::tl_ostream & s, const std::optional<::tlgen::service1::Value> & result) const noexcept;

  void read_result(::tlgen::basictl::tl_throwable_istream & s, std::optional<::tlgen::service1::Value> & result) const;
  void write_result(::tlgen::basictl::tl_throwable_ostream & s, const std::optional<::tlgen::service1::Value> & result) const;

  friend std::ostream& operator<<(std::ostream& s, const GetMaybeIface& rhs) {
    rhs.write_json(s);
    return s;
  }
};

} // namespace tlgen

