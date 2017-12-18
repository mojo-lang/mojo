#ifndef MOJO_PARSER_TYPE_HPP
#define MOJO_PARSER_TYPE_HPP

#include <mojo/grammar/grammar.hpp>
#include <mojo/parser/control.hpp>
#include <mojo/parser/term_state.hpp>

namespace mojo {
namespace parser {

struct type_annotation_state : array_term_state {
    type_annotation_state() : array_term_state("type_annotation") {
    }
};

struct package_identifier_state : array_term_state {
    package_identifier_state() : array_term_state("package_identifier") {
    }
};

struct generic_type_name_state : array_term_state {
    generic_type_name_state() : array_term_state("generic_type_name") {
    }
};

struct type_identifier_state : array_term_state {
    type_identifier_state() : array_term_state("type_identifier") {
    }
};

struct array_type_state : array_term_state {
    array_type_state() : array_term_state("array_type") {
    }
};

struct dictionary_type_state : array_term_state {
    dictionary_type_state() : array_term_state("dictionary_type") {
    }
};

struct tuple_type_element_state : array_term_state {
    tuple_type_element_state() : array_term_state("tuple_type_element") {
    }
};

struct tuple_type_state : array_term_state {
    tuple_type_state() : array_term_state("tuple_type") {
    }
};

struct function_type_parameters_state : array_term_state {
    function_type_parameters_state() : array_term_state("function_type_parameters") {
    }
};

struct function_type_state : array_term_state {
    function_type_state() : array_term_state("function_type") {
    }
};

struct type_inheritance_state : array_term_state {
    type_inheritance_state() : array_term_state("type_inheritance") {
    }
};

template <typename Rule>
struct type_action : pegtl::nothing<Rule> {};

template <>
struct type_action<grammar::attributes::predication> {
    template <typename Input>
    static void apply(const Input&, type_annotation_state& state) {
        state.push_element();
    }
};

template <>
struct type_action<grammar::type_name> {
    template <typename Input>
    static void apply(const Input& input, term_state& state) {
        state.term = make_term("type_name", input.string());
    }
};

template <>
struct type_action<grammar::generic_argument_clause::begin> {
    template <typename Input>
    static void apply(const Input&, generic_type_name_state& state) {
        state.push_element();
    }
};

template <>
struct type_action<grammar::package_identifier_separator> {
    template <typename Input>
    static void apply(const Input&, type_identifier_state& state) {
        state.push_element();
    }
};

template <>
struct type_action<grammar::generic_type_name_separator> {
    template <typename Input>
    static void apply(const Input&, type_identifier_state& state) {
        state.push_element();
    }
};

template <>
struct type_action<grammar::dictionary_type_separator> {
    template <typename Input>
    static void apply(const Input&, dictionary_type_state& state) {
        state.push_element();
    }
};

template <>
struct type_action<grammar::metatype_type_suffix> {
    template <typename Input>
    static void apply(const Input&, term_state& state) {
        auto new_term = make_term("metatype_type");
        new_term->terms.emplace_back(std::move(state.term));
        state.term = new_term;
    }
};

template <>
struct type_action<grammar::type_annotation::predication> {
    template <typename Input>
    static void apply(const Input&, tuple_type_element_state& state) {
        state.push_element();
    }
};

template <>
struct type_action<grammar::function_type_return_arrow> {
    template <typename Input>
    static void apply(const Input&, function_type_state& state) {
        state.push_element();
    }
};

template <>
struct control<grammar::type_name>
    : pegtl::change_action<grammar::type_name, type_action, errors> {};

template <>
struct control<grammar::generic_type_name>
    : pegtl::change_state_and_action<grammar::generic_type_name,
                                     generic_type_name_state,
                                     type_action,
                                     errors> {};

template <>
struct control<grammar::package_identifier>
    : pegtl::change_state_and_action<grammar::package_identifier,
                                     package_identifier_state,
                                     array_action,
                                     errors> {};

template <>
struct control<grammar::type_identifier>
    : pegtl::change_state_and_action<grammar::type_identifier,
                                     type_identifier_state,
                                     type_action,
                                     errors> {};

template <>
struct control<grammar::type_annotation_clause>
    : pegtl::change_state_and_action<grammar::type_annotation_clause,
                                     type_annotation_state,
                                     type_action,
                                     errors> {};

template <>
struct control<grammar::array_type::content>
    : pegtl::change_state_and_action<grammar::array_type::content,
                                     array_type_state,
                                     array_action,
                                     errors> {};

template <>
struct control<grammar::dictionary_type_content>
    : pegtl::change_state_and_action<grammar::dictionary_type_content,
                                     dictionary_type_state,
                                     type_action,
                                     errors> {};

template <>
struct control<grammar::metatype_type>
    : pegtl::change_action<grammar::metatype_type, type_action, errors> {};

template <>
struct control<grammar::tuple_type_element>
    : pegtl::change_state_and_action<grammar::tuple_type_element,
                                     tuple_type_element_state,
                                     type_action,
                                     errors> {};

template <>
struct control<grammar::tuple_type_body>
    : pegtl::change_state_and_action<grammar::tuple_type_body,
                                     tuple_type_state,
                                     array_action,
                                     errors> {};

template <>
struct control<grammar::function_type_parameters>
    : pegtl::change_state_and_action<grammar::function_type_parameters,
                                     function_type_parameters_state,
                                     array_action,
                                     errors> {};

template <>
struct control<grammar::function_type>
    : pegtl::change_state_and_action<grammar::function_type,
                                     function_type_state,
                                     type_action,
                                     errors> {};

template <>
struct control<grammar::simplified_one_parameter_function_type>
    : pegtl::change_state_and_action<grammar::simplified_one_parameter_function_type,
                                     function_type_state,
                                     type_action,
                                     errors> {};

template <>
struct control<grammar::type_inheritance_clause::content>
    : pegtl::change_state_and_action<grammar::type_inheritance_clause::content,
                                     type_inheritance_state,
                                     array_action,
                                     errors> {};
}
}

#endif  // MOJO_PARSER_TYPE_HPP