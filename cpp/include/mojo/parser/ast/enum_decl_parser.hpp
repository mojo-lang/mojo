#ifndef MOJO_ENUM_DECL_PARSER_HPP
#define MOJO_ENUM_DECL_PARSER_HPP

#include <mojo/lang/enum_decl.mojo.hpp>
#include <mojo/lang/term.mojo.hpp>

namespace mojo {
namespace parser {
namespace ast {

/**
 *
 */
class EnumDeclParser {
public:
    lang::EnumDeclPtr operator()(const lang::TermPtr& term, lang::PackagePtr& package) {
    }
};

}
}
}

#endif //MOJO_ENUM_DECL_PARSER_HPP
