#ifndef MOJO_PARSER_TERM_STATE_HPP
#define MOJO_PARSER_TERM_STATE_HPP

#include <list>
#include <map>
#include <mojo/lang/term.hpp>

namespace mojo {
namespace parser {

struct term_state {
    term_state() = default;

    term_state(const term_state&) = delete;
    void operator=(const term_state&) = delete;

    TermPtr term;
};

struct array_term_state : term_state {
    TermPtr array;

    array_term_state() = default;
    array_term_state(const char* name) : array(make_term(name)) {
    }

    void push_element() {
        if (this->term) {
            array->terms.push_back(std::move(term));
            term.reset();
        }
        else {
            // should log warning for push a null element
        }
    }

    void success(term_state& in) {
        if (this->term) {
            push_element();
        }
        in.term = std::move(array);
    }
};

struct dictionary_term_state : term_state {
    TermPtr key;
    TermPtr dictionary;
    const char* field_name = "field";

    dictionary_term_state() = default;
    dictionary_term_state(const char* name) : dictionary(make_term(name)) {
    }
    dictionary_term_state(const char* name, const char* field_name)
            : dictionary(make_term(name)), field_name(field_name) {
    }

    void add_key() {
        key = std::move(term);
    }

    void set_value() {
        auto field = make_term(field_name);
        field->terms.push_back(std::move(key));
        field->terms.push_back(std::move(term));

        dictionary->terms.push_back(std::move(field));

        key.reset();
        term.reset();
    }

    void success(term_state& in) {
        if (key) {
            set_value();
        }

        in.term = std::move(dictionary);
    }
};

template<typename Rule>
struct array_action : pegtl::nothing<Rule> {
};

template<char Char>
struct array_action<pegtl::one<Char>> {
    template<typename Input, typename State>
    static void apply(const Input&, State &state) {
        state.push_element();
    }
};

template<char Char>
struct array_action<grammar::separator<Char>> {
    template<typename Input, typename State>
    static void apply(const Input&, State &state) {
        state.push_element();
    }
};

}
}

#endif  // MOJO_PARSER_TERM_STATE_HPP
