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

// ObjectLiteralExpr - An expression of the form
// '{red: 1, blue: 0, green: 0, alpha: 1}' with a name and a list
// argument. The components of the list argument are meant to be themselves
// constant.
type ObjectLiteralExpr : LiteralExpr {
    type Field {
        /// position of first character belonging to the Field
        start_position: Position @1

        /// position of first character immediately after the Field
        end_position: Position @2

        name: String @3

        value: Expression @4
    }

    fields: [Field] @20
}
