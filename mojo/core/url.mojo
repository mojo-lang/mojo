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

/// URl type
/// 
/// Uniform Resource Identifier (URI) is a string of characters used to identify a resource.
///
/// A Uniform Resource Name (URN) may be compared to a person's name,
/// while a Uniform Resource Locator (URL) may be compared to their street address.
/// In other words, a URN identifies an item and a URL provides a method for finding it.
///
/// https://tools.ietf.org/html/rfc3986
///
///
@format("{scheme}://{authority.user_info@}{authority.host}{:authority.port}/{path}{#fragment}{?query}")
type Url {
    type Authority {
        user_info: String @1
        host: String @2 @format('')
        port: Int @3
    }

    //@format("{segments:/}")
    //type Path {
    //    segments:[String] @1
    //}

    ///
    @format("{key=value:&}")
    type Query: {String: [String]}

    ///
    scheme: String @1 //@format(r'[\w\d+-.]+')
    
    ///
    authority: Authority @2
    
    ///
    path: String @3
    
    ///
    query: Query @5
    
    ///
    fragment: String @7
}

//func addQuery(url: Url, key: String, value: String)

//func addQueries(url: Url, parameters: {String: String})

//func getQuery(url: Url, key: String) -> [String]