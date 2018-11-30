//
// Copyright © 2018 The IPFN Developers. All Rights Reserved.
// Copyright © 2015-2018 Topology LP. All Rights Reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL
// THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
// IN THE SOFTWARE.
//
#pragma once

#include <cppcodec/base32_crockford.hpp>

namespace cppcodec {

namespace detail {

// RFC 4648 uses a simple alphabet: A-Z starting at index 0, then 2-7 starting
// at index 26.
static constexpr const char base32_ipfn_alphabet[] = {
  '0', 'p', 'z', 'q', 'y', '9', 'x', '8', 'b', 'f', '2', 't', 'v', 'r',
  'w', 'd', 's', '3', 'j', 'n', '5', '4', 'k', 'h', 'c', 'e',  // at index 26
  '6', 'm', 'u', 'a', '7', 'l'};

class base32_ipfn : public cppcodec::base32_crockford {
 public:
  template <typename Codec>
  using codec_impl = stream_codec<Codec, base32_ipfn>;

  static CPPCODEC_ALWAYS_INLINE constexpr size_t alphabet_size() {
    static_assert(sizeof(base32_ipfn_alphabet) == 32,
                  "base32 alphabet must have 32 values");
    return sizeof(base32_ipfn_alphabet);
  }
  static CPPCODEC_ALWAYS_INLINE constexpr char symbol(alphabet_index_t idx) {
    return base32_ipfn_alphabet[idx];
  }
  static CPPCODEC_ALWAYS_INLINE constexpr char normalized_symbol(char c) {
    return c | 32;
  }
  static CPPCODEC_ALWAYS_INLINE constexpr bool generates_padding() {
    return false;
  }
  static CPPCODEC_ALWAYS_INLINE constexpr bool requires_padding() {
    return false;
  }
  static CPPCODEC_ALWAYS_INLINE constexpr bool is_padding_symbol(char) {
    return false;
  }
  static CPPCODEC_ALWAYS_INLINE constexpr bool is_eof_symbol(char c) {
    return c == '\0';
  }
  static CPPCODEC_ALWAYS_INLINE constexpr bool should_ignore(char c) {
    return c == '-';  // "Hyphens (-) can be inserted into strings [for
                      // readability]."
  }
};

}  // namespace detail

using base32_ipfn = detail::codec<detail::base32<detail::base32_ipfn>>;

}  // namespace cppcodec
