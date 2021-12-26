package paasio

import "io"

// Define readCounter and writeCounter types here.

type readCounter struct {
	io.Reader
	lock chan bool
	opCount int
	bytesRead int64
}

type writeCounter struct {
	io.Writer
	lock chan bool
	opCount int
	bytesWritten int64
}

type readWriteCounter struct {
	readCounter
	writeCounter
}

// For the return of the function NewReadWriteCounter, you must also define a type that satisfies the ReadWriteCounter interface.

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{
		Writer: writer, 
		lock: make(chan bool, 1),
		opCount: 0,
		bytesWritten: 0,
	}
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{
		Reader: reader, 
		lock: make(chan bool, 1),
		opCount: 0,
		bytesRead: 0,
	}
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{
		readCounter: readCounter{
			Reader: readwriter,
			lock: make(chan bool, 1),
			opCount: 0,
			bytesRead: 0,
		},
		writeCounter: writeCounter{
			Writer: readwriter,
			lock: make(chan bool, 1),
			opCount: 0,
			bytesWritten: 0,
		},
	}
}

func (rc *readCounter) Read(p []byte) (int, error) {

	bytesRead, err := rc.Reader.Read(p)

	if err != nil {
		return 0, err
	}

	rc.lock <- true
	rc.opCount++
	rc.bytesRead += int64(bytesRead)
	<- rc.lock 

	return bytesRead, nil
}

func (rc *readCounter) ReadCount() (int64, int) {
	var bytesRead int64
	var opCount int

	rc.lock <- true
	bytesRead = rc.bytesRead
	opCount = rc.opCount
	<- rc.lock

	return bytesRead, opCount
}

func (wc *writeCounter) Write(p []byte) (int, error) {

	bytesWritten, err := wc.Writer.Write(p)

	if err != nil {
		return 0, err
	}

	wc.lock <- true
	wc.opCount++
	wc.bytesWritten += int64(bytesWritten)
	<-wc.lock

	return bytesWritten, nil
}

func (wc *writeCounter) WriteCount() (int64, int) {

	var bytesWritten int64
	var opCount int

	wc.lock <- true
	bytesWritten = wc.bytesWritten
	opCount = wc.opCount
	<- wc.lock

	return bytesWritten, opCount
}
