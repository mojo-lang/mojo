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

/// Generic Array type
///
///
type Array<T>

///
/// Creates an array with a given length, filled with a default element
///
/// ```
/// repeat(0, 5) == [0,0,0,0,0]
/// repeat("cat", 3) == ["cat","cat","cat"]
/// "cat".repeat(3) == ["cat","cat","cat"]
/// ```
//func repeat<T>(T, UInt) -> Array<T>

/// ## Basics

/// 
//func empty<T>([T]) -> Bool

//func length<T>([T]) -> Size

//func push<T>([T], T) -> [T]

//func append<T>([T], T) -> [T]

//func append<T>([T], [T]) -> [T]

/// ## Get and Set
//func get<T>([T], Int) -> Optional<T>

//func set<T>([T], Int, T) -> [T]