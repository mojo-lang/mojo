#ifndef MOJO_ATTRIBUTE_PARSER_HPP
#define MOJO_ATTRIBUTE_PARSER_HPP

#include <mojo/lang/attribute.mojo.hpp>
#include <mojo/parser/ast/parser.hpp>
#include <mojo/parser/term_types.hpp>

#include <mojo/core/result.hpp>

namespace mojo {
namespace parser {
namespace ast {

class AttributeParser {
public:
    Result<lang::Attribute> operator()(const TermPtr& term) {
        auto p = Parser(term);
        if (p.expect(kAttribute)) {
        }
    }
};

}  // namespace ast
}  // namespace parser
}  // namespace mojo

#endif  // MOJO_ATTRIBUTE_PARSER_HPP
