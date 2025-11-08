/// 
type File {
    type Info {
        name: String @1     //< base name of the file
        suffix: String @2   //< suffix of the file name
        size: Int64  @3     //< length in bytes for regular files; system-dependent for others
        
        change_time: Timestamp @5 //< 文件的元数据发生变化。比如权限，所有者等
        modify_time: Timestamp @6 //< 文件内容被修改的最后时间
    }

    enum Mode {
        unspecified @0
        
        dir  @10
        // ModeSticky                                     // t: sticky
        // ModeAppend                                     // a: append-only
        // ModeExclusive                                  // l: exclusive use
        // ModeSetuid                                     // u: setuid
        // ModeSetgid                                     // g: setgid

        // ModeNamedPipe                                  // p: named pipe (FIFO)
        // ModeSocket                                     // S: Unix domain socket
        // ModeTemporary                                  // T: temporary file; Plan 9 only
        // ModeSymlink                                    // L: symbolic link
        // ModeDevice                                     // D: device file
        // ModeCharDevice                                 // c: Unix character device, when ModeDevice is set
        // Dir                                            // d: is a directory
        // ModeIrregular                                  // ?: non-regular file; nothing else is known about this file

        // // Mask for the type bits. For regular files, none will be set.
        // ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice | ModeCharDevice | ModeIrregular

        // ModePerm FileMode = 0777 // Unix permission bits
    }

    name: String @1 //< the name of the file, usually it is the full path of the file
    is_dir: Bool @2 //< whether is the directory or not

    mode: Mode @5
    info: Info @6

    files: [File] @15 //< sub files in this file if is a directory
}
