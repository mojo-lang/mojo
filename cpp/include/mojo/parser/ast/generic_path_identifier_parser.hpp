#ifndef MOJO_GENERIC_PATH_IDENTIFIER_PARSER_HPP
#define MOJO_GENERIC_PATH_IDENTIFIER_PARSER_HPP

#include <mojo/core/result.hpp>
#include <mojo/core/array.hpp>
#include <mojo/core/string.hpp>
#include <mojo/lang/term.mojo.hpp>
#include <mojo/lang/nominal_type_ptr.hpp>
#include <mojo/parser/ast/parser.hpp>
#include <mojo/parser/term_types.hpp>
#include <mojo/parser/ast/path_identifier_parser.hpp>

namespace mojo {
    namespace parser {
        namespace ast {

            class GenericPathIdentifierParser {
            public:
                Result<Array<String>, Array<lang::NominalTypePtr>> operator()(const TermPtr &term) {
                    Array<String> pathIdentifier;
                    Array<lang::NominalTypePtr> genericArguments;
                    auto p = Parser(term);
                    if (p.expect(kGenericPathIdentifier)) {
                        auto identifier = PathIdentifierParser()(p.getTerm(kPathIdentifier));
                        if (identifier) {
                            auto arguments = p.getTerm(kGenericArguments);

                        }
                    } else {

                    }
                }
            };

        }  // namespace ast
    }  // namespace parser
}  // namespace mojo

#endif //MOJO_GENERIC_PATH_IDENTIFIER_PARSER_HPP
