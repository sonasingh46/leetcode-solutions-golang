package main

/*

Given a 2d grid map of '1's (land) and '0's (water), count the number of islands.
An island is surrounded by water and is formed by connecting adjacent lands
horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

Example 1:

Input:
11110
11010
11000
00000

Output: 1
Example 2:

Input:
11000
11000
00100
00011

Output: 3

*/

/*

*/
func numIslands(grid [][]byte) int {
	cluster:=0
	for i:=0;i<len(grid);i++{
		for j:=0;j<len(grid[0]);j++{
			if grid[i][j]==49{
				cluster++
				bfs(grid,i,j)
			}

		}
	}

	return cluster
}


func bfs(grid [][]byte,i,j int){

	if i<0 || i>=len(grid) || j<0 || j>=len(grid[i]) || grid[i][j]==48{
		return
	}
	grid[i][j]=48
	bfs(grid,i,j+1)
	bfs(grid,i,j-1)
	bfs(grid,i+1,j)
	bfs(grid,i-1,j)

}