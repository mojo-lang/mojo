#include <mojo/lang/string_type.hpp>

namespace mojo {
namespace lang {

const StringType* buildType(const TypeIdentifier<String>*) {
    static const StringType* p = new StringType;
    return p;
}

}
}
