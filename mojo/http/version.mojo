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

// The http version
// For client requests, these fields are ignored. The HTTP
// client code always uses either HTTP/1.1 or HTTP/2.
// See the docs on Transport for details.
@format("HTTP/{major}{.minor}")
type Version {
    major: Int32 @1
    minor: Int32 @2
}