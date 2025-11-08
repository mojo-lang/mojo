#ifndef LANG_TYPE_ALIAS_DECL_PTR_HPP
#define LANG_TYPE_ALIAS_DECL_PTR_HPP

#include <memory>

namespace mojo {
namespace lang {

struct TypeAliasDecl;
using TypeAliasDeclPtr = std::shared_ptr<TypeAliasDecl>;

}  // namespace lang
}  // namespace mojo

#endif  // LANG_TYPE_ALIAS_DECL_PTR_HPP
