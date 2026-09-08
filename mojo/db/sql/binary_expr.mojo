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

type BinaryExpr: lang.Expr {
    enum Operator {
        unspecified @0

        plus  @1
        minus @2
        mul   @3
        div   @4
        mod   @5

        lt        @10
        gt        @11
        lte       @12
        gte       @13
        equal     @14
        not_equal @15

        and     @20
        or      @21
        bit_and @25
        bit_or  @26
        bit_xor @27

        concat  @30

        extended @99
    }

    operator: Operator @10
    operator_symbol: String @11
    operator_precedence: Int32 @12

    left_argument: Expression @15
    right_argument: Expression @16
}