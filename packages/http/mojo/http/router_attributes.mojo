// Copyright 2022 Mojo-lang.org
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

/// the base path attribute for the method in interface
@target(DeclType.function, DeclType.interface)
attribute base_path: String

/// the get router attribute for the method in interface
@target(DeclType.function)
attribute get: Router

/// the post router attribute for the method in interface
@target(DeclType.function)
attribute post: Router

/// the put router attribute for the method in interface
@target(DeclType.function)
attribute put: Router

/// the patch router attribute for the method in interface
@target(DeclType.function)
attribute patch: Router

/// the delete router attribute for the method in interface
@target(DeclType.function)
attribute delete: Router

/// the head router attribute for the method in interface
@target(DeclType.function)
attribute head: Router

/// the trace router attribute for the method in interface
@target(DeclType.function)
attribute trace: Router

/// the options router attribute for the method in interface
@target(DeclType.function)
attribute options: Router

/// will replace the `{{key}}` block in the router, specially in the `path` field
@target(DeclType.function, DeclType.interface)
attribute router_template_values: Object
