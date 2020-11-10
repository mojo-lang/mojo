#ifndef MOJO_PARSER_TERM_TYPES_HPP
#define MOJO_PARSER_TERM_TYPES_HPP

#include <ncraft/core/nonconstructable.hpp>
#include <mojo/core/string.hpp>

namespace mojo {
namespace parser {

    /**
     * Lac
     */
    extern const String kDocument;

    extern const String kAttribute;
    extern const String kAttributeGroup;
    extern const String kAttributes;

    extern const String kGenericArguments;

    extern const String kGenericPathIdentifier;
    extern const String kPathIdentifier;
    extern const String kIdentifier;

    /**
     * Types Declaration
     */
    extern const String kTypeAliasDecl;

    extern const String kTypeDecl;

    extern const String kTypeName;
    extern const String kTypeInheritance;

    extern const String kTypeAnnotation;

    extern const String kStructMembers;
    extern const String kStructField;
    extern const String kStructFieldGroup;

    extern const String kEnumDecl;

    extern const String kEnumMember;
    extern const String kEnumMembers;

}
}

#endif //MOJO_PARSER_TERM_TYPES_HPP
