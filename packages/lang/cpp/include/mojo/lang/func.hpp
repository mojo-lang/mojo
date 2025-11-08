#ifndef MOJO_FUNCTION_HPP
#define MOJO_FUNCTION_HPP

#include <mojo/core/string.hpp>
#include <mojo/lang/callable.hpp>

namespace mojo {
namespace lang {

class Function : Callable {
public:
    String name;

    virtual void call(T&, ObjectVisitor& visitor) const = 0;
};

}
}

#endif //MOJO_FUNCTION_HPP
