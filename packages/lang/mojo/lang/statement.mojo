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

/// There are 3 main classes of nodes: Expressions nodes,
/// statement nodes, and declaration nodes. The node names usually
/// match the corresponding Go spec production names to which they
/// correspond. The node fields correspond to the individual parts
/// of the respective productions.
///
/// All nodes contain position information marking the beginning of
/// the corresponding source text segment; it is accessible via the
/// Pos accessor method. Nodes may contain additional position info
/// for language constructs where comments may be found between parts
/// of the construct (typically any larger, parenthesized subpart).
/// That position information is needed to properly position comments
/// when printing the construct.

///
///
@discriminator('@type')
@label_format('{}')
type Statement = ReturnStmt   @1 |
                 BreakStmt    @2 |
                 ContinueStmt @3 |
                 MatchStmt    @4 |
                 IfStmt       @5 |
                 ForStmt      @6 |
                 WhileStmt    @7 |
                 RepeatStmt   @8 |
                 BlockStmt    @9 |
                 Declaration  @10|
                 Expression   @11
