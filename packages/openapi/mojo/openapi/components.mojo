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

/// Holds a set of reusable objects for different aspects of the OAS. All objects defined within the components object will have no effect on the API unless they are explicitly referenced from properties outside the components object.
type Components {
    type Key = String @format

    /// An object to hold reusable Schema Objects.
    schemas: {Key: Schema} @1

    /// An object to hold reusable Response Objects.
    responses: {Key: Response} @2

    /// An object to hold reusable Parameter Objects.
    parameters: {Key: Parameter} @3

    /// An object to hold reusable Example Objects.
    examples: {Key: Example} @4

    /// An object to hold reusable Request Body Objects.
    request_bodies: {Key: RequestBody} @5

    /// An object to hold reusable Header Objects.
    headers: {Key: Header} @6

    /// An object to hold reusable Security Scheme Objects.
    security_schemes: {Key: SecurityScheme} @7

    /// An object to hold reusable Link Objects.
    links: {Key: Link} @8

    /// An object to hold reusable Callback Objects.
    callbacks: {Key: Callback} @9

    /// An object to hold reusable Path Item Object.
    path_items: {Key: PathItem} @10
}
