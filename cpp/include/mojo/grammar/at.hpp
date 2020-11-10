#ifndef MOJO_GRAMMAR_AT_HPP
#define MOJO_GRAMMAR_AT_HPP

#include <pegtl/analysis/generic.hh>
#include <pegtl/internal/rule_conjunction.hh>
#include <pegtl/internal/skip_control.hh>
#include <pegtl/internal/trivial.hh>

namespace mojo {
namespace grammar {

// at which is enable action

template <typename... Rules>
struct at;

template <>
struct at<> : pegtl::internal::trivial<true> {};

template <typename... Rules>
struct at {
    using analyze_t = pegtl::analysis::generic<pegtl::analysis::rule_type::OPT, Rules...>;

    template <pegtl::apply_mode,
              template <typename...> class Action,
              template <typename...> class Control,
              typename Input,
              typename... States>
    static bool match(Input& in, States&&... st) {
        auto m = in.mark();
        return pegtl::internal::rule_conjunction<
            Rules...>::template match<pegtl::apply_mode::ACTION, Action, Control>(in, st...);
    }
};
}
}

namespace pegtl {
namespace internal {

template <typename... Rules>
struct skip_control<mojo::grammar::at<Rules...>> : std::false_type {};
}
}

#endif  // MOJO_GRAMMAR_AT_HPP
