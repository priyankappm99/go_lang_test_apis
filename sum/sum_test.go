package main 
 import "testing"

 func TestSum(t *testing.T) {
    // Test case 1: Positive numbers
    result := Sum(5, 3)
    expected := 8

	if result != expected {
		t.Errorf("Sum(5, 3) = %d; want %d", result, expected)
	}


	//Test case 2 for -ve numbers
	result= Sum(-3,-5)
	expected=-8
	if result != expected{
		t.Errorf("Sum(-3,-5)= %d ;  want %d",result, expected)
	}
	//Test case 3: for 0 and positive number
	result= Sum(0,5)
	expected=5
	if result != expected{
		t.Errorf("Sum(0,5)= %d ;  want %d",result, expected)
	}
}