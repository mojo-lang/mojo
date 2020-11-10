#ifndef MOJO_VALUE_HPP
#define MOJO_VALUE_HPP

namespace mojo {

// A value is anything RQL can pass around -- a datum, a sequence, a function, a
// selection, whatever.
class Value {
public:
    template <typename T>
            Value(T&& value);
};

using ValuePtr = std::unique_ptr<Value>;


}

#endif //MOJO_VALUE_HPP
