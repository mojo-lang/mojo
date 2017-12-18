#ifndef MOJO_PARSER_DOCUMENT_HPP
#define MOJO_PARSER_DOCUMENT_HPP

#include <mojo/grammar/grammar.hpp>
#include <mojo/parser/control.hpp>
#include <mojo/parser/term_state.hpp>

namespace mojo {
namespace parser {

template<bool Following = false>
struct document_state : term_state {
    std::string content;

    void set_line(std::string&& line) {
        content = std::move(line);
    }

    void success(term_state& in) {
        auto term = make_term(Following ? "following_document" : "document");
        term->value = std::move(content);
        in.term = term;
    }
};

struct documents_state : term_state {
    TermPtr document;

    documents_state() {
        document = make_term("document");
    }

    void add_line(bool new_line = true) {
        if (term) {
            document->value.append(std::move(term->value));
            if (new_line) {
                document->value.append("\n");
            }
            term.reset();
        }
        else {
            // should log warning for push a null element
        }
    }

    void success(term_state& in) {
        if (this->term) {
            add_line(false);
        }
        in.term = std::move(document);
    }
};

template <typename Rule>
struct document_action : pegtl::nothing<Rule> {};

template <>
struct document_action<grammar::document_line_begin> {
    template <typename Input>
    static void apply(const Input& input, documents_state& state) {
        state.add_line();
    }
};

template <>
struct document_action<grammar::document_string> {
    template <typename Input>
    static void apply(const Input& input, document_state<>& state) {
        state.set_line(input.string());
    }
};

template <>
struct document_action<grammar::following_document_string> {
    template <typename Input>
    static void apply(const Input& input, document_state<true>& state) {
        state.set_line(input.string());
    }
};

template <>
struct control<grammar::document>
    : pegtl::change_state_and_action<grammar::document,
                                     documents_state,
                                     document_action,
                                     errors> {};

template <>
struct control<grammar::document_string>
    : pegtl::change_state_and_action<grammar::document_string,
                                     document_state<>,
                                     document_action,
                                     errors> {};

template <>
struct control<grammar::following_document_string>
    : pegtl::change_state_and_action<grammar::following_document_string,
                                     document_state<true>,
                                     document_action,
                                     errors> {};
}
}

#endif  // MOJO_PARSER_DOCUMENT_HPP
