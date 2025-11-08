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

/// GeoJSON Polygon
///
///
type Polygon {
    /// 类型字段
    type: String @1 @required @const("Polygon")

    /// 包含的线段
    line_strings: [LineString] @5 @required @alias("coordinates") @type_format<PolygonCoordinates>
}

@disable_generate
type LinearRingCoordinates = [PointCoordinates] @min_length(4)

@disable_generate
type PolygonCoordinates = [LinearRingCoordinates]

///
///
///
//Polygon(points: LngLat...)