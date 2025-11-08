package geom

import (
    "math"
)

// hsin haversin(θ) function
func hsin(theta float64) float64 {
    return math.Pow(math.Sin(theta/2), 2)
}

// Distance function returns the distance (in meters) between two points of
//     a given longitude and latitude relatively accurately (using a spherical
//     approximation of the Earth) through the Haversin Distance Formula for
//     great arc distance on a sphere with accuracy for small distances
//
// point coordinates are supplied in degrees and converted into rad. in the func
//
// distance returned is METERS!!!!!!
// http://en.wikipedia.org/wiki/Haversine_formula
func Distance(lng1, lat1, lng2, lat2 float64) float64 {
    // convert to radians
    // must cast radius as float to multiply later
    var la1, ln1, la2, ln2, r float64
    la1 = lat1 * math.Pi / 180
    ln1 = lng1 * math.Pi / 180
    la2 = lat2 * math.Pi / 180
    ln2 = lng2 * math.Pi / 180

    r = EarthRadius // Earth radius in METERS

    // calculate
    h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(ln2-ln1)

    return 2 * r * math.Asin(math.Sqrt(h))
}
