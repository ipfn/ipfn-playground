/**
 * Copyright © 2017-2018 The IPFN Developers. All Rights Reserved.
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

#include <sstream>
#include <string>

#include "ipfn/framework/cell_path.h"


namespace ipfn {

// cell is a basic generic cell interface required for storage and
// transport, ...
class cell {
 public:
  virtual const std::optional<cell_path> &name() const = 0;
  virtual const std::optional<cell_path> &soul() const = 0;
};

class basic_cell : public cell {
 public:
  explicit basic_cell(std::string name, std::string soul);
  explicit basic_cell(const std::string& name, const cell_path& soul);
  
  virtual const std::optional<cell_path> &
  name() const {
    return name_;
  }

  virtual const std::optional<cell_path> &
  soul() const {
    return soul_;
  }

 private:
  const std::optional<cell_path> name_;
  const std::optional<cell_path> soul_;
};

basic_cell::basic_cell(std::string name, std::string soul)
  : name_(std::move(name)), soul_(std::move(soul)) {};

basic_cell::basic_cell(const std::string& name, const cell_path& soul)
  : name_(name), soul_(soul) {};

inline std::ostream &
operator<<(std::ostream &o, const cell &c) noexcept {
  o << "[ ";
  const auto& name = c.name();
  if (name) o << "name=" << *name << " ";
  const auto& soul = c.soul();
  if (soul) o << "soul=" << *soul;
  o << " ]";
  return o;
}

}  // namespace ipfn
