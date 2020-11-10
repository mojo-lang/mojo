#ifndef MOJO_ATTRIBUTES_PARSER_HPP
#define MOJO_ATTRIBUTES_PARSER_HPP

#include <mojo/core/array.hpp>
#include <mojo/lang/attribute.mojo.hpp>
#include <mojo/parser/ast/attribute_parser.hpp>
#include <mojo/parser/ast/parser.hpp>
#include <mojo/parser/term_types.hpp>

namespace mojo {
namespace parser {
namespace ast {

class AttributesParser {
public:
    Array<lang::Attribute> operator()(lang::TermPtr term) {
        auto p = Parser(term);
        if (p.expect(kAttributes)) {
            Array<lang::Attribute> attributes;
            for (auto&& t : term->terms) {
                auto p = Parser(t);
                if (p.expect(kAttribute)) {
                    attributes.push_back(AttributeParser{}(std::move(t)));
                }
            }
            return std::move(attributes);
        }
        else {
            return Array<lang::Attribute>{};
        }
    }
};

}  // namespace ast
}  // namespace parser
}  // namespace mojo

#endif  // MOJO_ATTRIBUTES_PARSER_HPP
