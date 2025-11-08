#ifndef NCRAFT_SDATA_WRITER_HPP
#define NCRAFT_SDATA_WRITER_HPP

#include <ncraft/sdata/reader.hpp>
#include <ncraft/sdata/types.hpp>

namespace mojo {
namespace lang {

struct Spectator;

template <typename Self, typename Subscript = Self>
struct DataWriterOf {
    bool operator<<(NothingType);

    bool operator<<(Boolean b);

    bool operator<<(Integer n);

    bool operator<<(Real r);

    bool operator<<(String str);

    // Convenience:
    bool operator<<(int n) {
        return *this << (Integer)n;
    }

    bool operator<<(const char* str) {
        return *this << std::string{str};
    }

    Subscript mutableSubscript(size_t idx);

    Subscript mutableSubscript(const String& key);

    bool erase(const String& key);

    template <typename T>
    bool write(const T& value);

protected:
    Writer() {
    }

private:
    // Could be an IWriter, could be something else...
    template <typename Self_ = Self>
    auto writer() -> decltype(std::declval<Self_>().writerInterface()) {
        return static_cast<Self_*>(this)->writerInterface();
    }
};

template <typename Self, typename Subscript>
bool Writer<Self, Subscript>::operator<<(NothingType) {
    return writer().setNothing();
}

template <typename Self, typename Subscript>
bool Writer<Self, Subscript>::operator<<(Boolean b) {
    return writer().setBoolean(b);
}

template <typename Self, typename Subscript>
bool Writer<Self, Subscript>::operator<<(Integer n) {
    return writer().setInteger(n);
}

template <typename Self, typename Subscript>
bool Writer<Self, Subscript>::operator<<(Real r) {
    return writer().setReal(r);
}

template <typename Self, typename Subscript>
bool Writer<Self, Subscript>::operator<<(String str) {
    return writer().setString(std::move(str));
}

template <typename Self, typename Subscript>
Subscript Writer<Self, Subscript>::mutableSubscript(size_t idx) {
    return writer().mutableAtIndex(idx);
}

template <typename Self, typename Subscript>
Subscript Writer<Self, Subscript>::mutableSubscript(const String& key) {
    return writer().mutableAtKey(key);
}

template <typename Self, typename Subscript>
bool Writer<Self, Subscript>::erase(const String& key) {
    return writer().erase(key);
}

template <typename Self, typename Subscript>
template <typename T>
bool Writer<Self, Subscript>::write(const T& value) {
    auto pushed = writer().append();
    return Self{pushed} << value;
}
}
}

#endif  // NCRAFT_SDATA_WRITER_HPP
