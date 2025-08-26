package mediantwosortedarrays

/*
	Say we have the two lists:

		[1,2,3,5,6,7,11,12]
		[1,2,3,7,11,33]

		A is the shortest List

		a = [1,2,3,7,11,33]
		b = [1,2,3,5,6,7,11,12]

		totalLen = len(a) + len(b) = 6 + 8 = 14
		half = totalLen / 2 = 7

		ALeft = a[0:len(a)/2] -> [1,2,3]
		BLeft = a[0:7-(len(a)/2)] -> [1,2,3,5]

		ALeftMax = 3
		BLeftMax = 5

		ARightMin
*/

func main(){

}

func solve(a,b []float64){
	// GetMaxMinValue of A and GetMinMaxValue of A
	AMaxMin := a[len(a)/2] // -> [1,5,7,8] -> [1,5] -> 5
	AMinMax := a[len(a)/2 + 1] // -> [1,5,7,8] -> [7,8] -> 7

	BMaxMin := b[len(b)/2] // -> [2,9,10,11] -> [2,9] -> 9
	BMinMax := b[len(b)/2 + 1] // -> [2,9,10,11] -> [10,11] -> 10

	aNew = []float64

	// Determine which array part to get
	// Get Left part of A if the maximum of the left side of A is bigger that the minimum of the right side of B
	if(AMaxMin > BMinMax){
		aNew = a[0:len(a)/2] // should be [1,5]
	}

	// Get Right side of A if the maximum of the left side of B is bigger than the minimum of the right side of A
	if(BMaxMin > AMinMax){
		aNew = a[len(a) /2 + 1:len(a)]
	}


}