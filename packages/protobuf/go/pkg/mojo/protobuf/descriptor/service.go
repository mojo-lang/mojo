// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// [source](https://github.com/protocolbuffers/protobuf-go//compiler/protogen/protogen.go)

package descriptor

import (
    "google.golang.org/protobuf/types/descriptorpb"
    "strings"
)

// Service describes an service.
type Service struct {
    Descriptor
    Proto *descriptorpb.ServiceDescriptorProto

    FullName string

    Methods []*Method // service method declarations
}

// Return a slice of all the EnumDescriptors defined within this File
//func WrapServiceDescriptors(file *FileDescriptor) []*Service {
//    sl := make([]*Service, 0, len(file.Service)+10)
//    // Top-level Enums.
//    for i, service := range file.Service {
//        sl = append(sl, NewService(service, nil, file, i))
//    }
//    return sl
//}

// NewService Construct the Service
func NewService(file *File) *Service {
    return &Service{
        Descriptor: Descriptor{File: file},
        Proto:      &descriptorpb.ServiceDescriptorProto{},
    }

    //sd.Path = fmt.Sprintf("%d,%d", ServiceTypeIndex, index)
    //return sd
}

func NewServiceFrom(file *File, proto *descriptorpb.ServiceDescriptorProto) *Service {
    service := &Service{
        Descriptor: Descriptor{File: file},
        Proto:      proto,
    }

    for _, method := range proto.Method {
        service.AppendMethod(NewMethodFrom(service, method))
    }

    return service
}

func (s *Service) proto() *descriptorpb.ServiceDescriptorProto {
    if s != nil {
        return s.Proto
    }
    return nil
}

func (s *Service) GetName() string {
    return s.proto().GetName()
}

func (s *Service) GetFullName() string {
    if s != nil {
        if len(s.FullName) == 0 {
            s.FullName = strings.Join([]string{s.GetPackageName(), s.GetName()}, ".")
        }
        return s.FullName
    }
    return ""
}

func (s *Service) GetPackageName() string {
    if s != nil && s.File != nil {
        return s.File.GetPackageName()
    }
    return ""
}

func (s *Service) SetName(name string) *Service {
    if s != nil && s.Proto != nil {
        s.Proto.Name = &name
        s.FullName = concatFullName(s.File.GetPackageName(), name)
    }
    return s
}

func (s *Service) HasMethod() bool {
    return s != nil && len(s.Methods) > 0
}

func (s *Service) GetMethod(name string) *Method {
    for _, m := range s.Methods {
        if m.GetName() == name {
            return m
        }
    }
    return nil
}

func (s *Service) IsMethodExist(name string) bool {
    return s.GetMethod(name) != nil
}

func (s *Service) AppendMethod(method *Method) *Service {
    if s != nil && s.Proto != nil {
        s.Methods = append(s.Methods, method)
        s.Proto.Method = append(s.Proto.Method, method.Proto)
    }
    return s
}
