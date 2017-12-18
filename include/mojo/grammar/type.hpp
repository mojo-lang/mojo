#ifndef MOJO_GRAMMAR_TYPE_HPP
#define MOJO_GRAMMAR_TYPE_HPP

#include <pegtl.hh>
#include <mojo/grammar/identifier.hpp>
#include <mojo/grammar/rules.hpp>

namespace mojo {
namespace grammar {

struct array_type;
struct dictionary_type;
struct tuple_type;
struct function_type;
struct simplified_one_parameter_function_type;

struct type_identifier;
struct metatype_type;
struct primary_type
    : pegtl::sor<array_type, dictionary_type, tuple_type, type_identifier> {};

/**
 * GRAMMAR OF A TYPE
 */
struct type : pegtl::sor<function_type, primary_type, metatype_type> {};

/**
 * GRAMMAR OF A TYPE ANNOTATION
 */
struct attributes;
struct type_annotation_predication : pegtl::one<':'> {};
struct type_annotation_clause
    : pegtl::sor<pegtl::seq<primary_type, pre_pad_opt<attributes, pegtl::ascii::blank>>,
                 function_type> {};

struct type_annotation
    : pegtl::seq<type_annotation_predication, blanks, type_annotation_clause> {
    using predication = type_annotation_predication;
    using content = type_annotation_clause;
};

/**
 * GRAMMAR OF A PACKAGE IDENTIFIER
 */
struct package_identifier : pegtl::list<identifier, pegtl::one<'.'>> {};

struct type_name_head : pegtl::ascii::upper {};
struct type_name_character : pegtl::sor<pegtl::ascii::alpha, pegtl::ascii::digit> {};

/**
 * GRAMMAR OF A TYPE NAME
 */
struct type_name : pegtl::seq<type_name_head, pegtl::star<type_name_character>> {};

/**
 * GRAMMAR OF A TYPE IDENTIFIER
 */
struct type_identifier;
struct generic_argument_clause;

struct generic_type_name : pegtl::seq<type_name, pegtl::opt<generic_argument_clause>> {};
struct package_identifier_separator : pegtl::one<'.'> {};
struct generic_type_name_separator : pegtl::one<'.'> {};
struct type_identifier
    : pegtl::seq<pegtl::opt<package_identifier, package_identifier_separator>,
                 pegtl::list<generic_type_name, generic_type_name_separator>> {};

struct tuple_type_element
    : pegtl::sor<type_annotation_clause, pegtl::seq<identifier, seps, type_annotation>> {
};

using tuple_type_body = pegtl::list<tuple_type_element, pegtl::one<','>, sep>;

/**
 * GRAMMAR OF A TUPLE TYPE
 */
struct tuple_type
    : pegtl::seq<pegtl::one<'('>, seps, tuple_type_body, seps, pegtl::one<')'>> {
    using content = tuple_type_body;
};

struct function_type_parameters : pegtl::list<type, pegtl::one<','>, sep> {};

/**
 * GRAMMAR OF A FUNCTION TYPE
 */
struct function_type_return_arrow : pegtl_string_t("->") {};
struct function_type : pegtl::sor<pegtl::seq<pegtl::one<'('>,
                                             seps,
                                             pegtl::opt<function_type_parameters>,
                                             seps,
                                             pegtl::one<')'>,
                                             seps,
                                             function_type_return_arrow,
                                             seps,
                                             type>,
                                  simplified_one_parameter_function_type> {};

/**
 * GRAMMAR OF A SIMPLIFIED ONE PARAMETER FUNCTION TYPE
 *
 * this is a syntactic sugar for the only one parameter function.
 */
struct simplified_one_parameter_function_type
    : pegtl::seq<primary_type, seps, function_type_return_arrow, seps, type> {};

/**
 * GRAMMAR OF AN ARRAY TYPE
 */
struct array_type_content : type {};
struct array_type
    : pegtl::seq<pegtl::one<'['>, seps, array_type_content, seps, pegtl::one<']'>> {
    using content = array_type_content;
};

/**
 * GRAMMAR OF A DICTIONARY TYPE
 */

struct dictionary_type_separator : pegtl::one<':'> {};
struct dictionary_type_content
    : pegtl::seq<type_identifier, seps, dictionary_type_separator, seps, type> {};
struct dictionary_type
    : pegtl::seq<pegtl::one<'{'>, seps, dictionary_type_content, seps, pegtl::one<'}'>> {
};

/**
 *  GRAMMAR OF A METATYPE TYPE
 */
struct metatype_type_suffix : pegtl_string_t(".Type") {};
struct metatype_type : pegtl::seq<primary_type, metatype_type_suffix> {};

/**
 * GRAMMAR OF A TYPE INHERITANCE CLAUSE
 */
struct type_inheritance_clause_content_separator : pegtl::one<','> {};
struct type_inheritance_clause_content
    : pegtl::list<type_identifier, type_inheritance_clause_content_separator, sep> {
    using separator = type_inheritance_clause_content_separator;
};
struct type_inheritance_predication : pegtl::one<':'> {};

struct type_inheritance_clause
    : pegtl::seq<type_inheritance_predication, seps, type_inheritance_clause_content> {
    using predication = type_inheritance_predication;
    using content = type_inheritance_clause_content;
};
}
}

#endif  // MOJO_GRAMMAR_TYPE_HPP