#ifndef NCRAFT_STRUCTURED_DATA_READER_HPP
#define NCRAFT_STRUCTURED_DATA_READER_HPP

#include <memory>
#include <sstream>
#include <vector>

namespace mojo {
namespace lang {

template <typename Self, typename Subscript = Self>
struct DataReaderOf {
    using SelfType = Self;
    using SubscriptType = Subscript;

    bool isNothing() const;

    operator bool() const;

    bool operator>>(Boolean& b) const;

    bool operator>>(Integer& n) const;

    bool operator>>(Real& r) const;

    bool operator>>(String& str) const;

    Subscript subscript(size_t idx) const;

    Subscript subscript(const String& key) const;

    bool has(const String& key) const;

    size_t length() const;

    struct iterator;

    iterator begin() const;

    iterator end() const;

    template <typename T>
    std::enable_if_t<std::is_same<T, bool>::value || !std::is_integral<T>::value,
                     Maybe<T>>
    read() const {
        T value;
        return ((*this) >> value) ? Maybe<T>(value) : Maybe<T>();
    }

    template <typename T>
    typename std::enable_if<std::is_integral<T>::value && !std::is_same<T, bool>::value,
                            Maybe<T>>::

        type
        read() const {
        Integer value;
        return ((*this) >> value) ? Maybe<T>(convert::to<T>(value)) : Maybe<T>();
    }

protected:
    DataReaderOf() {
    }

private:
    // Could be an DataReader, could be something else...
    template <typename Self_ = Self>
    auto reader() const -> decltype(std::declval<const Self_>().dataReader()) {
        return static_cast<const Self_*>(this)->dataReader();
    }
};

template <typename Self, typename Subscript>
struct DataReaderOf<Self, Subscript>::iterator {
    iterator(const iterator&) = default;

    iterator(iterator&&) = default;

    bool operator==(const iterator& other) const {
        if (enumerator_ == nullptr) {
            if (other.enumerator_ == nullptr) {
                return true;
            }
            else {
                return other.enumerator_->atEnd();
            }
        }
        else {
            if (other.enumerator_ == nullptr) {
                return enumerator_->atEnd();
            }
            else {
                return enumerator_->atEnd() == other.enumerator_->atEnd();
            }
        }
    }

    bool operator!=(const iterator& other) const {
        return !(*this == other);
    }

    const Subscript& operator*() const {
        return current_;
    }

    const Subscript* operator->() const {
        return &current_;
    }

    Maybe<String> key() const {
        return enumerator_->currentKey();
    }

    iterator& operator++() {
        enumerator_->moveNext();
        update_ptr();
        return *this;
    }

private:
    iterator(ReaderEnumeratorPtr e) : enumerator_{std::move(e)} {
        update_ptr();
    }

    ReaderEnumeratorPtr enumerator_;
    Subscript current_;
    friend struct DataReaderOf<Self, Subscript>;

    void update_ptr() {
        if (enumerator_ && !enumerator_->atEnd()) {
            auto c = enumerator_->currentValue();
            current_ = std::move(c);
        }
    }
};

template <typename Self, typename Subscript>
bool DataReaderOf<Self, Subscript>::isNothing() const {
    return reader().type() == SDataType::Nothing;
}

template <typename Self, typename Subscript>
DataReaderOf<Self, Subscript>::operator bool() const {
    return !isNothing();
}

template <typename Self, typename Subscript>
bool DataReaderOf<Self, Subscript>::operator>>(Boolean& b) const {
    if (reader().type() == SDataType::Boolean) {
        b = *reader().getBoolean();
        return true;
    }
    return false;
}

template <typename Self, typename Subscript>
bool DataReaderOf<Self, Subscript>::operator>>(Integer& n) const {
    auto type = reader().type();
    switch (type) {
        case SDataType::Integer: {
            n = *reader().getInteger();
            return true;
        }
        case SDataType::Real: {
            Real r = *reader().getReal();
            n = static_cast<Integer>(r);
            return true;
        }
        case SDataType::String: {
            std::stringstream ss(*reader().getString());
            return (ss >> n).eof();
        }
        default:
            return false;
    }
}

template <typename Self, typename Subscript>
bool DataReaderOf<Self, Subscript>::operator>>(Real& r) const {
    auto type = reader().type();
    switch (type) {
        case SDataType::Integer: {
            Integer n = *reader().getInteger();
            r = static_cast<Real>(n);
            return true;
        }
        case SDataType::Real: {
            r = *reader().getReal();
            return true;
        }
        case SDataType::String: {
            std::stringstream ss(*reader().getString());
            return (ss >> r).eof();
        }
        default:
            return false;
    }
}

template <typename Self, typename Subscript>
bool DataReaderOf<Self, Subscript>::operator>>(String& str) const {
    auto type = reader().type();
    switch (type) {
        case SDataType::Integer: {
            std::stringstream ss;
            ss << *reader().getInteger();
            str = ss.str();
            return true;
        }
        case SDataType::Real: {
            std::stringstream ss;
            ss << *reader().getReal();
            str = ss.str();
            return true;
        }
        case SDataType::String: {
            str = *reader().getString();
            return true;
        }
        default:
            return false;
    }
}

template <typename Self, typename Subscript>
typename DataReaderOf<Self, Subscript>::iterator DataReaderOf<Self, Subscript>::begin() const {
    return iterator{reader().enumerator()};
}

template <typename Self, typename Subscript>
typename DataReaderOf<Self, Subscript>::iterator DataReaderOf<Self, Subscript>::end() const {
    return iterator{ReaderEnumeratorPtr{new ReaderEnumeratorAtEnd}};
}

template <typename Self, typename Subscript>
size_t DataReaderOf<Self, Subscript>::length() const {
    return reader().length();
}

template <typename Self, typename Subscript>
Subscript DataReaderOf<Self, Subscript>::subscript(size_t idx) const {
    return reader().at(idx);
}

template <typename Self, typename Subscript>
Subscript DataReaderOf<Self, Subscript>::subscript(const String& key) const {
    return reader().get(key);
}

template <typename Self, typename Subscript>
bool DataReaderOf<Self, Subscript>::has(const String& key) const {
    return reader().has(key);
}

#if 0
template<typename T, typename Self, typename Subscript> inline
auto read(const DataReaderOf<Self, Subscript>& reader) -> decltype(std::declval<const DataReaderOf<Self, Subscript>>().read()) {
    return reader.read<T>();
}

template <typename T,
        typename R,
        typename = meta::EnableIf<std::is_base_of<DataReaderOf<typename R::SelfType, typename R::SubscriptType>, R>>> inline
auto read(const Maybe<R>& reader) -> decltype(std::declval<const R>().typename read<T>()) {
    if (reader) {
        return (*reader).typename read<T>();
    }
    else {
        return Nothing;
    }
};
#endif
}
}

#endif  // NCRAFT_STRUCTURED_DATA_READER_HPP
