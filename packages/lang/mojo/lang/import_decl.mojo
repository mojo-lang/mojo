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

type ImportDecl : Decl {
    /// 
    filter: String @11

    ///
    attributes: [Attribute] @12

    /// imported package name
    import_package_name:  String @13

    /// alias for the imported package name
    import_package_alias: String @14

    /// for some language, just import file not package
    import_file_name: String @15

    /// imported identifiers
    identifiers: [Identifier] @16
}