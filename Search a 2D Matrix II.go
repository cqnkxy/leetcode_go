func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 {
        return false
    }
    row, col := len(matrix), len(matrix[0])
    for i, j := 0, col-1; i < row && j >= 0; {
        if target == matrix[i][j] {
            return true
        } else if target > matrix[i][j] {
            i += 1
        } else {
            j -= 1
        }
    }
    return false
}
