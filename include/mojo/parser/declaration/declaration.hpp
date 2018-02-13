#ifndef MOJO_PARSER_DECLARATION_DECLARATION_HPP
#define MOJO_PARSER_DECLARATION_DECLARATION_HPP

#include <mojo/grammar/grammar.hpp>
#include <mojo/parser/control.hpp>
#include <mojo/parser/term_state.hpp>

namespace mojo {
namespace parser {
namespace declaration {

struct declaration_term_state : term_state {
    std::vector<TermPtr> terms;

    declaration_term_state() = default;

    void add_declaration() {
        if (term) {
            terms.push_back(std::move(term));
            term.reset();
        } else {
            // should log warning for push a null element
        }
    }

    void success(term_state &in) {
        if (term) {
            if (!terms.empty()) {
                if (terms.size() < 3) {
                    term->terms.insert(term->terms.begin(), terms.begin(), terms.end());
                    in.term = std::move(term);
                } else {
                    // error
                }
            } else {
                in.term = std::move(term);
            }
        } else {
            // error
        }
    }
};

template<typename Rule>
struct declaration_action : pegtl::nothing<Rule> {};

template<>
struct declaration_action<grammar::document_separator> {
    template<typename Input>
    static void apply(const Input &, declaration_term_state &state) {
        state.add_declaration();
    }
};

template<>
struct declaration_action<grammar::attribute_separator> {
    template<typename Input>
    static void apply(const Input &, declaration_term_state &state) {
        state.add_declaration();
    }
};

}

template<>
struct control<grammar::declaration>
    : pegtl::change_state_and_action<grammar::declaration,
                                     declaration::declaration_term_state,
                                     declaration::declaration_action,
                                     errors> {
};

template <>
struct control<grammar::type_member>
    : pegtl::change_state_and_action<grammar::type_member,
                                     declaration::declaration_term_state,
                                     declaration::declaration_action,
                                     errors> {};

template <>
struct control<grammar::interface_member>
    : pegtl::change_state_and_action<grammar::interface_member,
                                     declaration::declaration_term_state,
                                     declaration::declaration_action,
                                     errors> {};
}
}
#endif //MOJO_PARSER_DECLARATION_DECLARATION_HPP
