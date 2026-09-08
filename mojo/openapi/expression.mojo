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

/// Runtime expressions allow defining values based on information that will only be available within the HTTP message in an actual API call. This mechanism is used by Link Objects and Callback Objects.
///
/// The runtime expression is defined by the following ABNF syntax
/// ```abnf
/// expression = ( "$url" / "$method" / "$statusCode" / "$request." source / "$response." source )
/// source = ( header-reference / query-reference / path-reference / body-reference )
/// header-reference = "header." token
/// query-reference = "query." name  
/// path-reference = "path." name
/// body-reference = "body" ["#" json-pointer ]
/// json-pointer    = *( "/" reference-token )
/// reference-token = *( unescaped / escaped )
/// unescaped       = %x00-2E / %x30-7D / %x7F-10FFFF
///    ; %x2F ('/') and %x7E ('~') are excluded from 'unescaped'
/// escaped         = "~" ( "0" / "1" )
///   ; representing '~' and '/', respectively
/// name = *( CHAR )
/// token = 1*tchar
/// tchar = "!" / "#" / "$" / "%" / "&" / "'" / "*" / "+" / "-" / "." /
///   "^" / "_" / "`" / "|" / "~" / DIGIT / ALPHA
/// ```
///
/// Here, json-pointer is taken from RFC 6901, char from RFC 7159 and token from RFC 7230.
///
/// The name identifier is case-sensitive, whereas token is not.
///
/// The table below provides examples of runtime expressions and examples of their use in a value:
///
/// ##### Examples
/// 
/// | Source Location       | example expression         | notes                                                                                                                                               |
/// | --------------------- | -------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------- |
/// | HTTP Method           | `$method`                  | The allowable values for the `$method` will be those for the HTTP operation.                                                                        |
/// | Requested media type  | `$request.header.accept`   |                                                                                                                                                     |
/// | Request parameter     | `$request.path.id`         | Request parameters MUST be declared in the `parameters` section of the parent operation or they cannot be evaluated. This includes request headers. |
/// | Request body property | `$request.body#/user/uuid` | In operations which accept payloads, references may be made to portions of the `requestBody` or the entire body.                                    |
/// | Request URL           | `$url`                     |                                                                                                                                                     |
/// | Response value        | `$response.body#/status`   | In operations which return payloads, references may be made to portions of the response body or the entire body.                                    |
/// | Response header       | `$response.header.Server`  | Single header values only are available                                                                                                             |


@format("${type}{.source}")
type Expression {
    enum Type {
        unspecified @0

        url @1
        method @2
        status_code @3 @alias("statusCode")
        request @4
        source @5
        response @6
    }

    enum Location {
        unspecified @0
        path        @1
        query       @2
        header      @3
        body        @5
    }

    @format("{location}{.name}{#path}")
    type Source {
        location: Location @1
        name: String @2
        path: String @3
    }

    type: Type @1
    source: Source @2
}

