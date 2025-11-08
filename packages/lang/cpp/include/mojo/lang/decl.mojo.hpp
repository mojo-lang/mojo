#ifndef MOJO_LANG_DECL_MOJO_HPP
#define MOJO_LANG_DECL_MOJO_HPP

#include <mojo/core/string.hpp>
#include <mojo/lang/package_ptr.hpp>
#include <mojo/lang/source_file_ptr.hpp>

namespace mojo {
namespace lang {

struct Decl {
    /// position of first character belonging to the Expr
    //start_pos: Pos @1

    /// position of first character immediately after the Expr
    //end_pos  : Pos @2

    /**
     *
     */
    String document; //Document @3

    /**
     *
     */
    PackagePtr package;

    /**
     *
     */
    SourceFilePtr sourceFile;
};

}
}

#endif //MOJO_LANG_DECL_MOJO_HPP
