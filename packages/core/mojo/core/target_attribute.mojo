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


/// specificify the attribute, which can apply to the declaration types
///
/// #example
/// ```
/// @target(DeclType.value)
/// attribute alias: String
/// ```
/// this means that the alias attribute only can apply to the value declaration.
attribute target: [DeclType]

/// the declaration type of the mojo language.
enum DeclType {
    unspecified @0
    type @1               //< struct, interface (including annotation type), or enum declaration
    value @2              //< Field declaration (includes enum constants)
    function  @3          //< Method declaration
    constructor @4        //< Constructor declaration
    attribute @5          //< Annotation type declaration
    package @6            //< Package declaration
    generic @7            //< Generic Type parameter declaration
}