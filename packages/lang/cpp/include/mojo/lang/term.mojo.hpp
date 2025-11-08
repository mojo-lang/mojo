#ifndef MOJO_TERM_HPP
#define MOJO_TERM_HPP

#include <memory>
#include <string>
#include <vector>
#include <ostream>
#include <mojo/core/array.hpp>
#include <mojo/core/string.hpp>

namespace mojo {
namespace lang {

struct Term;
using TermPtr = ::std::shared_ptr<Term>;

struct Term {
    String type;
    String value;
    // Either<bool, int64_t, double, String> value;

    //Position position;

    TermPtr parent;
    Array<TermPtr> terms;
};

inline
String indent_str(int indent) {
    return String(indent * 2, ' ');
}

inline
void print(::std::ostream &o, const Term &term, int indent = 0) {
    o << "{\n";
    o << indent_str(indent) << "\"type\":\"" << term.type << "\"";
    if (!term.value.empty()) {
        o << ",\n" << indent_str(indent) << "\"value\":\"" << term.value << "\"\n";
    } else if (!term.terms.empty()) {
        o << ",\n" << indent_str(indent) << "\"terms\": [";
        print(o, *term.terms[0], indent + 1);
        for (::std::size_t i = 1; i < term.terms.size(); ++i) {
            o << ",";
            print(o, *term.terms[i], indent + 1);
        }

        o << "]\n";
    } else {
        o << "\n";
    }

    o << indent_str(indent) << "}";
}

inline ::std::ostream &operator<<(::std::ostream &o, const Term &term) {
    print(o, term);
    return o;
}

inline TermPtr make_term(const String &type) {
    auto ptr = ::std::make_shared<Term>();
    ptr->type = type;
    return ::std::move(ptr);
}

inline TermPtr make_term(String &&type) {
    auto ptr = ::std::make_shared<Term>();
    ptr->type = std::move(type);
    return ::std::move(ptr);
}

inline TermPtr make_term(const String &type, const String &value) {
    auto ptr = ::std::make_shared<Term>();
    ptr->type = type;
    ptr->value = value;
    return ::std::move(ptr);
}

inline TermPtr make_term(const String &type, ::std::initializer_list<TermPtr> terms) {
    auto ptr = make_term(type);
    for (auto &t : terms) {
        ptr->terms.push_back(::std::move(t));
    }
    return ::std::move(ptr);
}

inline bool operator==(const Term &l, const Term &r) {
    if (l.type == r.type && l.value == r.value && l.terms.size() == r.terms.size()) {
        for (::std::size_t i = 0; i < l.terms.size(); ++i) {
            if (!(*l.terms[i] == *r.terms[i])) {
                return false;
            }
        }
        return true;
    } else {
        return false;
    }
}
}

using Term = lang::Term;
using TermPtr = lang::TermPtr;
using lang::make_term;

}

#endif  // MOJO_TERM_HPP
