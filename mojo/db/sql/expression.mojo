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

@discriminator('@type')
@label_format('{}')
type Expression = NullLiteralExpr @1
                | BooleanLiteralExpr @2
                | IntegerLiteralExpr @3
                | FloatLiteralExpr @4
                | StringLiteralExpr @5
                | BinaryStringLiteralExpr @6
                | IdentifierExpr @7
                | BindParameterExpr @10
                | ColumnReferenceExpr @11
                | UnaryExpr @15
                | BinaryExpr @17
                | BetweenExpr @18
                | InExpr @19
                | CaseExpr @20
                | CastExpr @21
                | FunctionCallExpr @22
                | ExistsExpr @23
                | IsExpr @24
                | GlobExpr @25
                | LikeExpr @26
                | MatchExpr @27
                | RegexpExpr @28
                | NullExpr @30
                | RaiseFunctionExpr @31
                | SubqueryExpr @35
                | Expressions @40
                | ParenthesizedExpr @41

type Expressions = [Expression]
