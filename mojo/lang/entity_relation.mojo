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

enum RelationType {
    unspecified @0

    o2o @15
    o2o_two_types @1
    o2o_same_type @2
    o2o_bidirectional @3

    o2m @240 // 0xf << 4
    o2m_two_types @16
    o2m_same_type @32

    m2m @3840 // 0xf << 8
    m2m_two_types @256
    m2m_same_type @512
    m2m_bidirectional @768
}

type EntityRelation {
    /// '{from-entity-node-name}-{to-entity-node-name}' in the non-inverse edge
    name: String @1 @key
    type: RelationType @2

    from: EntityNode @5 @reference
    to: EntityNode @6 @reference

    edges: [EntityEdge] @10 @max_length(2)
}

type EntityRelations = [EntityRelation]
