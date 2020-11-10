# Mojo Map Library



扩展Mojo，支持配置地图显示风格的DSL语言；支持地图显示模板配置。

## 类型说明



### Zoom类型

```
type Zoom : UInt8 @max(24)
```


### Angle类型

```
type Angle : UInt8 @max(360)
```

```
TemplateStr = `xxx`;
Regular = /xxx/

reg = Regular("xx");
ts = TemplateStr("");


```

### Pixel

```
type Pixel : Int;

```

### Percentage类型

```
type Percentage : UInt8 @max(100)

```

### ratio类型

```
type Ratio : Real @min(0) @max(1)

type Ratio : Real @{min:0, max:1}
   ```

### Color类型

```
Type Color {
  red:UInt8,
  green:UInt8,
  blue:UInt8,
  alpha: UInt8 @max(100)
}
```



```
const {
    red = rgb()
    blue = rgb()
}
'#b323'
red
blue
rgb(r:UInt8, g:UInt8, b:UInt8)
rgba()
hsv(h, s, v)
```

#### color 构建函数

1. `Color(str:String)`
2. `Color(name:Symbol)`
3. `rgb(r:UInt8, g:UInt8, b[, a]) -> Color`
4. `hsl(hue:Degree, saturation:Percentage, lightness:Percentage[, alpha:Real])`

#### color操作函数

1. `lighten(c:Color, n:Integer) -> Color`    // `red.ligthen(10)`
2. `saturate(c:Color,)`
3. `spin(c:Color,)`
4. `fade(c:Color,)`
5. `mix(c:Color,)`
6. `darken(c:Color,)`

### Source类型

```
type Source {
  	id: String
}

type RasterSource : Source {
}

type VTileSource : Source {
  	tile: Url,
  	minzoom: Zoom,
  	maxzoom: Zoom
}

type DbSource : Source {

}
```



### Layer类型

attribute layout : Bool {
}

attribute paint : Bool {
}

type Translate {
    enum Anchor { map, viewport }

    ///
    ///
    anchor: Anchor = map

    ///
    ///
    offset: Point
}

enum Visibility { visible, none }

type Line {
    enum Cap { butt, round, square }
    enum Join { bevel, round, miter }

    /**
     */
    cap: Cap = butt @layout

    /**
     *
     */
    join: Join = miter

    /**
     */
    miter_limit: Real = 2

    /**
     */
    round_limit: Real = 1.05

    /**
     *
     */
    visibility: Visibility = visible

    /**
     *
     */
    opacity: Opacity @paint

    /**
     */
    antialias: Bool = true

    /**
     *
     */
    color: Color = black @paint

    /**
     *
     */
    translate: Translate

    /**
     */
    width: Pixel = 1

    /**
     */
    gap_width: Pixel = 0

    /**
     *
     */
    offset: Pixel = 0

    /**
     *
     */
    blur: Pixel = 0

    /**
     *
     */
    dasharray: Array(Pixel)

    /**
     * Name of image in sprite to use for drawing image lines. For seamless patterns,
     * image width must be a factor of two (2, 4, 8, ..., 512).
     */
    pattern: String
}

//Polygon
type Fill {
    /**
     *
     */
    visibility: Visibility = visible

    antialias: Bool

    color: Color
    opacity: Opacity
    translate: Translate
    pattern: String

    outline: Line
}

//Point
type Symbol {
    enum Placement { point, line }

    placement: Placement

    spacing: Pixel = 250 @requires(placement == line)

    avoid_edges: Bool

    icon: Icon

    text: Text

    visibility: Visibility = visible

}

enum Alignment {map, viewport, auto}

type Halo {
        color: Color = black
        width: Pixel = 0
        blur: Pixel = 0
    }

type Icon {
    allow_overlap: Bool = false
    ignore_placement: Bool = false
    optional: Bool = false
    rotation_alignment: Alignment = auto

    size: Int = 1
    text_fit : Int
    text_fit_padding: Array<Int, 4>

    image: String
    rotate: Degree
    padding: Pixel
    keep_upright: Bool
    offset: Point

    ///-> paint
    opacity: Opacity = 1
    color: Color = black
    halo: Halo


}

