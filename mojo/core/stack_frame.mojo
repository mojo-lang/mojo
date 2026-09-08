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

/// Represents a single frame in a stack
/// See src/google_breakpad/processor/stack_frame.h
type StackFrame {
  /// Indicates how well the instruction pointer derived during
  /// stack walking is trusted. Since the stack walker can resort to
  /// stack scanning, it can wind up with dubious frames.
  /// In rough order of "trust metric".
  enum Trust {
    unspecified @0
    scan @1      //< Scanned the stack, found this
    cfi_scan @2  //< Found while scanning stack using call frame info
    fp @3        //< Derived from frame pointer
    cfi @4       //< Derived from call frame info
    prewalked @5 //< Explicitly provided by some external stack walker.
    context @6   //< Given as instruction pointer in a context
  }

  /// The program counter location as an absolute virtual address.  For the
  /// innermost called frame in a stack, this will be an exact program counter
  /// or instruction pointer value.  For all other frames, this will be within
  /// the instruction that caused execution to branch to a called function,
  /// but may not necessarily point to the exact beginning of that instruction.
  instruction: Int64 @1 @required

  /// The module in which the instruction resides.
  module: CodeModule @2

  /// The function name, may be omitted if debug symbols are not available.
  function_name: String @3

  /// The start address of the function, may be omitted if debug symbols
  /// are not available.
  function_base: Int64 @4

  /// The source file name, may be omitted if debug symbols are not available.
  source_file_name: String @5

  /// The (1-based) source line number, may be omitted if debug symbols are
  /// not available.
  source_line: Int32 @6

  /// The start address of the source line, may be omitted if debug symbols
  /// are not available.
  source_line_base: Int64 @7

  /// Amount of trust the stack walker has in the instruction pointer
  /// of this frame.
  trust: Trust @8
}
