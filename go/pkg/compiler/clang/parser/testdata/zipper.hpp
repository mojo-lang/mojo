//
//  zipper.h
//
//  This code is based on 'Making MiniZip Easier to Use' by John Schember.
//  https://nachtimwald.com/2019/09/08/making-minizip-easier-to-use/
//
//  Copyright (c) 2021 Yuji Hirose. All rights reserved.
//  MIT License
//

#pragma once

#include <minizip/unzip.h>
#include <minizip/zip.h>

#include <cassert>
#include <fstream>
#include <functional>
#include <string>
#include <vector>

#define ZIPPER_BUF_SIZE 8192
#define ZIPPER_MAX_NAMELEN 256

namespace zipper {

using read_cb_t = std::function<void(const char *data, size_t len)>;

class Zip {
 public:
  Zip() = default;
  Zip(const std::string &zipname) { open(zipname); }
  ~Zip() { close(); }

  bool open(const std::string &zipname) {
    zfile_ = zipOpen64(zipname.data(), 0);
    return is_open();
  }

  bool is_open() const { return zfile_ != nullptr; }

  void close() {
    if (zfile_ != nullptr) {
      zipClose(zfile_, nullptr);
      zfile_ = nullptr;
    }
  }

  bool add_dir(std::string dirname) {
    assert(zfile_ && !dirname.empty());

    if (dirname.back() != '/') {
      dirname += '/';
    }

    auto ret = zipOpenNewFileInZip64(zfile_, dirname.data(), nullptr, nullptr,
                                     0, nullptr, 0, nullptr, 0, 0, 0);

    if (ret != ZIP_OK) {
      return false;
    }

    zipCloseFileInZip(zfile_);
    return ret == ZIP_OK ? true : false;
  }

  bool add_file(const std::string &path, const char *data, size_t len) {
    assert(zfile_ && data && len > 0);

    auto ret = zipOpenNewFileInZip64(
        zfile_, path.data(), nullptr, nullptr, 0, nullptr, 0, nullptr,
        Z_DEFLATED, Z_DEFAULT_COMPRESSION, (len > 0xffffffff) ? 1 : 0);

    if (ret != ZIP_OK) {
      return false;
    }

    ret = zipWriteInFileInZip(zfile_, data, static_cast<unsigned int>(len));

    zipCloseFileInZip(zfile_);
    return ret == ZIP_OK ? true : false;
  }

  bool add_file(const std::string &path, const std::string &data) {
    return add_file(path, data.data(), data.size());
  }

  operator zipFile() { return zfile_; }

 private:
  zipFile zfile_ = nullptr;
};

class UnZip {
 public:
  UnZip() = default;
  UnZip(const std::string &zipname) { open(zipname); }
  ~UnZip() { close(); }

  bool open(const std::string &zipname) {
    uzfile_ = unzOpen64(zipname.data());
    return is_open();
  }

  bool is_open() const { return uzfile_ != nullptr; }

  void close() {
    if (uzfile_ != nullptr) {
      unzClose(uzfile_);
      uzfile_ = nullptr;
    }
  }

  bool read(read_cb_t cb) const {
    assert(uzfile_);

    auto ret = unzOpenCurrentFile(uzfile_);
    if (ret != UNZ_OK) {
      return false;
    }

    char buf[ZIPPER_BUF_SIZE];
    int red;

    while ((red = unzReadCurrentFile(uzfile_, buf, sizeof(buf))) > 0) {
      cb(buf, red);
    }

    unzCloseCurrentFile(uzfile_);
    return red >= 0;
  }

  bool read(std::string &buf) const {
    return read([&](const char *data, size_t len) { buf.append(data, len); });
  }

  std::string file_path() const {
    assert(uzfile_);

    char name[ZIPPER_MAX_NAMELEN];
    unz_file_info64 finfo;

    auto ret = unzGetCurrentFileInfo64(uzfile_, &finfo, name, sizeof(name),
                                       nullptr, 0, nullptr, 0);
    if (ret != UNZ_OK) {
      return std::string();
    }

    return name;
  }

  bool is_dir() const {
    assert(uzfile_);

    char name[ZIPPER_MAX_NAMELEN];
    unz_file_info64 finfo;

    auto ret = unzGetCurrentFileInfo64(uzfile_, &finfo, name, sizeof(name),
                                       nullptr, 0, nullptr, 0);
    if (ret != UNZ_OK) {
      return false;
    }

    auto len = strlen(name);
    return finfo.uncompressed_size == 0 && len > 0 && name[len - 1] == '/';
  }

  bool is_file() const { return !is_dir(); }

  bool next() const {
    assert(uzfile_);
    return unzGoToNextFile(uzfile_) == UNZ_OK;
  }

  uint64_t file_size() const {
    assert(uzfile_);

    unz_file_info64 finfo;
    auto ret = unzGetCurrentFileInfo64(uzfile_, &finfo, nullptr, 0, nullptr, 0,
                                       nullptr, 0);
    if (ret != UNZ_OK) {
      return 0;
    }

    return finfo.uncompressed_size;
  }

  template <typename T>
  void enumerate(T callback) {
    if (!file_path().empty()) {
      do {
        callback(*this);
      } while (next());
    }
  }

  operator unzFile() { return uzfile_; }

 private:
  unzFile uzfile_ = nullptr;
};

template <typename T>
inline bool enumerate(const std::string &zipname, T callback) {
  UnZip unzip;
  if (unzip.open(zipname)) {
    unzip.enumerate(callback);
    return true;
  }
  return false;
}

};  // namespace zipper
