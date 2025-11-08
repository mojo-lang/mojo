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


/// Case Style types for name case style changing
/// 
/// for value, snake_case is the defualt style
/// for type, CamelCase is the default style
enum CaseStyle {
    unspecified @0
    snake @1           //< snake_case style
    screaming_snake @2 //< SCREAMING_SNAKE_CASE style
    kebab @3           //< kebab-case style
    screaming_kebab @4 //< SCREAMING-KEBAB-CASE style
    camel @5           //< CamelCase style
    lower_camel @6     //< lowerCamelCase style
}
