#ifndef MOJO_LANG_CALLABLE_TYPE_HPP
#define MOJO_LANG_CALLABLE_TYPE_HPP

#include <mojo/core/any.hpp>
#include <ncraft/concurrent/future.hpp>

namespace mojo {
namespace lang {

class Callable {
public:
    virtual ~Callable() = default;

    /**
     *
     * @param req
     * @return
     */
    virtual ncraft::Future<Any> call(AnyRef req) const = 0;
};

}
}

#endif //MOJO_LANG_CALLABLE_TYPE_HPP
