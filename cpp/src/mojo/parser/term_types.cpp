#include <string>
#include <mojo/parser/term_types.hpp>

namespace mojo {
namespace parser {

using namespace std::literals::string_literals;

/**
 * Lac
 */
const String kDocument = "document"s;

const String kAttribute = "attribute"s;
const String kAttributeGroup = "attribute_group"s;
const String kAttributes = "attributes"s;

const String kGenericArguments = "generic_arguments"s;
const String kGenericPathIdentifier = "generic_path_identifier"s;
const String kPathIdentifier = "path_identifier"s;
const String kIdentifier = "identifier"s;

/**
 * Types Declaration
 */
const String kTypeAliasDecl = "type_alias_declaration"s;

const String kTypeDecl = "type_declaration"s;
const String kTypeInheritance = "type_inheritance"s;

const String kTypeName = "type_name"s;

const String kStructMembers = "struct_members"s;
const String kStructField = "struct_field"s;
const String kStructFieldGroup = "struct_field_group"s;

const String kEnumDecl = "enum_declaration"s;

const String kEnumMember = "enum_member"s;
const String kEnumMembers = "enum_members"s;

}
}
