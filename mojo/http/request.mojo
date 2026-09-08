// Copyright 2022 Mojo-lang.org
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

/// A Request represents an HTTP request received by a server
/// or to be sent by a client.
type Request {
    /// method specifies the HTTP method (GET, POST, PUT, etc.).
	/// For client requests, an empty string means GET.
    method: Method @1

    /// url specifies either the URI being requested (for server
	/// requests) or the URL to access (for client requests).
    url: Url @2

    /// The protocol version for incoming server requests.
    version: Version @3

    /// Headers contains the request header fields either received
	/// by the server or to be sent by the client.
    headers: Headers @4

    /// body is the request's body, which ban be raw bytes or JSON object
    body: Value @5
}

type RequestOptions {
    timeout: Duration @1
    max_retries: Int32 @2
}
