package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type Matrix struct {
	num_rows uint64
	num_cols uint64
	mat      []int64
}

func input_num() int64 {
	var guess int64
	_, err := fmt.Scanf("%d", &guess)
	if err != nil {
		fmt.Printf("Error!!")
		os.Exit(3)
	}
	return guess
}

func input_len() uint64 {
	var guess uint64
	_, err := fmt.Scanf("%d", &guess)
	if err != nil {
		fmt.Printf("Error!!")
		os.Exit(3)
	}
	return guess
}

func get_ind(row uint64, col uint64, num_rows uint64) uint64 {
	return ((num_rows * col) + row)
}

func take_matrix_input() Matrix {
	fmt.Printf("Input the number of rows (m): \n")
	var num_rows uint64 = input_len()
	fmt.Printf("Input the number of cols (n): \n")
	var num_cols uint64 = input_len()
	fmt.Printf("Enter %d numbers: \n", num_rows*num_cols)
	var matrix = make([]int64, (num_rows * num_cols))
	for cur_row := uint64(0); cur_row < num_rows; cur_row++ {
		for cur_col := uint64(0); cur_col < num_cols; cur_col++ {
			matrix[get_ind(cur_row, cur_col, num_rows)] = input_num()
		}
	}
	var m = Matrix{num_rows, num_cols, matrix}
	//   MatrixMap m = MatrixMap(num_rows, num_cols, max_threads);
	//   m.take_stdin_matrix();
	//   cout        << "You just entered: " << endl;
	// m.print_matrix();
	return m

}
func multiply_row_and_col(a Matrix, cur_row uint64, b Matrix, cur_col uint64) int64 {
	var val int64 = 0
	var cc uint64 = get_ind(0, cur_col, b.num_rows)
	var last_ind uint64 = get_ind(cur_row, a.num_cols-1, a.num_rows)
	var cur_ind uint64 = get_ind(cur_row, 0, a.num_rows)
	for cur_ind <= last_ind {
		val += a.mat[cur_ind] * b.mat[cc]
		cc += 1
		cur_ind += a.num_rows
	}
	return val
}

func print_matrix(m Matrix) {
	for row := uint64(0); row < m.num_rows; row++ {
		for col := uint64(0); col < m.num_cols; col++ {
			fmt.Printf("%d \t", m.mat[get_ind(row, col, m.num_rows)])
		}
		fmt.Printf("\n")
	}
}
func single_calculation(m1 Matrix, cur_row uint64, m2 Matrix, cur_col uint64, num_rows uint64, mat []int64) {
	var v int64 = multiply_row_and_col(m1, cur_row, m2, cur_col)
	mat[get_ind(cur_row, cur_col, num_rows)] = v
}

func multiply(m1 Matrix, m2 Matrix) Matrix {
	var sz uint64 = m1.num_rows * m2.num_cols
	var num_rows uint64 = m1.num_rows
	var num_cols uint64 = m2.num_cols
	var mat = make([]int64, sz)
	for cur_row := uint64(0); cur_row < num_rows; cur_row++ {
		for cur_col := uint64(0); cur_col < num_cols; cur_col++ {
			single_calculation(m1, cur_row, m2, cur_col, num_rows, mat)
		}
	}
	var res Matrix = Matrix{m1.num_rows, m2.num_cols, mat}
	return res
}

func multiply_row(wg *sync.WaitGroup, m1 Matrix, m2 Matrix, cur_row uint64, num_rows uint64, num_cols uint64, mat []int64) {
	defer wg.Done()
	for cur_col := uint64(0); cur_col < num_cols; cur_col++ {
		single_calculation(m1, cur_row, m2, cur_col, num_rows, mat)
	}
}

func multiply_multi(m1 Matrix, m2 Matrix) Matrix {
	var wg sync.WaitGroup
	var sz uint64 = m1.num_rows * m2.num_cols
	var num_rows uint64 = m1.num_rows
	var num_cols uint64 = m2.num_cols
	var mat = make([]int64, sz)
	for cur_row := uint64(0); cur_row < num_rows; cur_row++ {
		wg.Add(1)
		go multiply_row(&wg, m1, m2, cur_row, num_rows, num_cols, mat)
	}
	wg.Wait()
	var res Matrix = Matrix{m1.num_rows, m2.num_cols, mat}
	return res
}

func parse_cmd_args() bool {
	var ret bool = false
	argsWithoutProg := os.Args[1:]
	for cur_col := 0; cur_col < len(argsWithoutProg); cur_col++ {
		if argsWithoutProg[cur_col] == "-s" {
			ret = true
		}
	}
	return ret
}
func main() {
	var isSingleThread = parse_cmd_args()
	var m1 Matrix = take_matrix_input()
	var m2 Matrix = take_matrix_input()
	if m1.num_cols != m2.num_rows {
		fmt.Printf("Matrix multiplication not possible")
		return
	}
	var start = time.Now()
	var m3 Matrix
	if isSingleThread {
		m3 = multiply(m1, m2)
	} else {
		m3 = multiply_multi(m1, m2)
	}
	fmt.Printf("Time taken by multiplication: %d microseconds\n", int64(time.Since(start)/1000))
	fmt.Printf("Product is: \n")
	print_matrix(m3)
	// else {
	// 		auto start = chrono::high_resolution_clock::now();
	// 		MatrixMap mu = MatrixMap();
	// 		MatrixMap *mup = &mu;
	// 		MatrixMap::multiply(m, b, mup);
	// 		auto stop = chrono::high_resolution_clock::now();
	// 		auto duration = chrono::duration_cast<chrono::microseconds>(stop - start);
	// 		printf("Time taken by multiplication: %lld microseconds\n", duration.count());
	// 		cout << "Product is: " << endl;
	// 		mu.print_matrix();
	// }
}
