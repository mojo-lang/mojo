#ifndef MOJO_TYPE_INHERITANCE_PARSER_HPP
#define MOJO_TYPE_INHERITANCE_PARSER_HPP

#include <ncraft/core/noncopyable.hpp>
#include <mojo/parser/ast/nominal_type_parser.hpp>

namespace mojo {
namespace parser {
namespace ast {

/**
 *
 */
class TypeInheritanceParser : ncraft::NonCopyable {
public:
    Array<lang::NominalTypePtr> operator()() {
    }
};

}  // namespace ast
}  // namespace parser
}  // namespace mojo

#endif  // MOJO_TYPE_INHERITANCE_PARSER_HPP
