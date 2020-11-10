#ifndef MOJO_GRAMMAR_GENERIC_HPP
#define MOJO_GRAMMAR_GENERIC_HPP

#include <pegtl.hh>
#include <mojo/grammar/type.hpp>

namespace mojo {
namespace grammar {

template <typename T>
struct basic_generic_list : pegtl::list<T, pegtl::one<','>, sep> {};

struct generic_clause_begin : pegtl::one<'<'> {};
struct generic_clause_end : pegtl::one<'>'> {};

template <typename T>
struct basic_generic_clause
    : pegtl::seq<generic_clause_begin, seps, basic_generic_list<T>, seps, generic_clause_end> {
    using begin = generic_clause_begin;
    using end = generic_clause_end;
    using element = T;
    using content = basic_generic_list<T>;
};

struct type_name_identifier_separator : pegtl::one<':'> {};

/**
 * GRAMMAR OF A GENERIC PARAMETER
 */
struct generic_parameter_constrict : pegtl::seq<pegtl::one<':'>, seps, type_identifier> {
    using predication = pegtl::one<':'>;
};
struct generic_parameter : pegtl::seq<type_name, pre_pad_opt<generic_parameter_constrict, sep>> {};

/**
 * GRAMMAR OF A GENERIC PARAMETER LIST
 */
using generic_parameters = basic_generic_list<generic_parameter>;

/**
 * GRAMMAR OF A GENERIC PARAMETER CLAUSE
 */
struct generic_parameter_clause : basic_generic_clause<generic_parameter> {};

/**
 * GRAMMAR OF A GENERIC ARGUMENT
 */
//struct generic_argument : type_annotation_clause {};

/**
 * GRAMMAR OF A GENERIC ARGUMENT LIST
 */
using generic_arguments = basic_generic_list<type_annotation_clause>;

/**
 * GRAMMAR OF A GENERIC ARGUMENT CLAUSE
 */
struct generic_argument_clause : basic_generic_clause<type_annotation_clause> {};
}
}

#endif  // MOJO_GRAMMAR_GENERIC_HPP