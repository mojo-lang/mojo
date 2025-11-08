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

/// Allows configuration of the supported OAuth Flows.
type OAuthFlows {
    /// Configuration for the OAuth Implicit flow
    implicit: OAuthFlow  @1 @ignore_fields('token_url')

    /// Configuration for the OAuth Resource Owner Password flow
    password: OAuthFlow  @2 @ignore_fields('authorization_url')

    /// Configuration for the OAuth Client Credentials flow.
    client_credentials: OAuthFlow @3 @ignore_fields('authorization_url')

    /// Configuration for the OAuth Authorization Code flow.
    authorization_code: OAuthFlow @4
}