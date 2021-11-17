#ifndef MOJO_PARSER_TERM_STATE_HPP
#define MOJO_PARSER_TERM_STATE_HPP

#include <list>
#include <map>
#include <mojo/lang/term.mojo.hpp>

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
    array_term_state(std::string&& name) : array(make_term(std::move(name))) {
    }

    array_term_state(const std::string& name) : array(make_term(name)) {
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

struct map_term_state : term_state {
    TermPtr key;
    TermPtr map;
    std::string pair_name = "field";

    map_term_state() = default;
    map_term_state(const std::string& name) : map(make_term(name)) {
    }
    map_term_state(const std::string& name, const std::string& pair_name)
        : map(make_term(name)), pair_name(pair_name) {
    }

    void add_key() {
        key = std::move(term);
    }

    void set_value() {
        auto field = make_term(pair_name);
        field->terms.push_back(std::move(key));
        field->terms.push_back(std::move(term));

        map->terms.push_back(std::move(field));

        key.reset();
        term.reset();
    }

    void success(term_state& in) {
        if (key) {
            set_value();
        }

        in.term = std::move(map);
    }
};

template <typename Rule>
struct array_action : pegtl::nothing<Rule> {};

template <char Char>
struct array_action<pegtl::one<Char>> {
    template <typename Input, typename State>
    static void apply(const Input&, State& state) {
        state.push_element();
    }
};

template <char Char, char AtChar>
struct array_action<grammar::separator<Char, AtChar>> {
    template <typename Input, typename State>
    static void apply(const Input&, State& state) {
        state.push_element();
    }
};

}  // namespace parser
}  // namespace mojo

#endif  // MOJO_PARSER_TERM_STATE_HPP
