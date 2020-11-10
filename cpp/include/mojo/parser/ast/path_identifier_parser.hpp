#ifndef MOJO_PATH_IDENTIFIER_PARSER_HPP
#define MOJO_PATH_IDENTIFIER_PARSER_HPP

#include <mojo/core/result.hpp>
#include <mojo/core/array.hpp>
#include <mojo/core/string.hpp>
#include <mojo/lang/term.mojo.hpp>
#include <mojo/parser/ast/parser.hpp>
#include <mojo/parser/term_types.hpp>

namespace mojo {
    namespace parser {
        namespace ast {

            class PathIdentifierParser {
            public:
                Result<Array<String>> operator()(const TermPtr &term, lang::PackagePtr &package) {
                    Array<String> pathIdentifier;
                    auto p = Parser(term);
                    if (p.expect(kPathIdentifier)) {
                        for (auto &t : p.terms()) {
                            if (t->type == kIdentifier) {
                                pathIdentifier.push_back(t->value);
                            } else {

                            }
                        }
                    } else {

                    }

                    return std::move(pathIdentifier);
                }
            };

        }  // namespace ast
    }  // namespace parser
}  // namespace mojo

#endif //MOJO_PATH_IDENTIFIER_PARSER_HPP
