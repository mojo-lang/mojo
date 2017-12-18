#ifndef MOJO_PARSER_CONTROL_HPP
#define MOJO_PARSER_CONTROL_HPP

#include <pegtl/contrib/changes.hh>
#include <mojo/parser/errors.hpp>

namespace mojo {
namespace parser {

template <typename Rule>
struct control : errors<Rule> {};

}
}

#endif //MOJO_PARSER_CONTROL_HPP
