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

///
///
///
type Range<T> {
    start: T @1 // first
    end: T @2   // last

    //step: Int @3 = 1 //

    start_excluded: Bool @9
    end_included: Bool @10
}

type IntRange = Range<Int64>

// 12..=45    12..45
// 12..<45    12..<45
// 12<..45    12<..45
// 12<.<45    12<..<45