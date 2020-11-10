#include <chrono>
#include <iostream>
#include <mojo/mojo_test.hpp>

namespace {
using namespace mojo;

TEST_CASE("object_literal_parser", "[expression]") {
    const char* object_str =
            "{"
                    "id: 'MapEMG',"
                    "name: 'Map for EMG',"
                    "sources: [{ id: '4445', host: 'localhost'}],"
                    "layers: ["
                    "   island | category == 999 | zoom in 1..16 | attach_style({"
                    "       fill.color: rgb(244, 243, 239),"
                    "       fill.outline.color: rgb(246, 242, 239)"
                    "   })"
                    "]"
                    "}";

    simple_success_parse<grammar::expression>(object_str);
}
}
