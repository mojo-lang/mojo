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

/// Object type
type Object : {String: Value}

/// Values type
type Values : [Value]

/// Value type
@disable_generate({"go": ["json"], "java": ["json"]})
type Value = Values @1 | Object @2 | Null @3 | Bool @4 | Positive @5 | Negative @6 | Double @7 | String @8 | Bytes @9

/// the kind of the Value
enum ValueKind {
    unspecified @0

    null    @1
    boolean @2
    integer @3
    number  @4
    string  @5
    bytes   @6

    array   @10 //<  An ordered list of instances, from the JSON "array" production
    object  @11 //< An unordered set of properties mapping a string to an instance, from the JSON "object" production
}
