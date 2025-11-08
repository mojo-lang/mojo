#ifndef LANG_PACKAGE_PTR_HPP
#define LANG_PACKAGE_PTR_HPP

#include <memory>

namespace mojo {
namespace lang {

struct Package;
using PackagePtr = std::shared_ptr<Package>;

}  // namespace lang
}  // namespace mojo

#endif  // LANG_PACKAGE_PTR_HPP
