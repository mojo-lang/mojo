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

/// This is the declaration of an enum.
///
/// For example:
///
/// \code
/// \endcode
///
/// The type of the decl itself is a MetatypeType; use getDeclaredType()
/// to get the declared type ("Bool" or "Optional" in the above example).
///
/// \brief This represents a single case of an 'enum' declaration.
///
/// For example, the X, Y, and Z in this enum:
///
/// \code
///   enum V: Int {
///     x, y, z
///   }
/// \endcode
///
/// The type of an EnumElementDecl is always the EnumType for the containing
/// enum. EnumElementDecls are represented in the AST as members of their
/// parent EnumDecl, although syntactically they are subordinate to the
/// EnumCaseDecl.
//class EnumElementDecl : public ValueDecl {
//}

///
type EnumDecl : TypeDecl @ignore_fields("generic_parameters") {
    ///
    type: EnumType @20

    scope:   Scope @30
}
