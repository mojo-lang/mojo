#ifndef MOJO_STRUCT_FIELD_GROUP_PARSER_HPP
#define MOJO_STRUCT_FIELD_GROUP_PARSER_HPP

#include <mojo/parser/ast/struct_field_parser.hpp>

namespace mojo {
namespace parser {
namespace ast {

class StructFieldGroupParser {
public:
    bool operator()(const TermPtr& term, lang::PackagePtr& package, lang::StructType& type) {
        // if (Parser(term).expect(kStructField)) {
        //    return ValueDeclParser<lang::StructType::Field>{}(term, package);
        //}
    }
};

}  // namespace ast
}  // namespace parser
}  // namespace mojo

#endif  // MOJO_STRUCT_FIELD_GROUP_PARSER_HPP
