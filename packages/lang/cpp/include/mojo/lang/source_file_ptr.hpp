#ifndef LANG_SOURCE_FILE_PTR_HPP
#define LANG_SOURCE_FILE_PTR_HPP

#include <memory>

namespace mojo {
namespace lang {

struct SourceFile;
using SourceFilePtr = std::shared_ptr<SourceFile>;

}  // namespace lang
}  // namespace mojo

#endif  // LANG_SOURCE_FILE_PTR_HPP