type Text {
    enum Transform {none, uppercase, lowercase}
    pitch_alignment : Alignment = auto
    rotation_alignment : Alignment = auto

    field: String
    font: Array(String)

     size: Int
     max_width: Em = 10

     line_height: Em = 1.2
     letter_spacing: Em = 0

     enum Justify { left, center, right }
     enum Anchor { center, left, right, top, bottom, top_left, top_right, bottom_left, bottom_right}
     justify: Justify = center
     anchor: Anchor = center

     max_angle: Degree = 45
     rotate: Degree = 0
     padding: Pixel = 2
     keep_upright: Bool = true
     transform: Transform = none

      offset: Em = 0

      allow_overlap: Bool = false
      ignore_placement: Bool = false

      optional: Bool = false


      ///-> paint
      opacity: Opacity = 1
      color: Color = black
      halo: Halo
      translate: Translate
}

type RasterStyle {
    opacity: Opacity = 1
    hue_rotate: Degree = 0
    brightness_min: Ratio = 0
    brightness_max: Ratio = 1
    saturation: Ratio = 0
    contrast: Ratio = 0
}

type Circle {
}

//Building
type FillExtrusion {
    visibility: Visibility = visible
    opacity: Opacity = 1
    color: Color = black
    translate: Translate
    pattern: String
    height: Int
    base: Int
}

```
type Layer {
  	id: String,
  	meta: Object,
  	min-zoom: Zoom,
  	max-zoom: Zoom,
  	styles: map<string,Style>,
  	visibility: Bool
}

source_id.roads | style({
		line-width: match(kind) { _: 1.5; 'motorway': 2; 'main': 4; },
	}
```

### Map类型

```c++
type Map {
  	id: String,
  	name: String,
  	center: Latlng,
  	zoom:Real,
  	bearing: Angle,
  	pitch:Integer,
  	sources:Source,
  	layers: Sequence<Layer>
}
```



### MapTemplate类型

```
type MapTemplate : Map {
}

tmpl = MapTemplate {
	name: '${name}',
	background: style({color: #fff}),
	layers: [
      	basic: tile({
	    	url:
	    }),
		european ｜ Style{ polygon-fill: #fff }
	]
};
```

#### map构建函数

```
map(tmpl:MapTemplate, params:Object);
```

### Viz类型

是对Map的扩展，可以定义整个地图的显示内容，包含control定义

```
Viz {
	map: xxx,
}
```

### Style:Object类型

是可嵌套的类型。

### 系统内置变量：

#### zoom

```
match(kind) {
	motorway: 2,
	main: 3,
	_: 0
}

if (kind) {} elif () {} else {}
```



```
roads | style({
		line.width: switch(kind) {
		                'motorway': scale(zoom, [4, 334], [5, 4], [8, 5])
		                'trunkway': scale()
		            }
	}

func poi_ldir(point:Point) {
	 planet_osm_line ｜ kind in ['motorway', 'trunk']
		            | has_field(way) | buffer(100)
		sort(distance(point), asc)
		limit(1)
		closest_point(point)
		azimuth(point)
		to_integer()
}

layer | zoom >= 4 and zoom <= 10 | new:style({
		line.color: red,
		line.width: match(zoom) {_:1.5; 5: 1; 7:2;}
		//line-width: scale(zoom, 1.5, {5, 1}, {7, 2})
} | attach_style(glow, {})

style({
	 bottomline.line.width: 6,
    middleline.line.width: 4,
}

make_style(object:Style)
attach_style(based:Style, name:Symbol, attached:Style)

cities |
	 (zoom>=4 && population>1M) ||
     (zoom>=5 && population>500000) |
    style({
    	text.name: name,
    	text.facename: ''
    }

tornadoes | case:style({
    marker.width: match(zoom) {
      	zoom <= 7: match(fscale) {
          	3: 12, 4: 14, 5: 15
      	}
      	_: match(fscale) {
      	}
    }

    marker-width: match(fscale) {
    	0: 0.5; 1: 1;
    }

    mark.fill:'#f45'
} | attach_style(fill, {

})
```

##
