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

/// Stmt - Base type for all statements in mojo.
type Stmt {
    // statement start position
    start_position: Position @1 

    // statement end position
    end_position: Position @2

    /// Kind - The subclass of Stmt that this is.
    kind : Int @4

    /// Implicit - Whether this statement is implicit.
    implicit : Bool @5
}