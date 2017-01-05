//Jared McDonald
//CSC 442

package main
import "fmt"
import "bufio"
import "os"
import "strconv"
import "strings"
import "bytes"
import "unicode"
import "math"

type DStack struct{
  top *DElement;
  size int;
}

type DElement struct {
	value float64
	next *DElement
}

func (s *DStack) push(value float64) {
  if value > 0{
  s.top = &DElement{value, s.top}
	s.size++
  }
}

func (s *DStack) isEmpty() bool {
	return s.size == 0
}
func (s *DStack) pop() (value float64) {
	if s.size > 0 {
	  value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return
}

type SStack struct{
  top *SElement;
  size int;
}

type SElement struct {
	value string
	next *SElement
}

func (s *SStack) push(value string) {
  s.top = &SElement{value, s.top}
	s.size++
}

func (s *SStack) isEmpty() bool {
	return s.size == 0
}
func (s *SStack) pop() (value string) {
	if s.size > 0 {
	  value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return
}

func main() {
    reader := bufio.NewScanner(os.Stdin)
    fmt.Printf("You can terminate the program by entering = instead of an equation.\n")
    var finished int = 0
    for finished < 1{
        fmt.Printf("Please enter an equation: ")
        reader.Scan()
        if reader.Text() != "="{
            var line string = reader.Text()
           var r float64 = calc(line)
           if r != -99.99{
            fmt.Printf("Answer: ")
            fmt.Printf("%g\n",r)
           }else{
             finished = 1
           }
        } else if reader.Text() == "="{
            finished = 1
            fmt.Println("Program terminated")
            os.Exit(0)
        }
    }
}
func calc(line string) float64{
  if line == "="{
    return -99.99
  }
  var result float64 = 0.0;
  line =strings.Replace(line,"sqrt ","|",-1)
  line =strings.Replace(line," ","",-1)
  if n, err := strconv.ParseFloat(line, 64); err == nil {
    result = n
    return result
  } else{
    var buffer bytes.Buffer
    numstack := new(DStack)
    opstack := new(SStack)
    temp := strings.Split(line,"=")
  line = temp[0]
    for _, r := range line{
      if unicode.IsDigit(rune(r)) || unicode.IsLetter(rune(r)) || string(r) == "."{
        buffer.WriteString(string(r))
      } else if string(r) == "+" || string(r) == "-" || string(r) == "/" || string(r) == "*" || string(r) == "^" || string(r) == "|"{
        if buffer.String() != "<nil>"{
          if n, err := strconv.ParseFloat(buffer.String(), 64); err == nil {
            numstack.push(n)
            buffer.Reset()
          }
        }
        if opstack.isEmpty(){
          opstack.push(string(r))
        } else{
          if !opstack.isEmpty(){
            var currentop string= opstack.pop()
            var operand float64 = numstack.pop()
            if currentop != "|"{
              var operandt float64 = numstack.pop()
              result = operate(operand,operandt,currentop)
            } else{
              result = operate(operand,0.5,currentop)
            }
            numstack.push(result)
          }
          opstack.push(string(r))
        }
      }
      numstack.push(result)
    }
    if buffer.String() != "<nil>"{
       if n, err := strconv.ParseFloat(buffer.String(), 64); err == nil {
            numstack.push(n)
            buffer.Reset()
          }
    } 
    for !opstack.isEmpty(){
      var currentop string= opstack.pop()
      var operand float64 = numstack.pop()
      if currentop != "|"{
        var operandt float64 = numstack.pop()
        result = operate(operand,operandt,currentop)
      } else{
        result = operate(operand,0.5,currentop)
      }
        numstack.push(result)
    }
    return result
  }
}
func operate(x float64, y float64, op string) float64{
  var r float64 = 0.0
  if op == "+"{
			r = y+x
	} else if op == "-"{
			r = (y-x)
  } else if op == "/"{
			r = (y/x)
	} else if op == "*"{
			r = (y*x)
	} else if op == "^"{
			r = math.Pow(y,x)
	} else if op == "|"{
			r = math.Pow(x,y)
	}
		return r;
}  
  