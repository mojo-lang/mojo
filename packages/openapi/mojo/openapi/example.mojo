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

type Example {
    /// Short description for the example.
    summary: String @1

    /// Long description for the example.
    description: String @2

    /// Embedded literal example.
    /// The value field and externalValue field are mutually exclusive.
    /// To represent examples of media types that cannot naturally represented in JSON or YAML,
    /// use a string value to contain the example, escaping where necessary.
    value: Value @3

    /// A URL that points to the literal example.
    /// This provides the capability to reference examples that cannot easily be included in JSON or YAML documents.
    /// The value field and externalValue field are mutually exclusive.
    external_value: Url @4
}