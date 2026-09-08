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

/// 几何拓扑判断函数集
///
///
func disjoint(self:Geometry, geom:Geometry) -> Bool

///
///
///
func touches(self:Geometry, geom:Geometry) -> Bool

///
///
///
func overlaps(self:Geometry, geom:Geometry) -> Bool

///
///
///
func contains(self:Geometry, geom:Geometry) -> Bool

///
///
///
func inside(self:Geometry, geom:Geometry) -> Bool

///
///
///
func within(self:Geometry, geom:Geometry) -> Bool

///
///
///
func crosses(self:Geometry, geom:Geometry) -> Bool

///
///
///
func intersects(self:Geometry, geom:Geometry) -> Bool
