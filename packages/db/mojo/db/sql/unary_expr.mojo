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

type UnaryExpr: lang.Expr {
    enum Operator {
        unspecified @0
        minus    @1
        plus     @2
        tilde    @3
        not      @4
        bit_not  @5

        extended @99
    }

    operator: Operator @10
    operator_symbol: String @11
    argument: Expression @12
}
