//
// Copyright © 2018 The IPFN Developers. All Rights Reserved.
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

#include <base-x.hpp>

namespace ipfn {
/**
 * @brief  Base58 codec.
 */
namespace base58 {

template <typename Result = std::string>
Result
encode(std::vector<uint8_t> &decoded) {
  return Base58::base58().encode(reinterpret_cast<const char *>(decoded.data()),
                                 decoded.size());
}

template <typename Result = std::string, typename T>
Result
encode(T &decoded) {
  return Base58::base58().encode(decoded);
}

template <typename Result = std::string>
Result
encode(std::vector<char> &decoded) {
  return Base58::base58().encode(decoded.data(), decoded.size());
}

template <typename Result = std::string, typename T>
Result
decode(T &decoded) {
  return Base58::base58().decode(decoded);
}

}  // namespace base58

}  // namespace ipfn
