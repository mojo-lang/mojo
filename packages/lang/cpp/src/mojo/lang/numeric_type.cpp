#include <mojo/lang/numeric_type.hpp>

namespace mojo {
namespace lang {

#define DEFINE_NUMERIC_TYPE(T)                                              \
    const NumericType<T>* buildType(const TypeIdentifier<T>*) {             \
        static const NumericType<T>* p = new NumericType<T>{#T, "core." #T}; \
        return p;                                                           \
    }

DEFINE_NUMERIC_TYPE(Int8)
DEFINE_NUMERIC_TYPE(Int16)
DEFINE_NUMERIC_TYPE(Int32)
DEFINE_NUMERIC_TYPE(Int64)
DEFINE_NUMERIC_TYPE(UInt8)
DEFINE_NUMERIC_TYPE(UInt16)
DEFINE_NUMERIC_TYPE(UInt32)
DEFINE_NUMERIC_TYPE(UInt64)
DEFINE_NUMERIC_TYPE(Float)
DEFINE_NUMERIC_TYPE(Double)

#undef DEFINE_NUMERIC_TYPE
}
}