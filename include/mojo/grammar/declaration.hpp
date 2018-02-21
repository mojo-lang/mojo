#ifndef MOJO_GRAMMAR_DECLARATION_HPP
#define MOJO_GRAMMAR_DECLARATION_HPP

#include <mojo/grammar/statement.hpp>
#include <mojo/grammar/type.hpp>
#include <pegtl.hh>

namespace mojo {
namespace grammar {

/**
 * GRAMMAR OF AN PACKAGE DECLARATION
 */
struct package_declaration
    : pegtl::seq<key_package, seps, package_identifier, seps, object_literal> {};

/**
 * GRAMMAR OF AN IMPORT DECLARATION
 */
struct import_wild_card : pegtl::one<'*'> {};
struct import_separator
    : pegtl::at<pegtl::sor<identifier_head, type_name_head, import_wild_card>> {};
struct import_package_identifier : pegtl::plus<pegtl::seq<identifier, pegtl::one<'.'>>> {
};
struct import_item_identifier : pegtl::sor<identifier, type_name> {};
struct import_item
    : pegtl::seq<import_item_identifier,
                 pegtl::pad_opt<pegtl::seq<key_as, sep, import_item_identifier>, sep>> {};

struct import_items : pegtl::list<import_item, pegtl::one<','>, sep> {};

struct import_clause
    : pegtl::seq<pegtl::opt<import_package_identifier, import_separator>, import_items> {
};

struct import_wild_card_clause
    : pegtl::seq<import_package_identifier, import_separator, import_wild_card> {};

struct import_declaration
    : pegtl::seq<key_import, seps, pegtl::sor<import_clause, import_wild_card_clause>> {};

/**
 *
 */
struct initializer_predication : pegtl::one<'='> {};
struct initializer_content : expression {};
struct initializer : pegtl::seq<initializer_predication, seps, initializer_content> {
    using predication = initializer_predication;
    using content = initializer_content;
};

struct pattern_initializer : pegtl::seq<pattern, initializer> {};

/**
 * GRAMMAR OF A CONSTANT DECLARATION
 */
struct constant_declaration : pegtl::seq<key_const, seps, pattern_initializer> {};
struct constant_group_content
    : pegtl::list_tail<pattern_initializer, value_separator, sep> {};

/**
 * GRAMMAR OF A CONSTANT GROUP DECLARATION
 */
struct constant_group_declaration : pegtl::seq<key_const,
                                               seps,
                                               pegtl::one<'{'>,
                                               seps,
                                               constant_group_content,
                                               seps,
                                               pegtl::one<'}'>> {};

/**
 * GRAMMAR OF A VARIABLE DECLARATION
 */
struct variable_explicit_declaration_content {};

struct variable_declaration
    : pegtl::sor<pegtl::seq<identifier, seps, type_annotation, seps, initializer>,
                 pegtl::seq<key_var,
                            seps,
                            identifier,
                            seps,
                            pegtl::opt<type_annotation>,
                            seps,
                            initializer>> {};

/**
 * GRAMMAR OF A FUNCTION DECLARATION
 */
struct function_parameter
    : pegtl::seq<identifier,
                 seps,
                 type_annotation,
                 pre_pad_opt<pegtl::sor<initializer, pegtl_string_t("...")>, sep>> {};

struct function_parameters : pegtl::list<function_parameter, pegtl::one<','>, sep> {};
struct function_parameters_begin : pegtl::one<'('> {};
struct function_parameters_end : pegtl::one<')'> {};
using function_name = pegtl::sor<identifier, binary_operator, prefix_operator>;
struct function_declaration_clause
    : pegtl::seq<function_name,
                 pegtl::pad_opt<generic_parameter_clause, sep>,
                 function_parameters_begin,
                 pegtl::pad_opt<function_parameters, sep>,
                 function_parameters_end,
                 seps,
                 pegtl_string_t("->"),
                 seps,
                 type,
                 pre_pad_opt<code_block, sep>> {
    using name = function_name;
    using parameter = function_parameter;
    using parameters = function_parameters;
};

struct type_construction_declaration
    : pegtl::seq<type_identifier,
                 seps,
                 function_parameters_begin,
                 pegtl::pad_opt<function_parameters, sep>,
                 function_parameters_end,
                 pre_pad_opt<code_block, sep>> {};

struct function_declaration : pegtl::seq<key_func, seps, function_declaration_clause> {};

/**
 * GRAMMAR OF AN ENUMERATION DECLARATION
 */
struct enum_member : pegtl::seq<pegtl::opt<document, document_separator>,
                                identifier,
                                pegtl::opt<inline_seps, attributes, attribute_separator>,
                                pre_pad_opt<initializer, sep>,
                                pegtl::opt<blanks, following_document>> {};

struct enum_members_begin : pegtl::one<'{'> {};
struct enum_members_end : pegtl::one<'}'> {};
struct enum_members : list<enum_member, ',', '}'> {};
struct enum_declaration_clause : pegtl::seq<type_name,
                                            pegtl::pad_opt<type_inheritance_clause, sep>,
                                            enum_members_begin,
                                            seps,
                                            enum_members,
                                            seps,
                                            enum_members_end> {
    using name = type_name;
    using member = enum_member;
    using members_begin = enum_members_begin;
    using members_end = enum_members_end;
    using members = enum_members;
};

struct enum_declaration : pegtl::seq<key_enum, seps, enum_declaration_clause> {};

struct type_declaration;

struct type_member_declaration : pegtl::seq<identifier,
                                            seps,
                                            type_annotation,
                                            pre_pad_opt<initializer, sep>,
                                            pegtl::opt<blanks, following_document>> {};

struct type_member_group_declaration;
struct type_dynamic_member_declaration {};

struct type_field
    : pegtl::seq<pegtl::opt<document, document_separator>, type_member_declaration> {};

struct type_member
    : pegtl::seq<pegtl::opt<document, document_separator>,
                 pegtl::sor<type_member_declaration,
                            pegtl::seq<pegtl::opt<attributes, attribute_separator>,
                                       pegtl::sor<type_declaration,
                                                  enum_declaration,
                                                  type_member_group_declaration>>>> {
    using member = type_member_declaration;
    using group_member = type_member_group_declaration;
    using inner_type = type_declaration;
    using inner_enum = enum_declaration;
};

struct type_alias_predication : pegtl::one<'='> {};
struct type_alias_clause : pegtl::seq<type_alias_predication, seps, data_type> {
    using predication = type_alias_predication;
};

struct type_members : list<type_member, ',', '}'> {};

struct type_members_begin : pegtl::one<'{'> {};
struct type_members_end : pegtl::one<'}'> {};
struct type_members_clause
    : pegtl::seq<type_members_begin, seps, type_members, seps, type_members_end> {
    using member = type_member;
    using members_begin = type_members_begin;
    using members_end = type_members_end;
    using members = type_members;
};

struct type_declaration_clause
    : pegtl::seq<type_name,
                 pegtl::pad_opt<generic_parameter_clause, sep>,
                 seps,
                 pegtl::sor<type_alias_clause,
                            pegtl::seq<type_inheritance_clause,
                                       pegtl::opt<pegtl::seq<seps, type_members_clause>>>,
                            type_members_clause>> {
};

struct type_member_group_members : list<type_field, ',', '}'> {};
struct type_member_group_members_begin : pegtl::one<'{'> {};
struct type_member_group_members_end : pegtl::one<'}'> {};
struct type_member_group_declaration : pegtl::seq<type_member_group_members_begin,
                                                  seps,
                                                  type_member_group_members,
                                                  seps,
                                                  type_member_group_members_end> {
    using member = type_field;
    using members_begin = type_member_group_members_begin;
    using members_end = type_member_group_members_end;
    using members = type_member_group_members;
};

/**
 * GRAMMAR OF A TYPE DECLARATION
 */
struct type_declaration : pegtl::seq<key_type, seps, type_declaration_clause> {};

struct method_declaration
    : pegtl::seq<pegtl::pad_opt<key_func, sep>, function_declaration_clause> {};

struct interface_member : pegtl::seq<pegtl::opt<document, document_separator>,
                                   pegtl::opt<attributes, attribute_separator>,
                                   method_declaration> {};

struct interface_members : list<interface_member, ',', '}'> {
    using member = method_declaration;
};
struct interface_members_begin : pegtl::one<'{'> {};
struct interface_members_end : pegtl::one<'}'> {};

struct interface_declaration_clause
    : pegtl::seq<type_name,
                 pegtl::pad_opt<type_inheritance_clause, sep>,
                 interface_members_begin,
                 seps,
                 interface_members,
                 seps,
                 interface_members_end> {
    using member = interface_members::member;
    using members_begin = interface_members_begin;
    using members_end = interface_members;
    using members = interface_members;
};

/**
 * GRAMMAR OF A INTERFACE DECLARATION
 */
struct interface_declaration : pegtl::seq<key_interface, seps, interface_declaration_clause> {};

struct attribute_declaration_clause
    : pegtl::seq<identifier, seps, type_annotation, pre_pad_opt<initializer, sep>> {};

/**
 * GRAMMAR OF A ATTRIBUTE DECLARATION
 */
struct attribute_declaration
    : pegtl::seq<key_attribute, seps, attribute_declaration_clause> {};

/**
 * GRAMMAR OF A DECLARATION
 */
struct declaration
    : pegtl::sor<
          import_declaration,
          constant_group_declaration,
          pegtl::seq<pegtl::opt<document, document_separator>,
                     pegtl::sor<pegtl::seq<pegtl::opt<attributes, attribute_separator>,
                                           pegtl::sor<type_declaration,
                                                      function_declaration,
                                                      enum_declaration,
                                                      interface_declaration,
                                                      type_construction_declaration>>,
                                package_declaration,
                                constant_declaration,
                                variable_declaration,
                                attribute_declaration>>> {};

struct declarations : pegtl::plus<declaration> {};
}  // namespace grammar
}  // namespace mojo

#endif  // MOJO_GRAMMAR_DECLARATION_HPP