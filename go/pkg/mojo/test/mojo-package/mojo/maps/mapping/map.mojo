
/// 
type Map {
  	id: String @1
  	name: String @2
  	center: LngLat @3
  	zoom: Zoom @4
  	bearing: Angle @5 = 0
  	pitch: Int @6 = 0
    sources: [Source] @7
    layers: [Layer] @8
}
