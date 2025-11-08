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

/// retain the types in the Union type when the declaration inherites the Union type, other types will be ignored.
/// #see
///  - `ignore_types`
@target(DeclType.type)
attribute retain_types: [String]

/// only retain the fields in the struct type when the declaration inherites the struct type, other fields will be ingored
/// #see
///  - `ignore_fields`
@target(DeclType.type)
attribute retain_fields: FieldMask

/// alias to the retain_types
@target(DeclType.type)
attribute types: [String]

/// alias to the retain_fields
@target(DeclType.type)
attribute fields: FieldMask