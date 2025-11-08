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


/// Status Code and Reason Phrase
/// The code element is a 3-digit integer result code of the attempt to understand and satisfy the request.
/// The reason phrase is intended to give a short textual description of the code.
@format
type Status {
    code: Int32 @1
    reason: String @2
}
