#ifndef MOJO_PARSER_ERRORS_HPP
#define MOJO_PARSER_ERRORS_HPP

#include <list>
#include <map>
#include <mojo/lang/term.hpp>
#include <pegtl/trace.hh>
#include <pegtl/normal.hh>

namespace mojo {
namespace parser {

template <typename Rule>
struct errors : public pegtl::normal<Rule> {
    static const std::string error_message;

    template <typename Input, typename... States>
    static void raise(const Input& in, States&&...) {
        throw pegtl::parse_error(error_message, in);
    }
};

template <typename Rule>
const std::string errors<Rule>::error_message = "error";

}
}

#endif  // MOJO_PARSER_ERRORS_HPP
