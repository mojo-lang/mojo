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

/// Describes a single API operation on a path.
type Operation {
    /// A list of tags for API documentation control. Tags can be used for logical grouping of operations by resources or any other qualifier.
    tags: [String] @1

    /// A short summary of what the operation does.
    /// For maximum readability in the swagger-ui, this field SHOULD be less than 120 characters.
    summary: String @2 @max_length(120)

    /// A verbose explanation of the operation behavior. GFM syntax can be used for rich text representation.
    description: Cached<document.Document> @3

    /// Additional external documentation for this operation.
    external_docs: ExternalDocument @4

    /// Unique string used to identify the operation.
    /// The id MUST be unique among all operations described in the API.
    /// Tools and libraries MAY use the operationId to uniquely identify an operation,
    /// therefore, it is recommended to follow common programming naming conventions.
    operation_id: String @5

    /// A list of parameters that are applicable for this operation. If a parameter is already defined at the Path Item, the new definition will override it but can never remove it. The list MUST NOT include duplicated parameters. A unique parameter is defined by a combination of a name and location. The list can use the Reference Object to link to parameters that are defined at the OpenAPI Object's components/parameters.
    parameters: [Referenceable<Parameter>] @10

    /// The request body applicable for this operation. The requestBody is only supported in HTTP methods where the HTTP 1.1 specification RFC7231 has explicitly defined semantics for request bodies. In other cases where the HTTP spec is vague, requestBody SHALL be ignored by consumers.
    request_body: Referenceable<RequestBody> @11

    /// The list of possible responses as they are returned from executing this operation.
    responses: Responses @12 @required

    /// A map of possible out-of band callbacks related to the parent operation.
    /// The key is a unique identifier for the Callback Object.
    /// Each value in the map is a Callback Object that describes a request
    /// that may be initiated by the API provider and the expected responses.
    /// The key value used to identify the callback object is an expression,
    /// evaluated at runtime, that identifies a URL to use for the callback operation.
    callbacks: {String: Referenceable<Callback>} @13

    /// Declares this operation to be deprecated.
    /// Consumers SHOULD refrain from usage of the declared operation.
    deprecated: Bool @15

    /// A declaration of which security mechanisms can be used for this operation.
    /// The list of values includes alternative security requirement objects that can be used.
    /// Only one of the security requirement objects need to be satisfied to authorize a request.
    /// This definition overrides any declared top-level security.
    /// To remove a top-level security declaration, an empty array can be used.
    security: SecurityRequirement @16

    /// An alternative server array to service this operation.
    /// If an alternative server object is specified at the Path Item Object or Root level,
    /// it will be overridden by this value.
    servers : [Server] @17
}