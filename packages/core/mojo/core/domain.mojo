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

/// Domain name type
///
/// For example, when naming a mail domain, the user should satisfy both the
/// rules of this memo and those in RFC-822.  When creating a new host name,
/// the old rules for HOSTS.TXT should be followed.  This avoids problems
/// when old software is converted to use domain names.
///
/// The following syntax will result in fewer problems with many
///
/// applications that use domain names (e.g., mail, TELNET).
///
/// ```
/// <domain> ::= <subdomain> | " "
/// <subdomain> ::= <label> | <subdomain> "." <label>
/// <label> ::= <letter> [ [ <ldh-str> ] <let-dig> ]
/// <ldh-str> ::= <let-dig-hyp> | <let-dig-hyp> <ldh-str>
/// <let-dig-hyp> ::= <let-dig> | "-"
/// <let-dig> ::= <letter> | <digit>
/// <letter> ::= any one of the 52 alphabetic characters A through Z in
/// upper case and a through z in lower case
/// <digit> ::= any one of the ten digits 0 through 9
/// ```
@format(labels:delimiter("."))
type Domain {
    labels: [String] @1
}
