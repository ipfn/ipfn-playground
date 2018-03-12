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

#include <list>
#include <string>

#include <boost/algorithm/string.hpp>

#include "ipfn/platform/experimental.h"


namespace ipfn {

class cell_path {
 public:
  using list = std::list<boost::iterator_range<std::string::iterator>>;

  // NOTE: expensive split (TODO?)
  cell_path(std::string);
  cell_path(std::string, list);

  /// \brief splits path into vector of elements and constructs path
  inline static cell_path split(std::string path) noexcept;

  inline const std::string &
  as_str() const {
    return path_;
  }

  inline const list &
  elements() const {
    return elements_;
  }

  inline const std::size_t
  size() const {
    return elements_.size();
  }

  inline const list &operator*() const { return elements_; }

 private:
  const std::string path_;
  const list elements_;
};

cell_path::list
split_path(std::string path) noexcept {
  cell_path::list elements;
  boost::split(elements, path, boost::is_any_of("/"));
  return elements;
}

cell_path
cell_path::split(std::string path) noexcept {
  cell_path::list elements = split_path(path);
  return cell_path(std::move(path), elements);
}

cell_path::cell_path(std::string path)
  : path_(path), elements_(split_path(path)){};

cell_path::cell_path(std::string path, list elements)
  : path_(std::move(path)), elements_(std::move(elements)){};

inline std::ostream &
operator<<(std::ostream &o, const cell_path &path) noexcept {
  return o << path.as_str();
}

}  // namespace ipfn
