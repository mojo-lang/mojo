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

/// Describes the operations available on a single path.
/// A Path Item MAY be empty, due to ACL constraints.
/// The path itself is still exposed to the documentation viewer but they will not know which operations and parameters are available.
@example({})
type PathItem {
    /// Allows for an external definition of this path item.
    /// The referenced structure MUST be in the format of a Path Item Object.
    /// If there are conflicts between the referenced definition and this Path Item's definition,
    /// the behavior is undefined.
    ref: String @1 @alias('$ref')

    /// An optional, string summary, intended to apply to all operations in this path.
    summary: String @2

    /// An optional, string description, intended to apply to all operations in this path.
    description: Cached<document.Document> @3

    /// A definition of a GET operation on this path.
    get: Operation @5

    /// A definition of a PUT operation on this path.
    put: Operation @6

    /// A definition of a POST operation on this path.
    post: Operation @7

    /// A definition of a DELETE operation on this path.
    delete: Operation @8

    /// A definition of a OPTIONS operation on this path.
    options: Operation @9

    /// A definition of a HEAD operation on this path.
    head: Operation @10

    /// A definition of a PATCH operation on this path.
    patch: Operation @11

    /// A definition of a TRACE operation on this path.
    trace: Operation @12

    /// An alternative server array to service all operations in this path.
    servers: [Server] @13

    /// A list of parameters that are applicable for all the operations described under this path. These parameters can be overridden at the operation level, but cannot be removed there. The list MUST NOT include duplicated parameters. A unique parameter is defined by a combination of a name and location. The list can use the Reference Object to link to parameters that are defined at the OpenAPI Object's components/parameters.
    parameters: [Referenceable<Parameter>] @15
}