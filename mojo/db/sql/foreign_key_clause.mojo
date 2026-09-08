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

type ForeignKeyClause: Clause {
    enum MatchType {
        unspecified @0

        /// if at least one referencing column is null, then the row of the referencing table passes the constraint check.
        /// If all referencing columns are not null, then the row passes the constraint check if and only if there is a
        /// row of the referenced table that matches all the referencing columns.
        simple @1

        /// if all referencing columns are null, then the row of the referencing table passes the constraint check.
        /// If at least one referencing columns is not null, then the row passes the constraint check if and only if
        /// there is a row of the referenced table that matches all the non-null referencing columns.
        partial @2

        /// if all referencing columns are null, then the row of the referencing table passes the constraint check.
        /// If all referencing columns are not null, then the row passes the constraint check if and only if
        /// there is a row of the referenced table that matches all the referencing columns.
        /// If some referencing column is null and another referencing column is non-null,
        /// then the row of the referencing table violates the constraint check.
        full @3
    }

    type TriggerAction {
        enum Trigger {
            updated @0
            deleted @1
        }
        enum Action {
            cascade @0
            set_null @1
            set_default @2
            restrict @3
            no_action @4
        }

        trigger: Trigger @1
        action: Action @2
    }

    table: String @10
    columns: [String] @11

    match: MatchType @12
    action: TriggerAction @13

    defer_strategy: DeferStrategy @15
}