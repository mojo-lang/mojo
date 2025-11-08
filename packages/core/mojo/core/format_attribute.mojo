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

/// 对于type的format标注，说明该type类型序列化后的表现形式，并且该Type默认的json等序列化后的类型为String
///
/// 对于字段的format标注，说明该字段字符串是字符串类型，则改字段需要符合的格式要求，一般为正则表达式
/// 如果该字段为结构体，则该接口体需要实现Format接口，并该字段encoding为字符串，该结构体在其他情况下，encoding不变
@apply_to("StructDecl", "ValueDecl")
attribute format: TemplateString | Regex

/// specify the field or the type to encode to struct style when set the format attribution to decl
attribute encoding_as_struct: Bool

/// specify the field or the type to encode to string style when set the format attribution to decl
attribute encoding_as_string: Bool
