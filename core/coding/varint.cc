//
// Copyright © 2018 The IPFN Developers. All Rights Reserved.
// Copyright © 2001-2018 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS-IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
#include <ipfn/util/coding/varint.hpp>

namespace varint {

const char *
parse32(const char *p, const char *l, uint32_t *OUTPUT) {
  return parse32_inline(p, l, OUTPUT);
}

const char *
parse64(const char *p, uint64_t *OUTPUT) {
  return parse64_inline(p, OUTPUT);
}

const char *
parse64_limit(const char *p, const char *l, uint64_t *OUTPUT) {
  return parse64_limit_inline(p, l, OUTPUT);
}

char *
encode32(char *sptr, uint32_t v) {
  return encode32_inline(sptr, v);
}

char *
encode64(char *sptr, uint64_t v) {
  return encode64_inline(sptr, v);
}

}  // namespace varint
