package main

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