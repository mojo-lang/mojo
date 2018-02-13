#ifndef MOJO_PARSER_TERM_TYPES_HPP
#define MOJO_PARSER_TERM_TYPES_HPP

using namespace std::literals::string_literals;

namespace mojo {
namespace parser {

const std::string kTypeAliasDecl = "type_alias_declaration"s;

const std::string kTypeDecl = "type_declaration"s;

const std::string kStructMember = "type_member"s;
const std::string kStructMembers = "type_members"s;
const std::string kStructMemberGroup = "type_member_group"s;


const std::string kEnumDecl = "enum_declaration"s;

const std::string kEnumMember = "enum_member"s;
const std::string kEnumMembers = "enum_members"s;

}
}

#endif //MOJO_PARSER_TERM_TYPES_HPP
