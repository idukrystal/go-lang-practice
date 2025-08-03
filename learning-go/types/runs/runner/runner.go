package runner

import "fmt"

type runner interface {
	Run()
}

type Result interface { 
	print()
	addResult(bool)
}

type inbuiltResult struct {
	wasSuccess bool
}

// new method added in new version
func (ir inbuiltResult) isSuccess() bool {
	return ir.wasSuccess
}


func (ir inbuiltResult) addResult(wasSuccess bool) {
	ir.wasSuccess = wasSuccess
}

func (ir inbuiltResult) print() {
	if ir.wasSuccess {
		fmt.Println("Success ...")
	} else {
		fmt.Println("Failled !!!")
	}
}

func RunInBackground(runner runner) /*inbuiltResult*/ Result {
	var res inbuiltResult

	fmt.Println("Preparing to run")
	runner.Run()
	res.addResult(true)
	fmt.Println("All Done")
	
	// functiona caller cant use new methods of inbuilt result
	return res

}

func FinnishUp(res Result) {
	fmt.Println("done")
}
