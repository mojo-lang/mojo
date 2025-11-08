// Copyright 2022 Mojo-lang.org
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

@case_style("screaming_snake")
enum Method {
    unsepecified @0

    /// The HTTP GET method requests a representation of the specified resource.
    ///
    /// Requests using GET should only be used to request data (they shouldn't include data).
    get @1

    post @2

    put @3

    patch @4

    delete @5

    options @6

    head @7

    trace @8

    connect @9
}
