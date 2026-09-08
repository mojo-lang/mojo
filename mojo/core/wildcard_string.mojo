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

///
/// *	Represents zero or more characters	                        bl* finds bl, black, blue, and blob
/// .   Represents a single character	                            h.t finds hot, hat, and hit
/// []	Represents any single character within the brackets	        h[oa]t finds hot and hat, but not hit
/// ^	Represents any character not in the brackets	            h[^oa]t finds hit, but not hot and hat
/// -	Represents any single character within the specified range	c[a-b]t finds cat and cbt
type WildcardString {
    type Wildcard {
        enum Type {
            unspecified @0
            single_character @1
            zero_or_more_characters @2
            in_characters @3
            not_in_characters @4
        }

        type: Type @1
        content: String @2
    }

    wildcards: [Wildcard] @1
}
