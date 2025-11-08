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

/// integer percentage
///
///
/// @format("{}%")
type Percentage: Int32 //@range(0..100)

/// float percentage
/// the value is from 0 to 100 with two decimal places precision normally, and then printed to '100.00%', '12.45%'
/// @format("{}%")
//type FloatPercentage: Float32 @precision(4)
