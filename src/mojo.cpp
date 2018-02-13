#include <iostream>
#include <fstream>
#include <pegtl.hh>
#include <pegtl/trace.hh>
#include <pegtl/analyze.hh>
#include <mojo/parser/parser.hpp>

int main(int argc, char* argv[]) {
    mojo::parser::term_state state;

    bool trace = false;
    std::string file_name;
    if (argc == 2) {
        file_name = argv[1];
    }
    else if (argc == 3) {
        trace = true;
        file_name = argv[2];
    }

    if (trace) {
        std::ifstream file;
        file.open(file_name.c_str());
        std::string name;
        pegtl::trace_istream<mojo::grammar::grammar>(file, "mojo", 4096, name);
    }
    else {
        pegtl::parse_file<mojo::grammar::grammar, pegtl::nothing, mojo::parser::control>(file_name, state);
        std::cout << *state.term << std::endl;
    }

    return 0;
}
