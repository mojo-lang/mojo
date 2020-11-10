#ifndef MOJO_AST_PARSER_HPP
#define MOJO_AST_PARSER_HPP

#include <mojo/lang/nominal_type.mojo.hpp>
#include <mojo/lang/struct_decl.mojo.hpp>
#include <mojo/parser/ast/attributes_parser.hpp>
#include <mojo/parser/ast/parser.hpp>
#include <mojo/parser/ast/type_inheritance_parser.hpp>

namespace mojo {
namespace parser {
namespace ast {

class TypeDeclParser {
public:
    PackagePtr operator()(PackagePtr package, TermPtr term) {
        if (parser_.expect(kTypeDecl)) {
            if (isDataType()) {
                parseDataDecl();
            }
            else if (isStructType()) {
                parseStructDecl();
            }
        }
    }

    DataDeclPtr parseDataDecl() {
    }

    String parseTypeName(const TermPtr& typeName) {
        if (parser_.expect(kTypeName)) {
            return std::move(typeName->value);
        }
        else {
            return "";
        }
    }

    StructDeclPtr parseStructDecl() {
        auto structDecl = std::make_shared<lang::StructDecl>();

        structDecl->name = parseTypeName(parser_.getTerm(kTypeName));

        structDecl->document = "";
        structDecl->attributes = std::move(AttributesParser{}(parser_.getTerm(kAttributes));

        auto structType = std::make_shared<lang::StructType>();
        structType->inherits = std::move(TypeInheritanceParser{}(parser_.getTerm(kTypeInheritance), package_));
        auto members = parser_.getTerm(kStructMembers);
        for (auto const& term : members->terms) {
            if (term->type == kStructField) {
            }
            else if (term->type == kStructFieldGroup) {
            }
            else if (term->type == kTypeDecl) {
            }
            else if (term->type == kTypeAliasDecl) {
            }
            else if (term->type == kEnumDecl) {
            }
            else {
                // error
            }
        }
    }

    bool isDataType() {
        if (!parser_.contains(kStructMembers) || !parser_.contains(kTypeInheritance)) {
            return true;
        }
        else if (parser_.termsCount(kTypeInheritance) == 1) {
            NominalTypePtr type;
            return isDataDecl(type->typeDeclaration);
        }
        else {
            return false;
        }
    }

    bool isStructType() {
        return !isDataType();
    }

private:
    Parser parser_;
    PackagePtr package_;
    TermPtr term_;
    TermPtr current_;
};

}  // namespace ast
}  // namespace parser
}  // namespace mojo

#endif  // MOJO_AST_PARSER_HPP
