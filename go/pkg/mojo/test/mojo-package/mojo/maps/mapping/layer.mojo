///
///
type Layer {
  	id: String @1
  	meta: Object @2
  	min_zoom: Zoom @3
  	max_zoom: Zoom @4
  	styles: {String: Style} @5
  	visible: Bool @6 = true
}

/// 
func attach_style(layer:Layer, style:Style) -> Layer {
	layer.styles.add('default', style)
}

/// 
func attach_style(layer:Layer, name:String, style:Style) -> Layer {
	layer.styles.add(name, style)
}