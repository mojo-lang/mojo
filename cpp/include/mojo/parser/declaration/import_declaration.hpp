#ifndef MOJO_PARSER_DECLARATION_IMPORT_DECLARATION_HPP
#define MOJO_PARSER_DECLARATION_IMPORT_DECLARATION_HPP

#include <mojo/grammar/grammar.hpp>
#include <mojo/parser/control.hpp>
#include <mojo/parser/term_state.hpp>

namespace mojo {
namespace parser {
namespace declaration {

struct import_state : array_term_state {
    import_state() : array_term_state("import_declaration") {
    }
};

struct import_package_identifier_state : array_term_state {
    import_package_identifier_state() : array_term_state("package_identifier") {
    }
};

struct import_items_state : array_term_state {
    import_items_state() : array_term_state("import_items") {
    }
};

struct import_item_state : array_term_state {
    import_item_state() : array_term_state("import_item") {
    }
};

template <typename Rule>
struct import_action : pegtl::nothing<Rule> {};

template<>
struct import_action<grammar::import_separator> {
    template<typename Input>
    static void apply(const Input&, import_state &state) {
        state.push_element();
    }
};

template <>
struct import_action<grammar::import_wild_card> {
    template <typename Input>
    static void apply(const Input&, import_state& state) {
        state.term = make_term("import_wild_card");
        state.term->value = "*";
    }
};

template <>
struct import_action<grammar::key_as> {
    template <typename Input>
    static void apply(const Input&, import_item_state& state) {
        state.push_element();
    }
};

}
template <>
struct control<grammar::import_clause>
    : pegtl::change_state_and_action<grammar::import_clause,
                                     declaration::import_state,
                                     declaration::import_action,
                                     errors> {};

template <>
struct control<grammar::import_wild_card_clause>
    : pegtl::change_state_and_action<grammar::import_wild_card_clause,
                                     declaration::import_state,
                                     declaration::import_action,
                                     errors> {};

template <>
struct control<grammar::import_package_identifier>
    : pegtl::change_state_and_action<grammar::import_package_identifier,
                                     declaration::import_package_identifier_state,
                                     array_action,
                                     errors> {};

template <>
struct control<grammar::import_items>
    : pegtl::change_state_and_action<grammar::import_items,
                                     declaration::import_items_state,
                                     array_action,
                                     errors> {};

template <>
struct control<grammar::import_item>
    : pegtl::change_state_and_action<grammar::import_item,
                                     declaration::import_item_state,
                                     declaration::import_action,
                                     errors> {};

}
}

#endif //MOJO_PARSER_DECLARATION_IMPORT_DECLARATION_HPP
