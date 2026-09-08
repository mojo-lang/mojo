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

/// ConstructorDecl - Declares a constructor for a type.  For example:
///
/// ```
/// type X {
///     x: Int
/// }
///
/// X(i : Int) {
///    x = i
/// }
/// ```
/// type X<T> {
///	    x : T
/// }
///
/// X<Int @argument>(i : Int)
///
/// X<T>(i: T) {
/// }
type ConstructorDecl : Decl {
    ///
    name: String @10

    ///
    full_name: String @11

    ///
    attributes: [Attribute] @12

    ///
    generic_parameters: [GenericParameter] @13

    /// 
    name_position: Position @19 @ignore

    ///
    signature: FunctionSignature @20

    ///
    body: BlockStmt @21

    ///
    scope: Scope @30
}
