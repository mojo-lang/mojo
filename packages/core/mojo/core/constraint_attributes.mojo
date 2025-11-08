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

// proto covert
// https://github.com/envoyproxy/protoc-gen-validate/blob/main/validate/validate.proto

///
attribute min_length: Int64

///
attribute max_length: Int64

/// make sure the count of the container type's values is equal to the fixed_length
/// for String/Array/Map container type
///
/// fixed_length specifies that the String target must be the specified number of
/// characters (Unicode code points). Note that the number of
/// characters may different from the number of bytes in the String type.
attribute fixed_length: Int64

// min_bytes specifies that this String/Bytest target must be the specified number of bytes
// at a minimum
attribute min_bytes: Int64

// max_bytes specifies that this String/Bytest target must be the specified number of bytes
// at a maximum
attribute max_bytes: Int64

// fixed_bytes specifies that this String/Bytes target must be the specified number of bytes
attribute fixed_bytes: Int64

/// make sure the number is less then and equal to the maximun
/// (le <=)
@target(DeclType.type)
attribute maximum: Int64

///  (lt <)
attribute exclusive_maximum: Int64

///  (ge >=)
/// for Integer, Float, Duration, TimeStamp
attribute minimum: Int64

///  (gt >)
attribute exclusive_minmum: Int64

/// (ne !=)
attribute not_equal: Int64

attribute precision: Int32

attribute in: [Int64]

///
@target<Real>
attribute not_in: [Int64]

/// The multiple_of attribute can be used to limit numbers to multiples of a given number. It can be set to any positive number.
/// # example
/// ```
/// type Value: Int32 @multiple_of(10) // 0  OK, 10 OK, 20 OK, 23 NOT OK (not the multiple of 10)
/// ```
attribute multiple_of: UInt64


attribute unique: Bool

attribute nonempty: Bool


// Pattern specifes that this String target must match against the specified
// regular expression (RE2 syntax). The included expression should elide
// any delimiters.
attribute pattern: Regex

// prefix specifies that this String target must have the specified substring at
// the beginning of the string.
attribute prefix: String

// Suffix specifies that this String target must have the specified substring at
// the end of the string.
attribute suffix: String

// Contains specifies that this String target must have the specified substring
// anywhere in the string.
attribute contains: String

// NotContains specifies that this String target cannot have the specified substring
// anywhere in the string.
attribute not_contains: String

// Within specifies that this Timestamp must be within this duration of the
// current time.
attribute within: Duration