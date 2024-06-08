package domain

import (
	"bytes"
	"time"
)

type MimeType string

const (
	MimeJPG  MimeType = "image/jpeg"
	MimePNG  MimeType = "image/png"
	MimeWEBP MimeType = "image/webp"
	MimeAVI  MimeType = "video/avi"
	MimeMP4  MimeType = "video/mp4"
	MimeMOV  MimeType = "video/mov"
)

type FileCachePolicy int

const (
	FileCachePolicyNever FileCachePolicy = iota
	FileCachePolicyMinimal
	FileCachePolicyMedium
	FileCachePolicyMaximal
)

type File struct {
	Key         string
	OwnerID     string
	CachePolicy FileCachePolicy
	Mime        MimeType
	CreatedAt   time.Time
	data        []byte
}

func (f *File) Close() error {
	return nil
}

func (f *File) Write(p []byte) (n int, err error) {
	m := len(f.data)
	n = len(p)
	d := make([]byte, m+n)
	copy(d[:m], f.data)
	copy(d[m:], p)
	f.data = d
	return
}

func (f *File) Size() int {
	return len(f.data)
}

func (f *File) Bytes() []byte {
	return f.data
}

func (f *File) Buffer() *bytes.Buffer {
	return bytes.NewBuffer(f.data)
}

func (f *File) SetData(b []byte) {
	f.data = b
}
