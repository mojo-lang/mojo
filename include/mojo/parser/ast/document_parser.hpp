#ifndef MOJO_DOCUMENT_PARSER_HPP
#define MOJO_DOCUMENT_PARSER_HPP

#include <mojo/parser/ast/parser.hpp>

namespace mojo {
namespace parser {
namespace ast {

class DocumentParser {
public:
    String operator()(const TermPtr& term);
};

}  // namespace ast
}  // namespace parser
}  // namespace mojo

#endif  // MOJO_DOCUMENT_PARSER_HPP
