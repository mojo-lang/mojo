#ifndef LANG_FUNC_DECL_PTR_HPP
#define LANG_FUNC_DECL_PTR_HPP

#include <memory>

namespace mojo {
namespace lang {

struct FuncDecl;
using FuncDeclPtr = std::shared_ptr<FuncDecl>;

}  // namespace lang
}  // namespace mojo

#endif  // LANG_FUNC_DECL_PTR_HPP
