package main 
import "testing"
func Sum(x,y int) int{
	return x+y
}

func TestSum(t *testing.T){
	result := Sum(2,3)
	if result != 5{
		t.Errorf("Sum was incorrect, got: %d, expected: %d",result,5)
	}
}

//testing with different data sets
func TestSumManyTimes(t *testing.T){
	tables:=[]struct{
		x int
		y int
		r int 
	}{
		{1,2,3},
		{2,3,5},
		{3,4,7},
		{4,5,9},
		{0,0,0},
	}
	for _, table := range tables{
		result:= Sum(table.x, table.y)
		if result!=table.r{
			t.Errorf("Sum of (%d + %d) was incorrect, got:%d, expected:%d", table.x,table.y,result,table.r)
		}
	}
}