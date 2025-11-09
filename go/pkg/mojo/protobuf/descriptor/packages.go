package descriptor

import "strings"

type Packages struct {
    Files map[string][]*File

    FilesByPath    map[string]*File
    EnumsByName    map[string]*Enum
    MessagesByName map[string]*Message
    ServicesByName map[string]*Service
}

func NewPackages() *Packages {
    return &Packages{
        Files:          make(map[string][]*File),
        FilesByPath:    make(map[string]*File),
        EnumsByName:    make(map[string]*Enum),
        MessagesByName: make(map[string]*Message),
        ServicesByName: make(map[string]*Service),
    }
}

func (p *Packages) GetMessage(name string) *Message {
    if p != nil {
        if msg, ok := p.MessagesByName[name]; ok {
            return msg
        }
    }
    return nil
}

func (p *Packages) AddFile(file *File) *Packages {
    if p != nil && file != nil {
        if _, ok := p.FilesByPath[file.GetName()]; !ok {
            pkg := file.GetPackageName()
            p.Files[pkg] = append(p.Files[pkg], file)

            p.FilesByPath[file.GetName()] = file
            for _, enum := range file.Enums {
                p.EnumsByName[enum.GetFullName()] = enum
            }
            for _, message := range file.Messages {
                p.MessagesByName[message.GetFullName()] = message
            }
            for _, service := range file.Services {
                p.ServicesByName[service.GetFullName()] = service
            }
        }
    }
    return p
}

func (p *Packages) Filter(pkg string, strict bool) []*File {
    var files []*File
    if p != nil {
        equal := func(name string) bool {
            if strict {
                return name == pkg
            } else {
                return name == pkg || strings.HasPrefix(name, pkg+".")
            }
        }

        for name, fs := range p.Files {
            if equal(name) {
                files = append(files, fs...)
            }
        }
    }

    return files
}
