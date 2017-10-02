package problem1

import "errors"

func Solve(max int) (int, error) {
	if(max < 0){
		return 0, errors.New("greater than zero please")
	}

	sum := 0

	for index:=1; index < max; index++  {
		if(index % 3 == 0){
			sum += index
		} else if(index % 5 == 0){
			sum += index
		}
	}

 return sum, nil
}