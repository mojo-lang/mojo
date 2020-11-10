# Mojo Database Libaray

[TOC]

## Types

### LatLng

```
struct LngLat {
	lat: Real,
	lng: Real,
	alt: Real
}

LngLat(Real...);
```



```
latlng = Latlng();
latlng = LatLng(lat:23.33, lng: 123.33);
```

### Geometry

```
struct Geometry {
	dimension: Integer,
	coordinateDimension: Integer,
	spatialDimension: Integer,
	empty: Bool,
	simple: Bool
}
```
### Feature

```
Feature {
		geom: Geometry
}
```
### Point

```
Point : Geometry {

}

Points = Sequence<Poi>

Point(Cooridate, ...) => Either<Point, Points>

```
### Line
```
Line:Geometry {
}
Lines = Sequence<Line>

make_line(Sequence<Cooridate>, ...) : Either<Line, Lines>
```

### Polygon
```
Polygon:Geometry {
}
Polygons = Sequence<Polygon>

make_polygon(Sequence<Cooridate>, ...) : Either<Polygon, Polygons>
```

## Functions

### core

#### equals

```
equals(Geometry, Geometry) : Bool
equals(Feature, Feature) : Bool
```

### 几何拓扑判断函数
```
disjoint(l:Geometry, r:Geometry) : Bool
touches(Geometry, Geometry) : Bool
overlaps(Geometry, Geometry) : Bool

contains(Geometry, Geometry) : Bool
inside(Geometry, Geometry) : Bool
within(Geometry, Geometry) : Bool

crosses(Geometry, Geometry) : Bool
intersects(Geometry, Geometry) : Bool
```

### 几何次序判断函数

```
left(Geometry, Geometry) : Bool
right(Geometry, Geometry) : Bool

north(Geometry, Geometry) : Bool
south(Geometry, Geometry) : Bool

over(Geometry, Geometry) : Bool
under(Geometry, Geometry) : Bool
```

###

### 几何计算函数

#### distance

```
distance(geom1: Geometry, geom2: Geometry) : Real
```

Returns the shortest distance in units between any two Points in the two geometric
objects as calculated in the spatial reference system of geom1.

#### manhaton_distace()
#### road_distance()

#### buffer
```
buffer(geom: Geometry, radius: double): Geometry
```

This function returns a geometric object that represents all Points whose distance from
geom1 is less than or equal to the radius measured in units. Calculations are in the
spatial reference system of geom1.

#### convex_hull
```
convexHull (geom1: ogc:geomLiteral): ogc:geomLiteral
```

This function returns a geometric object that represents all Points in the convex hull of
geom1. Calculations are in the spatial reference system of geom1

#### intersection
```
intersection(geom1: Geometry, geom2:Geometry) : Geometry
```
This function returns a geometric object that represents all Points in the intersection of
geom1 with geom2. Calculations are in the spatial reference system of geom1.

#### union
```
union(geom1: Geometry, geom2:Geometry) : Geometry
```
This function returns a geometric object that represents all Points in the union of geom1
with geom2. Calculations are in the spatial reference system of geom1.

#### difference
```
difference(geom1: Geometry, geom2:Geometry) : Geometry
```
This function returns a geometric object that represents all Points in the set difference of
geom1 with geom2. Calculations are in the spatial reference system of geom1

#### length
```
length(geometry) : number
```
#### area
```
area(geometry) : number
```
#### centroid
```
centroid(geometry) : geometry
```

#### envelop
```
envelop(g:Geometry) : Geometry
```

This function returns the minimum bounding box of geom1. Calculations are in the
spatial reference system of geom1.

#### boundary
```
envelop(g:Geometry) : Geometry
```
This function returns the closure of the boundary of geom1. Calculations are in the spatial
reference system of geom1.

### transform

#### as_text
```
as_text(g:Geometry) : String
```

#### as_binary
```
as_binary(g:Geometry) : String
```

#### travel

#### near
```
near(g:Geometry) : Features
```

#### path


#### mapmatch
```
mapmatch(p:Point) : Point
```

#### grid_clip
```
grid_clip(f:Feature, level:Integer) : Features
grid_clip(f:Feature, levels:Array<Integer>) : Features
```

### 几何逻辑判断函数
#### and
```
and(g1:Geometry, g2:Geometry) : Geometry
```

#### or

#### xor

#### no
