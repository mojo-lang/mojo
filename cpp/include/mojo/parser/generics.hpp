#ifndef MOJO_PARSER_GENERICS_HPP
#define MOJO_PARSER_GENERICS_HPP

#include <mojo/grammar/grammar.hpp>
#include <mojo/parser/control.hpp>
#include <mojo/parser/term_state.hpp>

namespace mojo {
namespace parser {

struct generic_parameter_state : array_term_state {
    generic_parameter_state() : array_term_state("generic_parameter") {
    }
};

template <typename T>
struct generic_clause_state : array_term_state {
    generic_clause_state()
        : array_term_state(std::is_same<T, grammar::generic_parameter>::value
                               ? "generic_parameters"
                               : "generic_arguments") {
    }
};

template <typename Rule>
struct generics_action : pegtl::nothing<Rule> {};

template <>
struct generics_action<grammar::type_name_identifier_separator> {
    template <typename Input>
    static void apply(const Input& input, generic_parameter_state& state) {
        state.push_element();
    }
};

template <>
struct control<grammar::generic_parameter>
    : pegtl::change_state_and_action<grammar::generic_parameter,
                                     generic_parameter_state,
                                     generics_action,
                                     errors> {};

template <typename T>
struct control<grammar::basic_generic_list<T>>
    : pegtl::change_state_and_action<grammar::basic_generic_list<T>,
                                     generic_clause_state<T>,
                                     array_action,
                                     errors> {};
}
}

#endif  // MOJO_PARSER_GENERICS_HPP
