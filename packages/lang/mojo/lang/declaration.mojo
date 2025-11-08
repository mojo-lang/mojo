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
@discriminator('@type')
@label_format('{}')
type Declaration = PackageDecl        @1
                 | ImportDecl         @2
                 | EnumDecl           @3
                 | StructDecl         @4
                 | TypeAliasDecl      @5
                 | InterfaceDecl      @6
                 | ConstantDecl       @10
                 | VariableDecl       @11
                 | AttributeDecl      @12
                 | AttributeAliasDecl @13
                 | FunctionDecl       @14
                 | ConstructorDecl    @15
                 | GenericParameter   @19
                 | GroupDecl          @20
