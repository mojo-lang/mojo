// Copyright 2022 Mojo-lang.org
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

type Table = TableName @1 | TableFunctionName @2 | TableQuery @3

type TableName {
    schema_name: String @10

    name: String @13 @required
    alias: String @15
}

type TableFunctionName {
    schema_name: String @10

    name: String @13 @required
    arguments: [Expression] @14

    alias: String @15
}

type TableQuery {
    query: SelectStmt @10
    alias: String @15
}