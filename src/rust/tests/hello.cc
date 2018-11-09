/**
 * Copyright © 2018 The IPFN Developers. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
#include <ipfn/core.h>
#include <ipfn/rust/ffi.h>

#include <iostream>

using namespace ipfn::ffi;

int
main(int argc, char *argv[]) {
  ipfn_test();
  // std::cout << addition(2, 2) << std::endl;
  auto song = theme_song_generate("test");
  std::cout << song << std::endl << std::flush;
  theme_song_free(song);
  std::cout << "hello.cc" << std::endl;
}
