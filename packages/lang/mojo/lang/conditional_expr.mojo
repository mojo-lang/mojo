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

/// The conditional expression 'x ? y : z'.
///
/// Besides false, possible falsy expressions are: null, NaN, 0, the empty string (""), and undefined.
/// If condition is any of these, the result of the conditional expression will be the result of executing
/// the expression expr_if_false.
type ConditionalExpr : Expr {
    /// an expression whose value is used as a condition
    condition: Expression @10

    /// An expression which is evaluated if the condition evaluates to a truthy value
    /// (one which equals or can be converted to true).
    then_branch: Expression @11

    /// An expression which is executed if the condition is falsy
    /// (that is, has a value which can be converted to false).
    else_branch: Expression @12
}
