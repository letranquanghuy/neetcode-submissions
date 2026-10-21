class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        shape = len(board)
        for i in range(shape):
            row_seen = set()
            col_seen = set()
            for j in range(shape):
                row_value = board[i][j]
                col_value = board[j][i]
                if row_value != "." and row_value in row_seen:
                    return False
                if col_value != "." and col_value in col_seen:
                    return False
                row_seen.add(row_value)
                col_seen.add(col_value)

        # square
        square_shape = shape // 3
        for i in range(square_shape):
            for j in range(square_shape):
                seen = set()
                for m in range(i*3,i*3+3):
                    for n in range(j*3,j*3+3):
                        value = board[m][n]
                        if value != "." and value in seen:
                            return False
                        seen.add(value)
        
        return True