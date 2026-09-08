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
//

type Window {
    type Frame {
        type Bound {
            expression: Expression @1  // expr, current_row, unbounded

            preceding: Bool @5
            following: Bool @6
        }

        type: String @1
        lower_bound: Bound @2
        upper_bound: Bound @3

        exclude: Bool @5
        exclude_no_others: Bool @6
        current_row: Bool @7
        group: Bool @8
        ties: Bool @9
    }

    name: String @10

    partition: PartitionClause @13
    order_by: OrderByClause @14
    frame: Frame @15
}
