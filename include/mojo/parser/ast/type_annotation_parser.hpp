#ifndef MOJO_NOMINAL_TYPE_PARSER_HPP
#define MOJO_NOMINAL_TYPE_PARSER_HPP

#include <mojo/lang/nominal_type.mojo.hpp>
#include <mojo/lang/package.mojo.hpp>
#include <mojo/parser/ast/parser.hpp>

namespace mojo {
namespace parser {
namespace ast {

class TypeAnnotationParser {
public:
    lang::NominalTypePtr operator()(const TermPtr& term, lang::PackagePtr package) {
    }
};

}  // namespace ast
}  // namespace parser
}  // namespace mojo

#endif  // MOJO_NOMINAL_TYPE_PARSER_HPP
