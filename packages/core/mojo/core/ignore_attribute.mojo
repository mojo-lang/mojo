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

/// 
@target(DeclType.value)
attribute ignore: Bool

/// will not encode when the field is empty
@target(DeclType.value)
attribute ignore_empty: Bool @default(true)

/// ingore the types in the Union type when the declaration inherites the Union type.
@target(DeclType.type)
attribute ignore_types: [String]

/// ignore the fields in the struct type when the declaration inherites the struct type.
@target(DeclType.type)
attribute ignore_fields: FieldMask
