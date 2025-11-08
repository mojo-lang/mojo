#ifndef MOJO_CORE_NUMERIC_HPP
#define MOJO_CORE_NUMERIC_HPP

#include <cstdint>
#include <cstdlib>

namespace mojo {
namespace core {

using Int8 = ::std::int8_t;
using Int16 = ::std::int16_t;
using Int32 = ::std::int32_t;
using Int64 = ::std::int64_t;

using UInt8 = ::std::uint8_t;
using UInt16 = ::std::uint16_t;
using UInt32 = ::std::uint32_t;
using UInt64 = ::std::uint64_t;

using Float32 = float;
using Float64 = double;

using Float = Float32;
using Double = Float64;

using Int = int;
using UInt = unsigned int;
using Size = ::std::size_t;

}

using Int8 = ::mojo::core::Int8;
using Int16 = ::mojo::core::Int16;
using Int32 = ::mojo::core::Int32;
using Int64 = ::mojo::core::Int64;

using UInt8 = ::mojo::core::UInt8;
using UInt16 = ::mojo::core::UInt16;
using UInt32 = ::mojo::core::UInt32;
using UInt64 = ::mojo::core::UInt64;

using Int = ::mojo::core::Int;
using UInt = ::mojo::core::UInt;
using Size = ::mojo::core::Size;

using Float32 = ::mojo::core::Float32;
using Float64 = ::mojo::core::Float64;

using Float = ::mojo::core::Float;
using Double = ::mojo::core::Double;

}

#endif //MOJO_CORE_NUMERIC_HPP
