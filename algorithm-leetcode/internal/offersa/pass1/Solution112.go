package pass1
/**
 * 剑指 Offer II 112. 最长递增路径
 * 输入，m x n 整数矩阵 matrix，
 * 对于每个单元格，你可以往上，下，左，右四个方向移动
 * 输出，最长递增路径 的长度
 */

 	/**
     * 记忆深度优先搜索
     * 从矩阵的每个点出发 深度优先向四周拓展 寻找最大路径，使用矩阵记录中间过程
     * @param matrix
     * @return
     */
func longestIncreasingPath(matrix [][]int) int {
	//
	if matrix==nil{
		return 0
	}
	if len(matrix)==0{
		return 0
	}
	if len(matrix[0])==0{
		return 0
	}
	// 上，下，左，右
	var steps [4][2]int=[4][2]int{{1,0},{-1,0},{0,-1},{0,1}}
	var maxr int=0
	// dfs
	// xcrj golang 方法变量声明时 不需要写 入参变量名称
	var dfs func(*[][]int,int,int,*[][]int) int
	dfs=func(matrix *[][]int,i int,j int,mem *[][]int) int{
		// xcrj golang (*mem)[][], 数组 *取值
		// i,j 已经被访问过
		if (*mem)[i][j]!=0{
			return (*mem)[i][j]
		}
		//
		(*mem)[i][j]++
		//
		for k:=0;k<4;k++{
			// xcrj golang steps是数组，接受的也应该是数组
			var step [2]int=steps[k]
			var nextI int=i+step[0];
			var nextJ int=j+step[1];
			if nextI>=0&&nextI<len(*matrix)&&nextJ>=0&&nextJ<len((*matrix)[0])&&(*matrix)[nextI][nextJ]>(*matrix)[i][j]{
				var nextIJLen=dfs(matrix,nextI,nextJ,mem)
				var ijLen=nextIJLen+1
				if (*mem)[i][j]<ijLen{
					(*mem)[i][j]=ijLen
				}
			}
		}
		//
		return (*mem)[i][j]
	}
	// xcrj golang 二维数组初始化
	var mem [][]int=make([][]int,len(matrix))
	for i:=0;i<len(matrix);i++{
		mem[i]=make([]int,len(matrix[i]))
	}
	// 从每个出发寻找最大长度
	for i:=0;i<len(matrix);i++{
		for j:=0;j<len(matrix[0]);j++{
			var r int=dfs(&matrix,i,j,&mem)
			if maxr < r{
				maxr=r
			}
		}
	}
	//
	return maxr
}

	/**
     * 最长递增路径=拓扑序列长度
     * 广度优先逆向拓扑排序
     * - 构建出度数组
     * - 将所有出度为0的结点放入队列
     * - 每次要将队列中的所有元素处理完
     */
func longestIncreasingPath1(matrix [][]int) int {
	//
	if matrix==nil{
		return 0
	}
	if len(matrix)==0{
		return 0
	}
	if len(matrix[0])==0{
		return 0
	}
	// 上，下，左，右
	var steps [4][2]int=[4][2]int{{1,0},{-1,0},{0,-1},{0,1}}
	var r int=0
	// 构建出度数组
	var out [][]int=make([][]int,len(matrix))
	for i:=0;i<len(matrix);i++{
		out[i]=make([]int,len(matrix[0]))
	}
	// 赋值出度数组
	for i:=0;i<len(matrix);i++{
		for j:=0;j<len(matrix[0]);j++{
			for k:=0;k<4;k++{
				var step [2]int=steps[k]
				var nextI int=i+step[0];
				var nextJ int=j+step[1];
				// matrix[nextI][nextJ]>matrix[i][j]
				if nextI>=0&&nextI<len(matrix)&&nextJ>=0&&nextJ<len(matrix[0])&&matrix[nextI][nextJ]>matrix[i][j]{
					out[i][j]++
				}
			}	
		}
	}
	// 所有出度为0的结点入队
	// xcrj 注意 二维slice的创建不需要 两个括号 {{}} 如果写两个括号就是初始化了二维slice的第一行数据
	var que [][]int=[][]int{}
	for i:=0;i<len(matrix);i++{
		for j:=0;j<len(matrix[0]);j++{
			if out[i][j]==0{
				var ij []int=[]int{i,j}
				que=append(que,ij)
			}
		}
	}
	// 广度优先遍历处理拓扑排序
	for len(que)!=0{
		// 一次把队列中所有出度为=的结点处理完毕
		r++
		var size=len(que)
		for m:=0;m<size;m++{
			var ij []int=que[0]
			que=que[1:]
			//
			var i int=ij[0]
			var j int=ij[1]
			//
			for k:=0;k<4;k++{
				var step [2]int=steps[k]
				var preI int=i+step[0];
				var preJ int=j+step[1];
				// matrix[preI][preJ]<matrix[i][j]
				if preI>=0&&preI<len(matrix)&&preJ>=0&&preJ<len(matrix[0])&&matrix[preI][preJ]<matrix[i][j]{
					out[preI][preJ]--
					if out[preI][preJ]==0{
						var preIJ []int=[]int{preI,preJ}
						que=append(que,preIJ)
					}
				}
			}
		}
	}
	//
	return r
}