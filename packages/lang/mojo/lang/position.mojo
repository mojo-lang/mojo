// Copyright 2021 Mojo-lang.org
//
// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
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


/// Positions
///
/// Position describes an arbitrary source position
/// including the file, line, and column location.
/// A Position is valid if the line number is > 0.
///
type Position {
  	filename: String @1 //< filename, if any
  	offset:   Int    @2 //< offset, starting at 0 (byte count)
  	line:     Int    @3 //< line number, starting at 1
  	column:   Int    @4 //< column number, starting at 1 (byte count)

    /// the comment before the AST node
    leading_comments: [Comment] @10

    /// usually exist in the end position of the AST node
    /// and only one of the following situations:
    /// ```mojo
    ///    a = 5 // following line comment 
    /// ```
    tailing_comments: [Comment] @13
}

/// valid reports whether the position is valid.
func valid(pos: Position) -> Bool {
    return pos.line > 0
}

// String returns a string in one of several forms:
//
//	file:line:column    valid position with file name
//	line:column         valid position without file name
//	file                invalid position with file name
//	-                   invalid position without file name
//
func to<String>(pos: Position) -> String {
}

// Pos is a compact encoding of a source position within a file.
// It can be converted into a Position for a more convenient, but much
// larger, representation.
//
// The Pos value for a given file is a number in the range [base, base+size],
// where base and size are specified when adding the file to the file set via
// AddFile.
//
// To create the Pos value for a specific source offset (measured in bytes),
// first add the respective file to the current file set using FileSet.AddFile
// and then call File.Pos(offset) for that file. Given a Pos value p
// for a specific file set fset, the corresponding Position value is
// obtained by calling fset.Position(p).
//
// Pos values can be compared directly with the usual comparison operators:
// If two Pos values p and q are in the same file, comparing p and q is
// equivalent to comparing the respective source file offsets. If p and q
// are in different files, p < q is true if the file implied by p was added
// to the respective file set before the file implied by q.
//
type Pos = Int

// The zero value for Pos is NoPos; there is no file and line information
// associated with it, and NoPos.IsValid() is false. NoPos is always
// smaller than any other Pos value. The corresponding Position value
// for NoPos is the zero value for Position.
//
const no_pos = 0

// IsValid reports whether the position is valid.
func valid(p: Pos) -> Bool {
    return p != no_pos
}

type PosSpan = [Int] @fixed_length(2)

func start(span: PosSpan) -> Pos {
    return span[0];
}

func end(span: PosSpan) -> Pos {
    return span[0] + span[1];
}