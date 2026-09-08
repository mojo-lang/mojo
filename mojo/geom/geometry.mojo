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

/// GeoJSON Geometry
///
/// A Geometry object represents points, curves, and surfaces in
/// coordinate space.  Every Geometry object is a GeoJSON object no
/// matter where it occurs in a GeoJSON text.
///
/// * The value of a Geometry object's "type" member MUST be one of the
///   seven geometry types: "Point", "MultiPoint", "LineString",
///   "MultiLineString", "Polygon", "MultiPolygon", and "GeometryCollection".
/// * A GeoJSON Geometry object of any type other than
///   "GeometryCollection" has a member with the name "coordinates".
///   The value of the "coordinates" member is an array.  The structure
///   of the elements in this array is determined by the type of
///   geometry.  GeoJSON processors MAY interpret Geometry objects with
///   empty "coordinates" arrays as null objects.      
@label_format('{}')
@discriminator("type")
type Geometry = Point              @1
              | MultiPoint         @2
              | LineString         @3
              | MultiLineString    @4
              | Polygon            @5
              | MultiPolygon       @6
              | GeometryCollection @7
