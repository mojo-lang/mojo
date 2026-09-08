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

/// Describes a single operation parameter.
///
/// A unique parameter is defined by a combination of a name and location.
///
/// #Parameter Locations
///
/// There are four possible parameter locations specified by the in field:
///
/// path - Used together with Path Templating, where the parameter value is actually part of the operation's URL. This does not include the host or base path of the API. For example, in /items/{itemId}, the path parameter is itemId.
/// query - Parameters that are appended to the URL. For example, in /items?id=###, the query parameter is id.
/// header - Custom headers that are expected as part of the request. Note that RFC7230 states header names are case insensitive.
/// cookie - Used to pass a specific cookie value to the API.
type Parameter {
    enum Location {
        unspecified @0
        path        @1
        query       @2
        header      @3
        cookie      @4
    }

    enum Style {
        unspecified     @0
        matrix          @1
        label           @2
        form            @3
        simple          @4
        space_delimited @5
        pipe_delimited  @6
        deep_object     @7
    }

    ///  The name of the parameter. Parameter names are case sensitive.
    ///
    /// * If in is "path", the name field MUST correspond to the associated path segment from the path field in the Paths Object. See Path Templating for further information.
    /// * If in is "header" and the name field is "Accept", "Content-Type" or "Authorization", the parameter definition SHALL be ignored.
    /// * For all other cases, the name corresponds to the parameter name used by the in property.
    name: String @1 @required

    ///  The location of the parameter.
    in: Location @2 @required

    /// A brief description of the parameter.
    /// This could contain examples of use. GFM syntax can be used for rich text representation.
    description: String @3

    /// Determines whether this parameter is mandatory.
    /// If the parameter is in "path", this property is required and its value MUST be true.
    /// Otherwise, the property MAY be included and its default value is false.
    required: Bool @4

    /// Specifies that a parameter is deprecated and SHOULD be transitioned out of usage.
    deprecated: Bool @5 @default(false)

    /// Sets the ability to pass empty-valued parameters. This is valid only for query parameters and allows sending a parameter with an empty value. Default value is false. If style is used, and if behavior is n/a (cannot be serialized), the value of allowEmptyValue SHALL be ignored.
    allow_empty_value : Bool @6

    // The rules for serialization of the parameter are specified in one of two ways. For simpler scenarios, a `schema` and `style` can describe the structure and syntax of the parameter.
    //{
        /// Describes how the parameter value will be serialized depending on the type of the parameter value. Default values (based on value of in): for query - form; for path - simple; for header - simple; for cookie - form.
        style: Style @10

        /// When this is true, parameter values of type array or object generate separate parameters for each value of the array or key-value pair of the map. For other types of parameters this property has no effect. When style is form, the default value is true. For all other styles, the default value is false.
        explode: Bool @11

        /// Determines whether the parameter value SHOULD allow reserved characters, as defined by RFC3986 :/?#[]@!$&'()*+,;= to be included without percent-encoding. This property only applies to parameters with an in value of query. The default value is false.
        allow_reserved: Bool @12

        /// The schema defining the type used for the parameter.
        schema: Referenceable<Schema> @13

        /// Example of the media type. The example SHOULD match the specified schema and encoding properties if present. The example object is mutually exclusive of the examples object. Furthermore, if referencing a schema which contains an example, the example value SHALL override the example provided by the schema. To represent examples of media types that cannot naturally be represented in JSON or YAML, a string value can contain the example with escaping where necessary.
        example: Value @14

        /// Examples of the media type. Each example SHOULD contain a value in the correct format as specified in the parameter encoding. The examples object is mutually exclusive of the example object. Furthermore, if referencing a schema which contains an example, the examples value SHALL override the example provided by the schema.
        examples: {String: Referenceable<Example>} @15
    //}

    // For more complex scenarios, the content property can define the media type and schema of the parameter. A parameter MUST contain either a schema property, or a content property, but not both. When example or examples are provided in conjunction with the schema object, the example MUST follow the prescribed serialization strategy for the parameter.
    //{
        /// A map containing the representations for the parameter. The key is the media type and the value describes it. The map MUST only contain one entry.
        content: {String: MediaType} @20
    //}
}