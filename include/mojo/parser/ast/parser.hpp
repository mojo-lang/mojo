#ifndef MOJO_PARSER_HPP
#define MOJO_PARSER_HPP

#include <mojo/lang/term.mojo.hpp>

namespace mojo {
namespace parser {
namespace ast {

class Parser {
public:
    Parser(const TermPtr& term) : term_(term) {
    }

    Parser(TermPtr&& term) : term_(std::move(term)) {
    }

    bool expect(const String& type) {
        return term_ && term_->type == type;
    }

    // contain the type of term in terms of the current term
    bool contains(const String& type) {
        return !!getTerm(type);
    }

    TermPtr getTerm(const String& type) const& {
        for (auto const& term : term_->terms) {
            if (term->type == type) {
                return term;
            }
        }
        return TermPtr{};
    }

    TermPtr getTerm(const String& type) && {
        for (auto&& term : term_->terms) {
            if (term->type == type) {
                return std::move(term);
            }
        }
        return TermPtr{};
    }

    std::size_t termsCount() const {
        return term_->terms.size();
    }

    std::size_t termsCount(const String& type) const {
        auto term = getTerm(type);
        if (term) {
            return term->terms.size();
        }
        else {
            return 0;
        }
    }

private:
    TermPtr term_;
    TermPtr current_;
};

}  // namespace ast
}  // namespace parser
}  // namespace mojo

#endif  // MOJO_PARSER_HPP
