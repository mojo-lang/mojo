#ifndef MOJO_RT_TERM_HPP
#define MOJO_RT_TERM_HPP

#include <mojo/lang/term.hpp>

namespace mojo {

/**
 *
 */
class RtTerm {
public:
    RtTerm(const Term& term);

    scoped_ptr_t<val_t> eval(scope_env_t *env, eval_flags_t eval_flags = NO_FLAGS) const;
    virtual deterministic_t is_deterministic() const = 0;
    virtual const char *name() const = 0;

    // Allocates a new value in the current environment.
    template<class... Args>
    scoped_ptr_t<val_t> new_val(Args... args) const {
        return make_scoped<val_t>(std::forward<Args>(args)..., backtrace());
    }
    scoped_ptr_t<val_t> new_val_bool(bool b) const {
        return new_val(datum_t::boolean(b));
    }

    virtual void accumulate_captures(var_captures_t *captures) const = 0;

    // Only terms that do a simple selection of a field, or an array of fields,
    // are allowed to be not updated with reql version. This is checked in is_acceptable_outdated
    // in sindex_manager.
    virtual bool is_simple_selector() const { return false; }

private:
    scoped_ptr_t<val_t> eval_on_current_stack(
            scope_env_t *env,
            eval_flags_t eval_flags) const;

    virtual scoped_ptr_t<val_t> term_eval(scope_env_t *env, eval_flags_t) const = 0;

private:
    const Term& term_;
};

}

/*
RtTerm root(Term(""));
auto v = root.eval(env);

auto map = v.get<Map>();
configMapnik(*map) {

}
*/

#endif //MOJO_RUNTIME_TERM_HPP
