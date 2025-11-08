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

/// \brief Represents a runtime type check query, 'a is T', where 'T' is a type
/// and 'a' is a value of some related type. Evaluates to a Bool true if 'a' is
/// of the type and 'a as T' would succeed, false otherwise.
///
/// FIXME: We should support type queries with a runtime metatype value too.
type IsExpr : CheckedCastExpr 


type CheckedCastExpr : Expr