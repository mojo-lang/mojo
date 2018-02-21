#ifndef MOJO_PARSER_TERM_TYPES_HPP
#define MOJO_PARSER_TERM_TYPES_HPP

namespace mojo {
namespace parser {

/**
 * Lac
 */
extern const std::string kDocument;

extern const std::string kAttribute;
extern const std::string kAttributes;

/**
 * Types Declaration
 */
extern const std::string kTypeAliasDecl;

extern const std::string kTypeDecl;

extern const std::string kTypeName;
extern const std::string kTypeInheritance;

extern const std::string kTypeAnnotation;

extern const std::string kStructMembers;
extern const std::string kStructField;
extern const std::string kStructFieldGroup;

extern const std::string kEnumDecl;

extern const std::string kEnumMember;
extern const std::string kEnumMembers;

}
}

#endif //MOJO_PARSER_TERM_TYPES_HPP
