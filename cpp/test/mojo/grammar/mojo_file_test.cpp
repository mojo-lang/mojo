#include <boost/filesystem.hpp>
#include <pegtl.hh>
#include <pegtl/trace.hh>
#include <mojo/grammar/grammar.hpp>
#include <mojo/parser/parser.hpp>
#include <catch.hpp>

std::string file_content(const std::string& filename) {
    std::ifstream ifs(filename);
    return std::string(std::istreambuf_iterator<char>(ifs), std::istreambuf_iterator<char>());
}

void get_files(const boost::filesystem::path& dir_path, std::vector<std::string> files) {
    if (!boost::filesystem::exists(dir_path)) {
        return;
    }

    boost::filesystem::directory_iterator end_itr;
    for (boost::filesystem::directory_iterator itr(dir_path); itr != end_itr; ++itr) {
        if (boost::filesystem::is_directory(itr->status())) {
            get_files(itr->path(), files);
        }
        else if (itr->path().extension() == "mojo") {
            files.push_back(itr->path().string());
        }
    }
}

template <typename Rule, template <typename> class Action>
inline void file_success_check(const std::string& filename) {
    std::string out;
    bool result = pegtl::parse_file<Rule, Action>(filename, out);
    REQUIRE((result && out == file_content(filename)));
}

template <typename Rule>
struct action : pegtl::nothing<Rule> {};

template <>
struct action<mojo::grammar::statements> {
    static void apply(const pegtl::action_input& in, std::string& name) {
        name = in.string();
    }
};

TEST_CASE("mojo_file_tests", "[mojo]") {
    std::vector<std::string> files;
    get_files("./test/mojo_file", files);
    for (auto const& file : files) {
        file_success_check<mojo::grammar::grammar, action>(file);
    }
}
