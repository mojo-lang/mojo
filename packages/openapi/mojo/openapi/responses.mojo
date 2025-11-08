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

/// A container for the expected responses of an operation. The container maps a HTTP response code to the expected response.
///
/// The documentation is not necessarily expected to cover all possible HTTP response codes because they may not be known in advance. However, documentation is expected to cover a successful operation response and any known errors.
/// 
/// The default MAY be used as a default response object for all HTTP codes that are not covered individually by the specification.
///
/// The Responses Object MUST contain at least one response code, and it SHOULD be the response for a successful operation call.
type Responses : {String: Referenceable<Response> }

enum HttpStatusCode {
    default
    code_200
}
