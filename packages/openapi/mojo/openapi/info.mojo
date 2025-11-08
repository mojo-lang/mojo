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

/// The object provides metadata about the API. The metadata MAY be used by the clients if needed, and MAY be presented in editing or documentation generation tools for convenience.
type Info {
    /// The title of the application.
    title: String @1 @required

    /// A short description of the application.
    /// CommonMark syntax syntax MAY be used for rich text representation.
    description: Cached<document.Document> @2

    /// A URL to the Terms of Service for the API.
    terms_of_service: String @3

    /// The contact information for the exposed API.
    contact: Contact @4

    /// The license information for the exposed API.
    license: License @5

    /// Provides the version of the application API (not to be confused with the specification version)
    version: Version @6 @required
}
