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

/// An object representing a Server Variable for server URL template substitution.
type ServerVariable {
    /// An enumeration of string values to be used if the substitution options are from a limited set.
    enum: [String] @1

    /// The default value to use for substitution, and to send, if an alternate value is *not* supplied.
    /// Unlike the `Schema` Object's `default`, this value *MUST* be provided by the consumer.
    default: String @2 @required

    /// An optional description for the server variable.
    /// CommonMark syntax MAY be used for rich text representation.
    description: Cached<document.Document> @4
}