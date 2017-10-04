#
# Copyright 2017 The IPFN Authors. All Rights Reserved.
# 
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# 
#     http://www.apache.org/licenses/LICENSE-2.0
# 
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#gra

@0xaec52273ad890df1;

struct Tensor {
  value :union {
    voids     @0  :List(Void);
    bools     @1  :List(Bool);
    int8s     @2  :List(Int8);
    int16s    @3  :List(Int16);
    int32s    @4  :List(Int32);
    int64s    @5  :List(Int64);
    uint8s    @6  :List(UInt8);
    uint16s   @7  :List(UInt16);
    uint32s   @8  :List(UInt32);
    uint64s   @9  :List(UInt64);
    float32s  @10 :List(Float32);
    float64s  @11 :List(Float64);
    text      @12 :List(Text);
    data      @13 :List(Data);
    list      @14 :List(Tensor);
  }
}
