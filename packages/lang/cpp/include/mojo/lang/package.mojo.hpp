#ifndef MOJO_LANG_PACKAGE_MOJO_HPP
#define MOJO_LANG_PACKAGE_MOJO_HPP

#include <vector>
#include <mojo/core/array.hpp>
#include <mojo/core/string.hpp>
#include <mojo/lang/package_ptr.hpp>
#include <mojo/lang/source_file_ptr.hpp>
#include <mojo/lang/struct_type_builder.hpp>

namespace mojo {
namespace lang {

struct Package {
    String name;
    String fullName;

    PackagePtr parent;
    Array<PackagePtr> children;

    Array<SourceFilePtr> sourceFiles;
};

}  // namespace lang
}  // namespace mojo

#endif  // MOJO_LANG_PACKAGE_MOJO_HPP
