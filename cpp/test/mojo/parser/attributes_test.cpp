#include <chrono>
#include <iostream>
#include <mojo/mojo_test.hpp>

namespace {
using namespace mojo;

TEST_CASE("attributes_parser", "[attributes]") {
    simple_success_parse<grammar::attributes>(
        "@1",
        make_term("attributes",
                  {make_term("attribute", {make_term("integer_literal", "1")})}));

    simple_success_parse<grammar::attributes>(
        "@1@str",
        make_term("attributes",
                  {make_term("attribute", {make_term("integer_literal", "1")}),
                   make_term("attribute",
                             {make_term("generic_path_identifier",
                                        {make_term("path_identifier",
                                                   {make_term("identifier",
                                                              "str")})})})}));

    simple_success_parse<grammar::attributes>(
        "@1@obj({a:2})",
        make_term(
            "attributes",
            {make_term("attribute", {make_term("integer_literal", "1")}),
             make_term(
                 "attribute",
                 {make_term("path_identifier", {make_term("identifier", "obj")}),
                  make_term(
                      "object_literal",
                      {make_term("object_field",
                                 {make_term("generic_path_identifier",
                                            {make_term("path_identifier",
                                                       {make_term("identifier", "a")}),
                                             make_term("integer_literal", "2")})})})})}));

    simple_success_parse<grammar::attributes>(
        "@1  @str('str', '23') \n@34",
        make_term("attributes",
                  {make_term("attribute", {make_term("integer_literal", "1")}),
                   make_term("attribute",
                             {make_term("path_identifier",
                                        {make_term("identifier", "str")}),
                              make_term("string_literal", "str")}),
                   make_term("attribute", {make_term("integer_literal", "34")})}));

    simple_success_parse<grammar::attributes>(
        "@1@{foo:bar, fb: 'foobar'}",
        make_term("attributes",
                  {make_term("attribute", {make_term("integer_literal", "1")}),
                   make_term("group_attribute",
                             {make_term("attribute",
                                        {make_term("path_identifier",
                                                   {make_term("identifier", "foo")}),
                                         make_term("identifier", "bar")}),
                              make_term("attribute",
                                        {make_term("path_identifier",
                                                   {make_term("identifier", "fb")}),
                                         make_term("string_literal", "foobar")})})}));
}
}  // namespace
