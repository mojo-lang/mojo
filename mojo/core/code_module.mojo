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

/// Carries information about code modules that are loaded into a process.
/// See src/google_breakpad/processor/code_module.h
type CodeModule {
  /// The base address of this code module as it was loaded by the process.
  base_address: Int64 @1

  /// The size of the code module.
  size: Int64 @2

  /// The path or file name that the code module was loaded from.
  code_file: String @3

  /// An identifying string used to discriminate between multiple versions and
  /// builds of the same code module.  This may contain a uuid, timestamp,
  /// version number, or any combination of this or other information, in an
  /// implementation-defined format.
  code_identifier: String @4

  /// The filename containing debugging information associated with the code
  /// module.  If debugging information is stored in a file separate from the
  /// code module itself (as is the case when .pdb or .dSYM files are used),
  /// this will be different from code_file.  If debugging information is
  /// stored in the code module itself (possibly prior to stripping), this
  /// will be the same as code_file.
  debug_file: String @5

  /// An identifying string similar to code_identifier, but identifies a
  /// specific version and build of the associated debug file.  This may be
  /// the same as code_identifier when the debug_file and code_file are
  /// identical or when the same identifier is used to identify distinct
  /// debug and code files.
  debug_identifier: String @6

  /// A human-readable representation of the code module's version.
  version: String @7
}
