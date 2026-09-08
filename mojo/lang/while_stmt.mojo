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

/// WhileStmt - while statement. After type-checking, the condition is of
/// type Builtin.Int1.
type WhileStmt : LoopStmt {
    condition: Expression @20

    /// A statement that is executed at least once and is re-executed each time the condition evaluates to true.
    /// equals to `repeat { statements } while( condition )`
    execute_at_least_once: Bool @21 
}
