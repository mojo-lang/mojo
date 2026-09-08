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

/// A media type (also known as a Multipurpose Internet Mail Extensions or MIME type) is a standard
/// that indicates the nature and format of a document, file, or assortment of bytes.
/// It is defined and standardized in IETF's [RFC 6838](https://tools.ietf.org/html/rfc6838).
///
/// The simplest MIME type consists of a type and a subtype;
/// these are each strings which, when concatenated with a slash (/) between them, comprise a MIME type. 
/// No whitespace is allowed in a MIME type:
@format('{type}/{subtype}{;{parameter.key}={parameter.value}}')
type MediaType {
    type Parameter {
        key:String @1
        value: Value @2 @format<String>
    }

    type: String @1
    subtype: String @2

    parameter: Parameter @3
}

//
// construct MediaType from string using the template
//
//MediaType(str:String)
