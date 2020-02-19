

// 问题描述
/*
判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-sudoku
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 解题思路1 逐行、逐列、3*3方块遍历
// 链接：https://leetcode-cn.com/problems/valid-sudoku/solution/chang-gui-bian-li-bu-yong-qi-ta-shu-ju-jie-gou-by-/
func isValidSudoku(board [][]byte) bool {
	var target byte
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			target = board[i][j]
			if target != '.' {
				// 判断是否有效
				for k := 0; k < 9; k++ {
					// 检测行
					if k != j && board[i][k] != '.' && board[i][k] == target {
						return false
					}
					// 检测列
					if i != k && board[k][j] != '.' && board[k][j] == target {
						return false
					}
				}
				// 测试3*3方块
				for m := i / 3 * 3; m < (i/3+1)*3; m++ {
					for n := j / 3 * 3; n < (j/3+1)*3; n++ {
						if m != i && n != j && board[m][n] != '.' && board[m][n] == target {
							return false
						}
					}
				}
			}
		}
	}
	return true
}




// 解题思路2 位图
/*
使用 uint16 , 来作为位图，压缩空间复杂度。

使用了3个 [9]uint16 数组 作为 位图，使用内存固定为 3 * 9 * 2byte = 54byte，空间复杂度为 O(1)

两层嵌套for循环，时间复杂度为 O(1)

链接：https://leetcode-cn.com/problems/valid-sudoku/solution/shi-yong-uint16-wei-tu-shi-xian-cha-zhong-by-pppob/

*/
func isValidSudoku(board [][]byte) bool {
	// 行、列、块都声明为9个uint16的数组
    var row, col, block [9]uint16
    var cur uint16
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if board[i][j] == '.' {
                continue
            }
            cur = 1 << (board[i][j] - '1')  // 当前数字的 二进制数位 位置，board[i][j]中存放的是字符对应的ascii码值
            bi := i/3 + j/3*3  // 将3x3的块 划分为0~8的索引号
            if (row[i] & cur) | (col[j] & cur) | (block[bi] & cur) != 0 { // 使用与运算查重
                return false
            }
            // 在对应的位图上，加上当前数字
            row[i] |= cur
            col[j] |= cur
            block[bi] |= cur
        }
    }
    return true
}



// 解题思路3 位图 

/*
主要两个要点：
1、只遍历一次如何储存数据；
2、判断是在一个3*3的框中的方法。

1、使用了2进制的9个位数，如果是第一个数是1，那么统计标志就是0000000010(二进制 1左移1位)，如果第二个数是3那么统计标识变为0000001010(二进制 1左移3位再加上原来的)，每次判断有没有重复就右移相应位数之后整除2即可。
2、同官方解法int boxNum = i / 3 * 3 + j / 3;如果是0,1,2行的话整除3就是0，然后再加上列数整除3，这样就把整个9*9分为了编号0-8的9个3*3的区域。

链接：https://leetcode-cn.com/problems/valid-sudoku/solution/36you-xiao-de-shu-du-ti-jie-java-3ms-by-zhaomin666/
来源：力扣（LeetCode）
*/
public boolean isValidSudoku(char[][] board) {
	int[] rowCnt = new int[9];
	int[] colCnt = new int[9];
	int[] boxCnt = new int[9];
	for (int i = 0; i < 9; i++) {
		for (int j = 0; j < 9; j++) {
			if ('.' == board[i][j]) {
				continue;
			}
			int num = board[i][j] - 48;
			// 处理行
			if ((rowCnt[i] >> num) % 2 == 1) {
				return false;
			} else {
				rowCnt[i] += 1 << num;
			}
			// 处理列
			if ((colCnt[j] >> num) % 2 == 1) {
				return false;
			} else {
				colCnt[j] += 1 << num;
			}
			// 处理框
			int boxNum = i / 3 * 3 + j / 3;
			if ((boxCnt[boxNum] >> num) % 2 == 1) {
				return false;
			} else {
				boxCnt[boxNum] += 1 << num;
			}
		}
	}
	return true;
}

