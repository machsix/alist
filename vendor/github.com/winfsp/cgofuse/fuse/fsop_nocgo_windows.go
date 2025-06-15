//go:build !cgo && windows
// +build !cgo,windows

/*
 * fsop_nocgo_windows.go
 *
 * Copyright 2017-2022 Bill Zissimopoulos
 */
/*
 * This file is part of Cgofuse.
 *
 * It is licensed under the MIT license. The full license text can be found
 * in the License.txt file at the root of this project.
 */

package fuse

// Error codes reported by FUSE file systems.
const (
	E2BIG           = 7
	EACCES          = 13
	EADDRINUSE      = 100
	EADDRNOTAVAIL   = 101
	EAFNOSUPPORT    = 102
	EAGAIN          = 11
	EALREADY        = 103
	EBADF           = 9
	EBADMSG         = 104
	EBUSY           = 16
	ECANCELED       = 105
	ECHILD          = 10
	ECONNABORTED    = 106
	ECONNREFUSED    = 107
	ECONNRESET      = 108
	EDEADLK         = 36
	EDESTADDRREQ    = 109
	EDOM            = 33
	EEXIST          = 17
	EFAULT          = 14
	EFBIG           = 27
	EHOSTUNREACH    = 110
	EIDRM           = 111
	EILSEQ          = 42
	EINPROGRESS     = 112
	EINTR           = 4
	EINVAL          = 22
	EIO             = 5
	EISCONN         = 113
	EISDIR          = 21
	ELOOP           = 114
	EMFILE          = 24
	EMLINK          = 31
	EMSGSIZE        = 115
	ENAMETOOLONG    = 38
	ENETDOWN        = 116
	ENETRESET       = 117
	ENETUNREACH     = 118
	ENFILE          = 23
	ENOATTR         = ENODATA
	ENOBUFS         = 119
	ENODATA         = 120
	ENODEV          = 19
	ENOENT          = 2
	ENOEXEC         = 8
	ENOLCK          = 39
	ENOLINK         = 121
	ENOMEM          = 12
	ENOMSG          = 122
	ENOPROTOOPT     = 123
	ENOSPC          = 28
	ENOSR           = 124
	ENOSTR          = 125
	ENOSYS          = 40
	ENOTCONN        = 126
	ENOTDIR         = 20
	ENOTEMPTY       = 41
	ENOTRECOVERABLE = 127
	ENOTSOCK        = 128
	ENOTSUP         = 129
	ENOTTY          = 25
	ENXIO           = 6
	EOPNOTSUPP      = 130
	EOVERFLOW       = 132
	EOWNERDEAD      = 133
	EPERM           = 1
	EPIPE           = 32
	EPROTO          = 134
	EPROTONOSUPPORT = 135
	EPROTOTYPE      = 136
	ERANGE          = 34
	EROFS           = 30
	ESPIPE          = 29
	ESRCH           = 3
	ETIME           = 137
	ETIMEDOUT       = 138
	ETXTBSY         = 139
	EWOULDBLOCK     = 140
	EXDEV           = 18
)

// Flags used in FileSystemInterface.Create and FileSystemInterface.Open.
const (
	O_RDONLY  = 0x0000
	O_WRONLY  = 0x0001
	O_RDWR    = 0x0002
	O_APPEND  = 0x0008
	O_CREAT   = 0x0100
	O_TRUNC   = 0x0200
	O_EXCL    = 0x0400
	O_ACCMODE = O_RDONLY | O_WRONLY | O_RDWR
)

// File type and permission bits.
const (
	S_IFMT   = 0170000
	S_IFBLK  = 0060000
	S_IFCHR  = 0020000
	S_IFIFO  = 0010000
	S_IFREG  = 0100000
	S_IFDIR  = 0040000
	S_IFLNK  = 0120000
	S_IFSOCK = 0140000

	S_IRWXU = 00700
	S_IRUSR = 00400
	S_IWUSR = 00200
	S_IXUSR = 00100
	S_IRWXG = 00070
	S_IRGRP = 00040
	S_IWGRP = 00020
	S_IXGRP = 00010
	S_IRWXO = 00007
	S_IROTH = 00004
	S_IWOTH = 00002
	S_IXOTH = 00001
	S_ISUID = 04000
	S_ISGID = 02000
	S_ISVTX = 01000
)

// BSD file flags (Windows file attributes).
const (
	UF_HIDDEN   = 0x00008000
	UF_READONLY = 0x00001000
	UF_SYSTEM   = 0x00000080
	UF_ARCHIVE  = 0x00000800
)

// Access flags
const (
	F_OK      = 0
	R_OK      = 4
	W_OK      = 2
	X_OK      = 1
	DELETE_OK = 0x40000000 // Delete access check [Windows only]
)

// Options that control Setxattr operation.
const (
	XATTR_CREATE  = 1
	XATTR_REPLACE = 2
)

// Flags used in Utimens and Utimens3.
const (
	UTIME_NOW  = (1 << 30) - 1
	UTIME_OMIT = (1 << 30) - 2
)

// Flags used in FileSystemRename3.Rename3.
const (
	RENAME_NOREPLACE = 1 << 0
	RENAME_EXCHANGE  = 1 << 1
	RENAME_WHITEOUT  = 1 << 2
)

// Notify actions.
const (
	NOTIFY_MKDIR    = 0x0001
	NOTIFY_RMDIR    = 0x0002
	NOTIFY_CREATE   = 0x0004
	NOTIFY_UNLINK   = 0x0008
	NOTIFY_CHMOD    = 0x0010
	NOTIFY_CHOWN    = 0x0020
	NOTIFY_UTIME    = 0x0040
	NOTIFY_CHFLAGS  = 0x0080
	NOTIFY_TRUNCATE = 0x0100
)
