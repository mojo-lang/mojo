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

type UpdateStmt: DataManipulationStmt {
    enum Type {
        update @0
        update_or_replace  @5
        update_or_rollback @6
        update_or_abort    @7
        update_or_fail     @8
        update_or_ignore   @9
    }

    with: WithClause @10

    type: Type @11
    table_name: QualifiedTableName @12

    set: SetClause @13
    from: FromClause @14
    where: WhereClause @15

    returning: ReturningClause @16
}