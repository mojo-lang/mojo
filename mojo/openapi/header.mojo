// Copyright 2021 Mojo-lang.org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/// The Header Object follows the structure of the Parameter Object with the following changes:
/// 1. name MUST NOT be specified, it is given in the corresponding headers map.
/// 1. in MUST NOT be specified, it is implicitly in header.
/// 1. All traits that are affected by the location MUST be applicable to a location of header (for example, style).
/// @example({
///    description: "The number of allowed requests in the current period"
///    schema: {
///        type: integer
///    }
///})
type Header : Parameter @ignore_fields('name', 'in')