#ifndef MOJO_PARSER_PARSER_HPP
#define MOJO_PARSER_PARSER_HPP

#include <mojo/parser/attributes.hpp>
#include <mojo/parser/declaration/attribute_declaration.hpp>
#include <mojo/parser/declaration/const_declaration.hpp>
#include <mojo/parser/declaration/declaration.hpp>
#include <mojo/parser/declaration/enum_declaration.hpp>
#include <mojo/parser/declaration/function_declaration.hpp>
#include <mojo/parser/declaration/import_declaration.hpp>
#include <mojo/parser/declaration/initializer.hpp>
#include <mojo/parser/declaration/package_declaration.hpp>
#include <mojo/parser/declaration/service_declaration.hpp>
#include <mojo/parser/declaration/type_construction_declaration.hpp>
#include <mojo/parser/declaration/type_declaration.hpp>
#include <mojo/parser/declaration/variable_declaration.hpp>
#include <mojo/parser/document.hpp>
#include <mojo/parser/expression/array_literal.hpp>
#include <mojo/parser/expression/binary_expression.hpp>
#include <mojo/parser/expression/identifer.hpp>
#include <mojo/parser/expression/lambda_expression.hpp>
#include <mojo/parser/expression/literal.hpp>
#include <mojo/parser/expression/match_expression.hpp>
#include <mojo/parser/expression/object_literal.hpp>
#include <mojo/parser/expression/object_type_construction.hpp>
#include <mojo/parser/expression/parenthesized_expression.hpp>
#include <mojo/parser/expression/postfix_expression.hpp>
#include <mojo/parser/expression/prefix_expression.hpp>
#include <mojo/parser/expression/string_prefix_literal.hpp>
#include <mojo/parser/expression/wildcard_expression.hpp>
#include <mojo/parser/generics.hpp>
#include <mojo/parser/pattern.hpp>
#include <mojo/parser/statement/break_statement.hpp>
#include <mojo/parser/statement/continue_statement.hpp>
#include <mojo/parser/statement/for_in_statement.hpp>
#include <mojo/parser/statement/if_statement.hpp>
#include <mojo/parser/statement/match_statement.hpp>
#include <mojo/parser/statement/repeat_while_statement.hpp>
#include <mojo/parser/statement/return_statement.hpp>
#include <mojo/parser/statement/statements.hpp>
#include <mojo/parser/statement/while_statement.hpp>
#include <mojo/parser/types.hpp>

#endif  // MOJO_PARSER_PARSER_HPP
