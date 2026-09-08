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

/// StructDecl - This is the declaration of a struct, for example:
///
/// ``` mojo
/// struct Complex { r : Double, i : Double }
/// ```
///
/// The type of the decl itself is a MetaType; use `Complex.type`
/// to get the `Complex`'s declared type.
///
type StructDecl : TypeDecl {
    ///
    type: StructType @20

    ///
    type_alias_decls: [TypeAliasDecl] @21

    ///
    enum_decls: [EnumDecl] @22

    ///
    struct_decls: [StructDecl] @23

    scope:       Scope @30
}