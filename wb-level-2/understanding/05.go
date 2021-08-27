package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test1() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test1()
	//fmt.Printf("%#v\t%#v\n", nil, err)   // <nil>   (*main.customError)(nil)
	// (https://medium.com/golangspec/equality-in-golang-ff44da79b7f1)
	// Неравенство будет выполнено, так как для выполнения равенства err == nil необходимо, чтобы сравниваемые
	// величины имели одинаковый тип и значение, в данном случае совпадает только значение.
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
