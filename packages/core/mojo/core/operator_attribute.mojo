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

/// attribute for the operator functions
///
/// #example
/// ```
/// @infix
/// func +(left: String, right: String) -> String
///
/// // r'.*'
/// @literal
/// @prefix
/// func r(str: String) -> Regex
///
/// // 5ms
/// @literal
/// @postfix
/// func ms(number: Int) -> Duration
/// ```
attribute prefix: Bool

/// attribute for the postfix operator functions
attribute postfix: Bool

/// attribute for the infix operator functions
attribute infix: Bool

/// attribute for the literal operator functions
attribute literal: Bool
