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

/// TupleExpr - Parenthesized expressions like `(a: x+x)` and '(x, y, 4)'.  Also
/// used to represent the operands to a binary operator.  Note that
/// expressions like '(4)' are represented with a ParenExpr.
type TupleExpr : Expr {
    /// Whether this tuple has any labels.
    has_element_labels : Bool @10

    elements: [Argument] @20
}
