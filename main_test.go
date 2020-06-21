package main
import (
"testing"
)

func TestCalculate(t *testing.T){
	
	var tests = []struct{
	input int
	expected int 
	}{
		{2,7},
		{5,10},
		{10,15},
	}

for _,test := range tests{
	if output:=Calculate(test.input) ; output!=test.expected{
		t.Error("Test Falied  {} Input {} Expected {} Output",test.input,test.expected,output)
	}
}	
}

/* CMD

C:\Users\DELL\Desktop\New folder>go test
PASS
ok      _/C_/Users/DELL/Desktop/New_folder      0.046s

C:\Users\DELL\Desktop\New folder>go test
--- FAIL: TestCalculate (0.00s)
    main_test.go:19: Test Falied  {} Input {} Expected {} Output 2 7 6
    main_test.go:19: Test Falied  {} Input {} Expected {} Output 5 10 9
    main_test.go:19: Test Falied  {} Input {} Expected {} Output 10 15 14
FAIL
exit status 1
FAIL    _/C_/Users/DELL/Desktop/New_folder      0.076s


Verbose Test Output
Exactly how and what time test case are running by –v command

C:\Users\DELL\Desktop\New folder>go test -v
=== RUN   TestCalculate
--- PASS: TestCalculate (0.00s)
PASS
ok      _/C_/Users/DELL/Desktop/New_folder      0.029s

Checking Test Coverage

C:\Users\DELL\Desktop\New folder>go test -cover
PASS
coverage: 50.0% of statements
ok      _/C_/Users/DELL/Desktop/New_folder      0.038s

C:\Users\DELL\Desktop\New folder>go test -coverprofile=coverage.out
PASS
coverage: 50.0% of statements
ok      _/C_/Users/DELL/Desktop/New_folder      0.026s

C:\Users\DELL\Desktop\New folder>go tool cover -html=coverage.out


*/
