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

@label_format('{}')
type IPAddress = IPv4 @1 | IPv6 @2

@irregular_case_rule("IPv4", "Ipv4")
type IPv4: UInt32

@irregular_case_rule("IPv6", "Ipv6")
type IPv6: UInt64
