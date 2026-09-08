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

type ColumnConstraint {
    type PrimaryKey {
        order: Order @1
        conflict: ConflictClause @2
        auto_increment: Bool @3
    }

    type NotNull: ConflictClause {
    }

    type Unique: ConflictClause {
    }

    type Default {
        expression: Expression @2
    }

    name: String @1
    constraint: PrimaryKey @2
              | NotNull @3
              | Unique @4
              | CheckClause @5
              | Default @6
              | CollateClause @7
              | ForeignKeyClause @8
}