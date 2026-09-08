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

/// 经纬度
@format('{longitude},{latitude}{,altitude}')
type LngLat {
    /// the longitude of the `LngLat`
    longitude: Double @1 @required //< 经度
    
    /// the latitude of the `LngLat`
    latitude: Double @2 @required //< 纬度
    
    /// the altitude of the `LngLat` in meters.
    altitude: Double @3 //< 高度
}

///
/// parse the format lik `'\{longitude},\{latitude}'` to LngLat object
///
//LngLat(String)

/// convert the double array to LngLat object
/// #notes
///  * the double array size must equals to `2` or `3`
///
//LngLat(Double... @max_size(3) @min_size(2))

///
//LngLat(values: [Double] @max_size(3) @min_size(2)) {
//    LngLat(values...);
//}