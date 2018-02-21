#include <string>

namespace mojo {
namespace parser {

using namespace std::literals::string_literals;

/**
 * Lac
 */
const std::string kDocument = "document"s;

const std::string kAttribute = "attribute"s;
const std::string kAttributes = "attributes"s;

/**
 * Types Declaration
 */
const std::string kTypeAliasDecl = "type_alias_declaration"s;

const std::string kTypeDecl = "type_declaration"s;
const std::string kTypeInheritance = "type_inheritance"s;

const std::string kTypeName = "type_name"s;

const std::string kStructMembers = "struct_members"s;
const std::string kStructField = "struct_field"s;
const std::string kStructFieldGroup = "struct_field_group"s;

const std::string kEnumDecl = "enum_declaration"s;

const std::string kEnumMember = "enum_member"s;
const std::string kEnumMembers = "enum_members"s;

}
}
