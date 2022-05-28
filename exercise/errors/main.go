package main

//func Positive(n int) (bool, error) {
//	if n == 0 {
//		return false, errors.New("Undefined")
//	}
//	return n > -1, nil
//}
//
//func Check(n int) {
//	pos, err := Positive(n)
//	if err != nil {
//		fmt.Println(n, err)
//	}
//
//	if pos {
//		fmt.Println(n, "is positive")
//	} else {
//		fmt.Println(n, "is negative")
//	}
//}

func main() {

	//err := test()
	//switch err := err.(type) {
	//case nil:
	//	// call success,nothing to do
	//case *myError:
	//	fmt.Println("error on line", err.Line)
	//default:
	//	// unknown error
	//}
}

//type myError struct {
//	Msg  string
//	File string
//	Line int
//}

//func (e *myError) Error() string {
//	return fmt.Sprintf("%s:%d: %s", e.File, e.Line, e.Msg)
//}
//
//func test() error {
//	return &myError{"something wrong happened", "server.go", 42}
//}
//
//type temporary interface {
//	Temporary() bool
//}
//
//func IsTemporary(err error) bool {
//	te, ok := err.(temporary)
//	return ok && te.Temporary()
//}

//func Countlines(r io.Reader) (int, error) {
//	var (
//		br    = bufio.NewReader(r)
//		lines int
//		err   error
//	)
//
//	for {
//		_, err = br.ReadString('\n')
//		lines++
//		if err != nil {
//			break
//		}
//	}
//
//	if err != io.EOF {
//		return 0, err
//	}
//
//	return lines, nil
//}
//
//func CountLines(r io.Reader) (int, error) {
//	sc := bufio.NewScanner(r)
//	lines := 0
//
//	for sc.Scan() {
//		lines++
//	}
//
//	return lines, sc.Err()
//}
