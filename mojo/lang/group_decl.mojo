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
///
type GroupDecl : Decl {
    /// come from `group_decl_name` attribute, otherwise will be the index from "0"
    name: String @10

    ///
    attributes: [Attribute] @12

    /// may be the `struct`, `enum`, `constant`, `variable`, `attribute` `function`
    type: Identifier.Kind @20

    ///
    declarations: [Declaration @reference] @21
}
