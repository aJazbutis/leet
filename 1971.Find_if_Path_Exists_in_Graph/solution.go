/*
** https://www.hackerearth.com/practice/notes/disjoint-set-union-union-find/
*/

// func find(arr []int, a, b int) bool {
// 	return arr[a] == arr[b]
// }

func find(arr []int, a, b int) bool {
	return root(arr, a) == root(arr, b)
}

func root(arr []int, i int) int {
	for arr[i] != i {
		i = arr[i]
	}
	return i
}

// func union(arr []int, a, b int) {
// 	tmp := arr[a]
// 	for i := range arr {
// 		if arr[i] == tmp {
// 			arr[i] = arr[b]
// 		}
// 	}
// }

func union(arr []int, a, b int) {
	rootA := root(arr, a)
	rootB := root(arr, b)
	arr[rootA] = rootB
}

func validPath(n int, edges [][]int, source int, destination int) bool {
	if source == destination {
		return true
	}
	//disjointSet
	ds := make([]int, n)
	for i := range ds {
		ds[i] = i
	}
	for _, v := range edges {
		union(ds, v[0], v[1])
	}
	return find(ds, source, destination)
}
