#ifndef MOJO_VALUE_DECL_PARSER_HPP
#define MOJO_VALUE_DECL_PARSER_HPP

#include <memory>
#include <ncraft/core/noncopyable.hpp>
#include <mojo/parser/ast/document_parser.hpp>
#include <mojo/parser/ast/parser.hpp>
#include <mojo/parser/ast/type_annotation_parser.hpp>
#include <mojo/parser/term_types.hpp>

namespace mojo {
namespace parser {
namespace ast {

template <typename T, typename Enable = void>
class ValueDeclParser : ncraft::NonCopyable {
public:
    T operator()(const TermPtr& term, lang::PackagePtr package) {
        auto p = Parser(term);
        auto ptr = std::make_shared<T>();

        ptr->document = DocumentParser{}();
        ptr->type = TypeAnnotationParser{}(p.getTerm(kTypeAnnotation), package);

        return ptr;
    }
};

}  // namespace ast
}  // namespace parser
}  // namespace mojo

#endif  // MOJO_VALUE_DECL_PARSER_HPP
