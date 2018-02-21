#ifndef MOJO_STRUCT_FIELD_PARSER_HPP
#define MOJO_STRUCT_FIELD_PARSER_HPP

#include <mojo/lang/struct_field.hpp>
#include <mojo/parser/ast/value_decl_parser.hpp>

namespace mojo {
namespace parser {
namespace ast {

class StructFieldParser {
public:
    lang::StructType::Field operator()(const TermPtr& term, lang::PackagePtr& package) {
        if (Parser(term).expect(kStructField)) {
            return ValueDeclParser<lang::StructType::Field>{}(term, package);
        }
    }
};

}  // namespace ast
}  // namespace parser
}  // namespace mojo

#endif  // MOJO_STRUCT_FIELD_PARSER_HPP
