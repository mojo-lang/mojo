#ifndef LANG_ENUM_DECL_PTR_HPP
#define LANG_ENUM_DECL_PTR_HPP

#include <memory>

namespace mojo {
namespace lang {

struct EnumDecl;
using EnumDeclPtr = std::shared_ptr<EnumDecl>;

}
}

#endif //LANG_ENUM_DECL_PTR_HPP
