#ifndef LANG_DATA_DECL_HPP
#define LANG_DATA_DECL_HPP

#include <mojo/lang/type_decl.mojo.hpp>

namespace mojo {
namespace lang {

struct DataDecl : TypeDecl {
    DataType* type;
};

}
}

MOJO_BUILD_STRUCT_TYPE(mojo, lang, DataDecl) {
    field(&mojo::lang::DataDecl::type, "type", 1);
}

#endif //LANG_DATA_DECL_HPP
