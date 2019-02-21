package lager

import (
	"io"
	"sync"
	"bytes"
	"github.com/supeanut/logtool/lager/color"
)

type Sink interface {
	Log(level LogLevel, payload []byte)
}

type writerSink struct {
	writer io.Writer
	minLogLevel LogLevel
	name string
	writeL *sync.Mutex
}

func (sink *writerSink) Log(level LogLevel, log []byte) {
	if level < sink.minLogLevel {
		return
	}
	if sink.name == "stdout" {
		if bytes.Contains(log, []byte("WARN")) {
			log = bytes.Replace(log, []byte("WARN"), color.WarnByte, -1)
		}else if bytes.Contains(log, []byte("ERROR")) {
			log = bytes.Replace(log, []byte("ERROR"), color.ErrorByte, -1)
		} else if bytes.Contains(log, []byte("FATAL")) {
			log = bytes.Replace(log, []byte("FATAL"), color.FatalByte, -1)
		} else if bytes.Contains(log, []byte("INFO")) {
			log = bytes.Replace(log, []byte("INFO"), color.InfoByte, -1)
		}
	}
	log = append(log, '\n')
	sink.writeL.Lock()
	sink.writer.Write(log)
	sink.writeL.Unlock()
}

func NewWriterSink(name string, writer io.Writer, minLogLevel LogLevel) Sink {
	return &writerSink{
		writer: writer,
		minLogLevel: minLogLevel,
		writeL: new(sync.Mutex),
		name: name,
	}
}
