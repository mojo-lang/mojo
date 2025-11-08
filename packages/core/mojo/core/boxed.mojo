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

/// Wrapper type for the Generic T type
///
/// 实例化名称规格， TValue  TValues  TValuesMap
type Boxed<T> {
    val: T @1 //< the boxed value
}

/// Associative Values for string key based map
type AssocValues<T> = {String: T}

/// Wrapper type for `Bool`
///
/// The JSON representation for `BoolValue` is JSON boolean.
type BoolValue = Boxed<Bool>

/// Wrapper type for `Int32`
///
/// The JSON representation for `Int32Value` is JSON number.
type Int32Value = Boxed<Int32>

/// Wrapper type for `UInt32`
///
/// The JSON representation for `UInt32Value` is JSON number.
type UInt32Value = Boxed<UInt32>

/// Wrapper type for `Int64`
///
/// The JSON representation for `Int64Value` is JSON number.
type Int64Value = Boxed<Int64>

/// Wrapper type for `UInt64`
///
/// The JSON representation for `UInt64Value` is JSON number.
type UInt64Value = Boxed<UInt64>

/// Wrapper type for `Float32`
///
/// The JSON representation for `FloatValue` is JSON number.
type Float32Value = Boxed<Float32>

/// Alias type for `Float64Value`
///
/// The JSON representation for `FloatValue` is JSON number.
type FloatValue = Float32Value

/// Wrapper type for `Double`
///
/// The JSON representation for `DoubleValue` is JSON number.
type Float64Value = Boxed<Float64>

/// Alias type for `Float64Value`
///
/// The JSON representation for `DoubleValue` is JSON number.
type DoubleValue = Float64Value

/// Wrapper type for `String`
///
/// The JSON representation for `StringValue` is JSON string.
type StringValue = Boxed<String>

/// Wrapper type for `Bytes`
///
/// The JSON representation for `BytesValue` is JSON string.
type BytesValue = Boxed<Bytes>

/// Wrapper type for `Array<Bool>`
///
/// The JSON representation for `BoolValues` is JSON boolean array.
type BoolValues: [Bool]

/// Wrapper type for `Array<Int32>`
///
/// The JSON representation for `Int32Values` is JSON number array.
type Int32Values: [Int32]

/// Wrapper type for `Array<UInt32>`
///
/// The JSON representation for `UInt32Values` is JSON number array.
type UInt32Values: [UInt32]

/// Wrapper type for `Array<Int64>`
///
/// The JSON representation for `Int64s` is JSON number array.
type Int64Values: [Int64]

/// Alias type for `Int64Values`
///
/// The JSON representation for `IntValues` is JSON number array.
type IntValues = Int64Values

/// Wrapper type for `Array<UInt64>`
///
/// The JSON representation for `UInt64Values` is JSON number array.
type UInt64Values: [UInt64]

/// Alias type for `UInt64Values`
///
/// The JSON representation for `UIntValues` is JSON number array.
type UIntValues = UInt64Values

/// Wrapper type for `Array<Float32>`
///
/// The JSON representation for `Float32Values` is JSON number array.
type Float32Values: [Float32]

/// Alias type for `Float32Values`
///
/// The JSON representation for `FloatValues` is JSON number array.
type FloatValues = Float32Values

/// Wrapper type for `Array<Float64>`
///
/// The JSON representation for `Float64Values` is JSON number array.
type Float64Values: [Float64]

/// Alias type for `Float64Values`
///
/// The JSON representation for `DoubleValues` is JSON number array.
type DoubleValues = Float64Values

/// Wrapper type for `Array<String>`
///
/// The JSON representation for `StringValues` is JSON string array.
type StringValues: [String]

/// Wrapper type for `Map<String, String>`
///
/// The JSON representation for `StringMap` is JSON object.
type StringMap: {String: String}

/// Wrapper type for `Map<String, [String]>`
///
/// The JSON representation for `StringValues` is JSON object.
type StringMultiMap: {String: [String]}
