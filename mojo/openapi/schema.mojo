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

/// The Schema Object allows the definition of input and output data types. These types can be objects, but also primitives and arrays. This object is an extended subset of the JSON Schema Specification Wright Draft 00.
///
/// For more information about the properties, see JSON Schema Core and JSON Schema Validation. Unless stated otherwise, the property definitions follow the JSON Schema.
type Schema {
    enum Type {
        unspecified @0
        null    @1
        boolean @2
        integer @3
        number  @4
        string  @5
        
        array   @10 //<  An ordered list of instances, from the JSON "array" production
        object  @11 //< An unordered set of properties mapping a string to an instance, from the JSON "object" production
    }

    /// Adds support for polymorphism. The discriminator is an object name that is used to differentiate between other schemas which may satisfy the payload description. See Composition and Inheritance for more details.
    discriminator: Discriminator @1

    /// This MAY be used only on properties schemas. It has no effect on root schemas. Adds additional metadata to describe the XML representation of this property.
    xml: Xml @2

    /// Additional external documentation for this schema.
    external_docs: ExternalDocument @3

    /// A free-form property to include an example of an instance for this schema. To represent examples that cannot be naturally represented in JSON or YAML, a string value can be used to contain the example with escaping where necessary.
    @deprecated("The example property has been deprecated in favor of the JSON Schema examples keyword. Use of example is discouraged, and later versions of this specification may remove it.")
    example: Value @4

    // The following properties are taken directly from the JSON Schema definition and follow the same specifications
    //{
        title: String @10

        description: Cached<document.Document> @11 //< support GFM

        /// Value MUST be a string. Multiple types via an array are not supported.
        type: Type @12

        /// format is an open string-valued property,
        /// and can have any value. Formats such as "Email", "Uuid", and so on,
        /// MAY be used even though undefined by this specification.
        /// Types that are not accompanied by a format property follow the type definition in the JSON Schema.
        /// Tools that do not recognize a specific format MAY default back to the type alone, as if the format is not specified.
        format: String @13

        /// The value of "multipleOf" MUST be a number, strictly greater than 0.
        ///
        /// A numeric instance is valid only if division by this keyword's value results in an integer.
        ///
        /// The `multiple_of` will be less than 1, when the numeric has a fixed decimals.
        /// For example, `multiple_of` will be `0.01`, when the numeric need to hold two decimals, and any greater precision will be rejected.
        multiple_of: Double @14 @exclusive_maximum(0)

        maximum: Value @15
        exclusive_maximum: Value @16

        minimum: Value @17
        exclusive_minimum: Value @18

        max_length: UInt @19

        /// A string instance is valid against this keyword if its length is greater than, or equal to, the value of this keyword.
        /// The length of a string instance is defined as the number of its characters as defined by RFC 7159.
        /// Omitting this keyword has the same behavior as a value of 0.
        min_length: UInt @20

        pattern: Regex @21

        /// Value MUST be an object and not an array. Inline or referenced schema MUST be of a Schema Object and not a standard JSON Schema. items MUST be present if the type is array.
        items: Referenceable<Schema> @22

        max_items: UInt @23
        
        min_items: UInt @24 @default(0)
        
        unique_items: Bool @25 @default(false)

        properties: {String: Referenceable<Schema>} @26

        additional_properties: Referenceable<Schema> @27

        max_properties: UInt @28

        min_properties: UInt @29 @default(0)
        
        required: [String] @30 @unique

        enum: [Value] @31 @unique @nonempty
    
        /// Inline or referenced schema MUST be of a Schema Object and not a standard JSON Schema.
        /// validates the value against all the sub-schemas
        all_of: [Referenceable<Schema>] @35 @nonempty

        /// validates the value against exactly one of the sub-schemas
        one_of: [Referenceable<Schema>] @36

        /// validates the value against any (one or more) of the sub-schemas
        any_of: [Referenceable<Schema>] @37

        /// make sure the value is not valid against the specified schema.
        not: Referenceable<Schema> @38

        default: Value @40
    //}

    /// Allows sending a null value for the defined schema. Default value is false.
    nullable: Bool @50 = false

    /// Relevant only for Schema "properties" definitions. Declares the property as "read only". This means that it MAY be sent as part of a response but SHOULD NOT be sent as part of the request. If the property is marked as readOnly being true and is in the required list, the required will take effect on the response only. A property MUST NOT be marked as both readOnly and writeOnly being true. Default value is false.
    read_only: Bool @51

    /// Relevant only for Schema "properties" definitions. Declares the property as "write only". Therefore, it MAY be sent as part of a request but SHOULD NOT be sent as part of the response. If the property is marked as writeOnly being true and is in the required list, the required will take effect on the request only. A property MUST NOT be marked as both readOnly and writeOnly being true. Default value is false.
    write_only: Bool @52

    /// Specifies that a schema is deprecated and SHOULD be transitioned out of usage. Default value is false.
    deprecated: Bool @53 = false
}