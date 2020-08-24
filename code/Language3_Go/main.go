package main
 
import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)
 // Represents 2D field of cells 
type Field struct {
	board[][]bool
	width, height int
}

//Set new field with width and height 
func NewField(width, height int) Field {
	s := make([][]bool, height)
	for i := range s {
		s[i] = make([]bool, width)
	}
	return Field{board: s, width: width, height: height}
}

//Set the state b at location x,y on the board
func (f Field) Set(x, y int, b bool) {
	f.board[y][x] = b
}

// Return next state of the cell x, y at next time 
func (f Field) Next(x, y int) bool {
	alive := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if f.State(x+i, y+j) && !(j == 0 && i == 0) {
				alive++
			}
		}
	}
	return alive == 3 || alive == 2 && f.State(x, y)
}
 
func (f Field) State(x, y int) bool {
	for y < 0 {
		y += f.height
	}
	for x < 0 {
		x += f.width
	}
	return f.board[y%f.height][x%f.width]
}
 
 //Stores state of a round 
type Life struct {
	w, h int
	a, b Field
}
 // Returns a new game with with random initial state 
func NewLife(w, h int) *Life {
	a := NewField(w, h)
	for i := 0; i < (w * h / 2); i++ {
		a.Set(rand.Intn(w), rand.Intn(h), true)
	}
	return &Life{
		a: a,
		b: NewField(w, h),
		w: w, h: h,
	}
}
 
// Advances game by one instant
func (l *Life) Step() {
	for y := 0; y < l.h; y++ {
		for x := 0; x < l.w; x++ {
			l.b.Set(x, y, l.a.Next(x, y))
		}
	}
	l.a, l.b = l.b, l.a
}

 // Returns game board as a string
func (l *Life) String() string {
	var buf bytes.Buffer
	for y := 0; y < l.h; y++ {
		for x := 0; x < l.w; x++ {
			b := byte(' ')
			if l.a.State(x, y) {
				b = '*'
			}
			buf.WriteByte(b)
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}
 
func main() {
	l := NewLife(80, 15)
	for i := 0; i < 300; i++ {
		l.Step()
		fmt.Print("\x0c") // Clear the screen
		fmt.Println(l)
		time.Sleep(time.Second / 30)
	}
}
