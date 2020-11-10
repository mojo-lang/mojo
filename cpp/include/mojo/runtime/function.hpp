#ifndef MOJO_FUNCTION_HPP
#define MOJO_FUNCTION_HPP

#include <mojo/lang/runtime/value.hpp>
#include <mojo/lang/runtime/enviroment.hpp>

namespace mojo {

class Function {
public:
    virtual ValuePtr call(Enviroment& env) const = 0;

};

}

#endif //MOJO_FUNCTION_HPP
