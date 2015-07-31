package main

import (
	"crypto/sha512"
	"io"
	"math/rand"
	"time"
)

type FiniteByteReader struct {
	Source []byte
	Times  int
}

func (f *FiniteByteReader) Read(p []byte) (n int, err error) {
	if f.Times == 0 {
		return 0, io.EOF
	}
	length := len(p)

	if length <= len(f.Source) {
		// write one instance
		for i := 0; i < length; i++ {
			p[i] = f.Source[i]
		}
		f.Times--
		return length, nil

	} else {
		// write multiple instances
		written := 0
		for written+len(f.Source) < length {
			for i := written; i < written+len(f.Source); i++ {
				p[i] = f.Source[i-written]
			}
			written += len(f.Source)
			f.Times--

			if f.Times <= 0 {
				return written, nil
			}
		}
		return written, nil
	}

}
func sha512Hash() {
	f := &FiniteByteReader{
		Source: []byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."),
		Times:  rand.Intn(1000),
	}
	s := sha512.New()
	io.Copy(s, f)
	s.Sum([]byte{})
}

func doCPUWork() {
	t := time.After(*loadTime)
	for {
		select {
		case <-t:
			// mark as no longer under load
			loadMutex.Lock()
			underLoad = false
			loadMutex.Unlock()

			return
		default:
			for i := 0; i < 50; i++ {
				sha512Hash()
			}
		}
	}
}
