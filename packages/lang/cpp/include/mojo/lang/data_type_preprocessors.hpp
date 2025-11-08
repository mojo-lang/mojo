#ifndef MOJO_LANG_DATA_TYPE_PREPROCESSORS_HPP
#define MOJO_LANG_DATA_TYPE_PREPROCESSORS_HPP

#include <boost/preprocessor/seq/cat.hpp>
#include <boost/preprocessor/seq/fold_left.hpp>
#include <boost/preprocessor/seq/reverse.hpp>
#include <boost/preprocessor/seq/seq.hpp>
#include <boost/preprocessor/stringize.hpp>
#include <boost/preprocessor/variadic/to_seq.hpp>

#define MOJO_CAT(s, state, x) BOOST_PP_CAT(state, BOOST_PP_CAT(_, x))
#define MOJO_STR_CAT(s, state, x) state "." BOOST_PP_STRINGIZE(x)
#define MOJO_COMPOSE(s, state, x) state::x

#define MOJO_TYPES(...) BOOST_PP_VARIADIC_TO_SEQ(__VA_ARGS__)

#define MOJO_TYPES_CAT(types) \
    BOOST_PP_SEQ_FOLD_LEFT(MOJO_CAT, BOOST_PP_SEQ_HEAD(types), BOOST_PP_SEQ_TAIL(types))

#define MOJO_TYPES_COMPOSE(types)                    \
    BOOST_PP_SEQ_FOLD_LEFT(MOJO_COMPOSE,             \
                           BOOST_PP_SEQ_HEAD(types), \
                           BOOST_PP_SEQ_TAIL(types))

#define MOJO_TYPES_FULL_NAME(types)                  \
    BOOST_PP_SEQ_FOLD_LEFT(MOJO_STR_CAT,             \
                           BOOST_PP_STRINGIZE(BOOST_PP_SEQ_HEAD(types)), \
                           BOOST_PP_SEQ_TAIL(types))

#define MOJO_TYPES_LAST_NAME(types) \
    BOOST_PP_STRINGIZE(BOOST_PP_SEQ_HEAD(BOOST_PP_SEQ_REVERSE(types)))

#endif //MOJO_LANG_DATA_TYPE_PREPROCESSORS_HPP
