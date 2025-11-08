
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

//@format('{left_bottom.longitude},{left_bottom.latitude},{right_top.longitude},{right_top.latitude}')
//@format('{left_bottom.longitude},{left_bottom.latitude},{left_bottom.altitude},{right_top.longitude},{right_top.latitude},{right_top.altitude}')
type BoundingBox {
    left_bottom: LngLat @1
    right_top: LngLat @2
}