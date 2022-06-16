// Convert C++ code to Go:
// #include <iostream>
// using namespace std;
//
// int main() {
//   int foo [5] = {6, 2, 77, 4, 12};
//   int count=0;
//   int sum=0;
//   while (count<5){
//     sum+=foo[count];
//     count++;
//   }
//   cout<<"The sum is: "<<sum<<endl;
//   return 0;
// }

package main

func GetSum(array []int) int {
	count := 0
	sum := 0

	for count < len(array) {
		sum += array[count]
		count += 1
	}

	return sum
}
