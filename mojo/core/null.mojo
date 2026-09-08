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
/// a unit type is a type that allows only one value (and thus can hold no information). The carrier (underlying set) associated with a unit type can be any singleton set. There is an isomorphism between any two such sets, so it is customary to talk about the unit type and ignore the details of its value. One may also regard the unit type as the type of 0-tuples, i.e. the product of no types.
///
/// the unit type is called None or () and its only value is `none`, reflecting the 0-tuple interpretation.
///
type Null