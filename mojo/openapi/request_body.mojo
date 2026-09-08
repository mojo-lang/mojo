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

/// Describes a single request body.
type RequestBody {
    /// A brief description of the request body.
    /// This could contain examples of use. CommonMark syntax MAY be used for rich text representation.
    description: Cached<document.Document> @1

    /// The content of the request body.
    /// The key is a media type or media type range and the value describes it.
    /// For requests that match multiple keys, only the most specific key is applicable.
    /// e.g. text/plain overrides text/*
    content: {String: MediaType} @2 @required

    /// Determines if the request body is required in the request.
    required: Bool @3 @default(false)
}
