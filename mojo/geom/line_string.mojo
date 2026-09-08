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

/// Geojson LineString object
///
/// # example
/// ``` json
/// {
///     "type": "LineString",
///     "coordinates": [
///         [102.0, 0.0],
///         [103.0, 1.0],
///         [104.0, 0.0],
///         [105.0, 1.0]
///     ]
///  }
/// ```
type LineString {
    type: String @1 @required @const("LineString")
    
    coordinates: [LngLat] @5 @required @type_format<LineStringCoordinates>
}

@disable_generate
type LineStringCoordinates = [PointCoordinates] @min_length(2)

///
///
///
//LineString(points: LngLat...)
