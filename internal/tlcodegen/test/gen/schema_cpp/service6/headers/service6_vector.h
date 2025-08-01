// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/vector.h"
#include "service6/types/service6.findWithBoundsResult.h"
#include "service6/types/service6.findResultRow.h"
#include "service6/types/service6.error.h"
#include "__common_namespace/types/Either.h"

namespace tlgen { namespace details { 

void VectorEitherService6ErrorVectorService6FindResultRowReset(std::vector<::tlgen::Either<::tlgen::service6::Error, std::vector<::tlgen::service6::FindResultRow>>>& item) noexcept;

bool VectorEitherService6ErrorVectorService6FindResultRowWriteJSON(std::ostream& s, const std::vector<::tlgen::Either<::tlgen::service6::Error, std::vector<::tlgen::service6::FindResultRow>>>& item) noexcept;
bool VectorEitherService6ErrorVectorService6FindResultRowRead(::tlgen::basictl::tl_istream & s, std::vector<::tlgen::Either<::tlgen::service6::Error, std::vector<::tlgen::service6::FindResultRow>>>& item) noexcept; 
bool VectorEitherService6ErrorVectorService6FindResultRowWrite(::tlgen::basictl::tl_ostream & s, const std::vector<::tlgen::Either<::tlgen::service6::Error, std::vector<::tlgen::service6::FindResultRow>>>& item) noexcept;
bool VectorEitherService6ErrorVectorService6FindResultRowReadBoxed(::tlgen::basictl::tl_istream & s, std::vector<::tlgen::Either<::tlgen::service6::Error, std::vector<::tlgen::service6::FindResultRow>>>& item);
bool VectorEitherService6ErrorVectorService6FindResultRowWriteBoxed(::tlgen::basictl::tl_ostream & s, const std::vector<::tlgen::Either<::tlgen::service6::Error, std::vector<::tlgen::service6::FindResultRow>>>& item);

}} // namespace tlgen::details

namespace tlgen { namespace details { 

void VectorService6FindResultRowReset(std::vector<::tlgen::service6::FindResultRow>& item) noexcept;

bool VectorService6FindResultRowWriteJSON(std::ostream& s, const std::vector<::tlgen::service6::FindResultRow>& item) noexcept;
bool VectorService6FindResultRowRead(::tlgen::basictl::tl_istream & s, std::vector<::tlgen::service6::FindResultRow>& item) noexcept; 
bool VectorService6FindResultRowWrite(::tlgen::basictl::tl_ostream & s, const std::vector<::tlgen::service6::FindResultRow>& item) noexcept;
bool VectorService6FindResultRowReadBoxed(::tlgen::basictl::tl_istream & s, std::vector<::tlgen::service6::FindResultRow>& item);
bool VectorService6FindResultRowWriteBoxed(::tlgen::basictl::tl_ostream & s, const std::vector<::tlgen::service6::FindResultRow>& item);

}} // namespace tlgen::details

namespace tlgen { namespace details { 

void VectorService6FindWithBoundsResultReset(std::vector<::tlgen::service6::FindWithBoundsResult>& item) noexcept;

bool VectorService6FindWithBoundsResultWriteJSON(std::ostream& s, const std::vector<::tlgen::service6::FindWithBoundsResult>& item) noexcept;
bool VectorService6FindWithBoundsResultRead(::tlgen::basictl::tl_istream & s, std::vector<::tlgen::service6::FindWithBoundsResult>& item) noexcept; 
bool VectorService6FindWithBoundsResultWrite(::tlgen::basictl::tl_ostream & s, const std::vector<::tlgen::service6::FindWithBoundsResult>& item) noexcept;
bool VectorService6FindWithBoundsResultReadBoxed(::tlgen::basictl::tl_istream & s, std::vector<::tlgen::service6::FindWithBoundsResult>& item);
bool VectorService6FindWithBoundsResultWriteBoxed(::tlgen::basictl::tl_ostream & s, const std::vector<::tlgen::service6::FindWithBoundsResult>& item);

}} // namespace tlgen::details

