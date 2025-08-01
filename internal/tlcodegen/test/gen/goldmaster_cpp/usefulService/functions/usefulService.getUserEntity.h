// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "usefulService/types/usefulService.getUserEntityResult.h"


namespace tlgen { namespace usefulService { 
struct GetUserEntity {
  uint32_t fields_mask = 0;
  std::string stage_id;

  // tl type info
  static constexpr uint32_t TL_TAG = 0x3c857e52;
  static constexpr std::string_view TL_NAME = "usefulService.getUserEntity";

  uint32_t tl_tag() const { return 0x3c857e52; }
  std::string_view tl_name() const { return "usefulService.getUserEntity"; }

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
  using ResultType = std::optional<::tlgen::usefulService::GetUserEntityResult>;

  bool read_result(::tlgen::basictl::tl_istream & s, std::optional<::tlgen::usefulService::GetUserEntityResult> & result) const noexcept;
  bool write_result(::tlgen::basictl::tl_ostream & s, const std::optional<::tlgen::usefulService::GetUserEntityResult> & result) const noexcept;

  void read_result(::tlgen::basictl::tl_throwable_istream & s, std::optional<::tlgen::usefulService::GetUserEntityResult> & result) const;
  void write_result(::tlgen::basictl::tl_throwable_ostream & s, const std::optional<::tlgen::usefulService::GetUserEntityResult> & result) const;

  friend std::ostream& operator<<(std::ostream& s, const GetUserEntity& rhs) {
    rhs.write_json(s);
    return s;
  }
};

}} // namespace tlgen::usefulService

