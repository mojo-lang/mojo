#ifndef MOJO_PARSER_EXPRESSION_STRING_PREFIX_LITERAL_HPP
#define MOJO_PARSER_EXPRESSION_STRING_PREFIX_LITERAL_HPP

#include <mojo/grammar/grammar.hpp>
#include <mojo/parser/control.hpp>
#include <mojo/parser/term_state.hpp>

namespace mojo {
namespace parser {
namespace expression {

struct string_prefix_literal_state : term_state {
    TermPtr literal;

    void success(term_state& in) {
        if (literal) {  // really a prefix literal
            if (this->term) {
                literal->terms.push_back(std::move(this->term));
                in.term = std::move(literal);
            }
            else {
            }
        }
        else {  // will be an error
            std::cout << "lack of prefix literal" << std::endl;
        }
    }
};

template <typename Rule>
struct string_prefix_literal_action : pegtl::nothing<Rule> {};

template <>
struct string_prefix_literal_action<grammar::string_prefix> {
    template <typename Input>
    static void apply(const Input& in, string_prefix_literal_state& state) {
        state.literal = make_term("string_prefix_literal");

        auto term = make_term("string-prefix");
        term->value = in.string();

        state.literal->terms.push_back(std::move(term));
    }
};

}  // namespace expression

template <>
struct control<grammar::string_prefix_literal>
    : pegtl::change_state_and_action<grammar::string_prefix_literal,
                                     expression::string_prefix_literal_state,
                                     expression::string_prefix_literal_action,
                                     errors> {};

}  // namespace parser
}  // namespace mojo

#endif  // MOJO_PARSER_EXPRESSION_STRING_PREFIX_LITERAL_HPP
