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
///
type StructType {
    /// position of the character "{" belonging to the StructType
    start_position: Position @1

    /// position of the character "}"
    end_position:   Position @2

    /// A Field represents a Field declaration list in a struct type,
    /// a method list in an interface type, or a parameter/result declaration
    /// in a signature.
    fields:   [ValueDecl] @10

    ///
    inherits: [NominalType] @11

    ///
    groups:   [GroupDecl] @12

    /// 
    inherit_position: Position @19
}
