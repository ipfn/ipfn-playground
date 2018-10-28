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
#include <iostream>
#include <memory>
#include <secp256k1.h>
#include <secp256k1/hash.h>
#include <secp256k1/hash_impl.h>
#include <cstring>

int main()
{
    const char *data_ = "test";
    secp256k1_sha256 ctx;
    secp256k1_sha256_initialize(&ctx);
    secp256k1_sha256_write(&ctx, reinterpret_cast<const unsigned char *>(data_), strlen(data_));
    unsigned char hash[32];
    secp256k1_sha256_finalize(&ctx, hash);

    std::cout << hash << std::endl;
}