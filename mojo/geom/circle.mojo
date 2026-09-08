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

/// Circle Feature
///
///
type Circle {
    type: String @1 @required @const("Circle")

    /// the center of the circle
    center: LngLat @2 @required @alias("coordinates") @type_format<PointCoordinates>
    
    /// the radius of the circle
    radius: Float  @5 @exclusive_minmum(0)
}
