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

/// The format is: `{algorithm} {value}` and the algorithm should be one of the `MD5`, `SHA1`, `SHA256` and `SHA512`.
@format('{algorithm} {value}')
type Checksum {
    /// commonly used algorithm for Checksum
    @case_style("screaming_snake")
    enum Algorithm {
        unspecified @0

        md5 @1    //< md5 algorithm
        sha1 @2   //< sha1 algorithm
        sha256 @3 //< sha256 algorithm
        sha512 @4 //< sha512 algorithm
    }

    /// the algorithm of the checksum
    algorithm: Algorithm @1
    
    /// the checksum value
    val: String @2
}
