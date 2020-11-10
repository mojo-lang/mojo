from conans import ConanFile, CMake

class NcraftConan(ConanFile):
   settings = "os", "compiler", "build_type", "arch"
   generators = "cmake"
   requires = "catch2/2.2.1@bincrafters/stable","double-conversion/3.0.0@bincrafters/stable"

   def imports(self):
      self.copy("*.dll", dst="bin", src="bin") # From bin to bin
      self.copy("*.dylib*", dst="bin", src="lib") # From lib to bin

   def build(self):
      cmake = CMake(self)
      cmake.configure()
      cmake.build()
