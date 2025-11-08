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

/// GeoJSON Feature
///
///
type Feature {
    type: String @1 @required @const("Feature")

    id: Id @2
    geometry: Geometry @3

    bbox: BoundingBox @4
    spatial_reference: SpatialReference @5 @alias("crs")

    //
    properties: {String: Value} @6 @protobuf.encoding({type: 'map', field_numbers: [13, 14, 15]})

    //keys: [string] @13  // global arrays of unique keys
    //values: [Value] @14 // unique values
    //encoded_properties: [UInt] @15
}

//func encode_properties(self: Feature)

//func encode_properties(self: Feature, keys: {String: UInt}, values: {Value: Uint})

//func decode_properties(self: Feature)