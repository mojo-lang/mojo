#ifndef MOJO_GRAMMAR_TYPE_HPP
#define MOJO_GRAMMAR_TYPE_HPP

#include <mojo/grammar/identifier.hpp>
#include <mojo/grammar/rules.hpp>
#include <pegtl.hh>

namespace mojo {
namespace grammar {

struct union_type;
struct array_type;
struct map_type;
struct tuple_type;
struct function_type;
struct simplified_one_parameter_function_type;

struct type_identifier;

struct primary_type
    : pegtl::sor<array_type, map_type, tuple_type, type_identifier> {};

/**
 * GRAMMAR OF A TYPE
 */
struct data_type : pegtl::sor<union_type, primary_type> {};
struct type : pegtl::sor<function_type, data_type> {};

/**
 * GRAMMAR OF A TYPE ANNOTATION
 */
struct attributes;
struct required_attribute : pegtl::one<'!'> {};
struct optional_attribute : pegtl::one<'?'> {};
struct type_annotation_predication : pegtl::one<':'> {};

using primary_type_annotation = pegtl::seq<
    primary_type,
    pre_pad_opt<pegtl::sor<required_attribute, optional_attribute>, pegtl::ascii::blank>,
    pre_pad_opt<attributes, pegtl::ascii::blank>>;

struct primary_type_annotation_clause : primary_type_annotation {};
struct data_type_annotation_clause : pegtl::sor<union_type, primary_type_annotation> {};
struct type_annotation_clause
    : pegtl::sor<union_type, function_type, primary_type_annotation> {};

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

struct function_type_parameters
    : pegtl::list<type_annotation_clause, pegtl::one<','>, sep> {};

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
                                             type_annotation_clause>,
                                  simplified_one_parameter_function_type> {};

/**
 * GRAMMAR OF A SIMPLIFIED ONE PARAMETER FUNCTION TYPE
 *
 * this is a syntactic sugar for the only one parameter function.
 */
struct simplified_one_parameter_function_type : pegtl::seq<primary_type_annotation_clause,
                                                           seps,
                                                           function_type_return_arrow,
                                                           seps,
                                                           type_annotation_clause> {};

/**
 * GRAMMAR OF AN ARRAY TYPE
 */
struct array_type_content : type_annotation_clause {};
struct array_type
    : pegtl::seq<pegtl::one<'['>, seps, array_type_content, seps, pegtl::one<']'>> {
    using content = array_type_content;
};

/**
 * GRAMMAR OF A DICTIONARY TYPE
 */
struct map_type_separator : pegtl::one<':'> {};
struct map_key
    : pegtl::seq<type_identifier, pre_pad_opt<attributes, pegtl::ascii::blank>> {};
struct map_type_content : pegtl::seq<map_key,
                                            seps,
                                            map_type_separator,
                                            seps,
                                            type_annotation_clause> {};
struct map_type
    : pegtl::seq<pegtl::one<'{'>, seps, map_type_content, seps, pegtl::one<'}'>> {
};

/**
 * GRAMMAR OF A UNION TYPE
 */
struct union_type_separator : pegtl::one<'|'> {};
struct union_type
    : pegtl::seq<
          primary_type_annotation_clause,
          pegtl::plus<seps, union_type_separator, seps, primary_type_annotation_clause>> {
};

/**
 * GRAMMAR OF A INTERSECTION TYPE
 */

/**
 * GRAMMAR OF A TYPE INHERITANCE CLAUSE
 */
struct type_inheritance_clause_content_separator : pegtl::one<','> {};
struct type_inheritance_clause_content
    : pegtl::list<data_type_annotation_clause,
                  type_inheritance_clause_content_separator,
                  sep> {
    using separator = type_inheritance_clause_content_separator;
};
struct type_inheritance_predication : pegtl::one<':'> {};

struct type_inheritance_clause
    : pegtl::seq<type_inheritance_predication, seps, type_inheritance_clause_content> {
    using predication = type_inheritance_predication;
    using content = type_inheritance_clause_content;
};
}  // namespace grammar
}  // namespace mojo

#endif  // MOJO_GRAMMAR_TYPE_HPP