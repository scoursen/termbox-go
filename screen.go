// +build !windows

package termbox

import "bytes"
import "errors"
import "fmt"
import "os"

type screen struct {
	out            *os.File
	in             int
	orig_tios      syscall_Termios
	back_buffer    cellbuf
	front_buffer   cellbuf
	termw          int
	termh          int
	input_mode     InputMode
	output_mode    OutputMode
	lastfg         Attribute
	lastbg         Attribute
	lastx          int
	lasty          int
	cursor_x       int
	cursor_y       int
	foreground     Attribute
	background     Attribute
	inbuf          []byte
	outbuf         bytes.Buffer
	input_comm     chan input_event
	interrupt_comm chan struct{}
	intbuf         []byte
}

func SetTerm(s *screen) (*screen, error) {
	if _, ok := screens[s]; ok {
		if s != nil {
			if current != nil {
				copy_globals_to(current)
			}
			old := current
			current = s
			copy_globals_from(current)
			return old, nil
		} else {
			return nil, errors.New("nil terminal")
		}
	} else {
		return nil, errors.New(fmt.Sprintf("%v not valid terminal", s))
	}
}

func DelTerm(s *screen) error {
	if _, ok := screens[s]; ok {
		delete(screens, s)
		return nil
	} else {
		return errors.New(fmt.Sprintf("%v not valid terminal", s))
	}
}

func copy_globals_to(s *screen) {
	s.orig_tios = orig_tios
	s.back_buffer = back_buffer
	s.front_buffer = front_buffer
	s.termw = termw
	s.termh = termh
	s.input_mode = input_mode
	s.output_mode = output_mode
	s.lastfg = lastfg
	s.lastbg = lastbg
	s.lastx = lastx
	s.lasty = lasty
	s.cursor_x = cursor_x
	s.cursor_y = cursor_y
	s.foreground = foreground
	s.background = background
	s.inbuf = inbuf
	s.outbuf = outbuf
	s.input_comm = input_comm
	s.interrupt_comm = interrupt_comm
	s.intbuf = intbuf
}

func copy_globals_from(s *screen) {
	orig_tios = s.orig_tios
	back_buffer = s.back_buffer
	front_buffer = s.front_buffer
	termw = s.termw
	termh = s.termh
	input_mode = s.input_mode
	output_mode = s.output_mode
	lastfg = s.lastfg
	lastbg = s.lastbg
	lastx = s.lastx
	lasty = s.lasty
	cursor_x = s.cursor_x
	cursor_y = s.cursor_y
	foreground = s.foreground
	background = s.background
	inbuf = s.inbuf
	outbuf = s.outbuf
	input_comm = s.input_comm
	interrupt_comm = s.interrupt_comm
	intbuf = s.intbuf
}
