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
/// An email address identifies an email box to which email messages are delivered.
///
/// The general format of an email address is local-part@domain,
/// and a specific example is jsmith@example.com. An address consists of two parts.
/// The part before the @ symbol (local-part) identifies the name of a mailbox.
/// This is often the username of the recipient, e.g., jsmith.
/// The part after the @ symbol (domain) is a domain name that represents the administrative realm for the mail box,
/// e.g., a company's domain name, example.com.
///
@example("jsmith@example.com")
@format('{local_part}@{domain}')
type EmailAddress {
    ///
    local_part: String @1

    ///
    domain: Domain @2
}
