#ifndef LANG_STRUCT_DECL_PTR_HPP
#define LANG_STRUCT_DECL_PTR_HPP

#include <memory>

namespace mojo {
namespace lang {

struct StructDecl;
using StructDeclPtr = std::shared_ptr<StructDecl>;

}
}

#endif //LANG_STRUCT_DECL_PTR_HPP
