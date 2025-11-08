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

@format('#{red:0x}{green:0x}{blue:0x}{alpha_value:0x}')
type Color {
    red:   UInt8 @1
    green: UInt8 @2
    blue:  UInt8 @3
    alpha: FloatValue @4 @in(0..1)
}

//type RGB = Color

//type HLS {
//}

//const {
//    red = '#ff0000'
//    blue = '#ff0000'
//}

//
//Color(str:String)

//
//Color(r:UInt8, g:UInt8, b:UInt8)

///
//func rgb(r:UInt8, g:UInt8, b:UInt8, a) -> Color

///
//func hsl(hue: Degree, saturation: Percentage, lightness: Percentage, alpha:Real) -> Color

///
//func lighten(c:Color, n:Integer) -> Color    // `red.ligthen(10)`

///
//func saturate(c:Color,)

///
//func spin(c:Color,)

///
//func fade(c:Color,)

///
//func mix(c:Color,)

///
//func darken(c:Color,)
