package graph

import (
	base "golang-study/algorithm/basic"
)

// UndirectedGraph 无向图
type UndirectedGraph struct {
	vertexes int
	edges    int
	adj      []*base.List
}

// NewUndirectedGraph 创建无向图，通过邻接矩阵初始化
func NewUndirectedGraph(matrix [][]bool) *UndirectedGraph {
	g := &UndirectedGraph{
		vertexes: len(matrix),
		edges:    0,
		adj:      make([]*base.List, len(matrix)),
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			if matrix[i][j] {
				g.addEdge(i, j)
			}
		}
	}
	return g
}

func (g *UndirectedGraph) addEdge(v, w int) {
	g.adj[v].Add(w)
	g.adj[w].Add(v)
	g.edges++
}

// SearchConnected 搜索指定节点能到达的所有节点
func (g *UndirectedGraph) SearchConnected(sourceVertex int) []bool {
	marked := make([]bool, g.vertexes)
	g.dfs(sourceVertex, marked)
	return marked
}

func (g *UndirectedGraph) dfs(currentVertex int, marked []bool) {
	marked[currentVertex] = true
	for i := g.adj[currentVertex].Iterator(); i.HasNext(); {
		nextVertex := i.Next().(int)
		if !marked[nextVertex] {
			g.dfs(nextVertex, marked)
		}
	}
}

// SearchPaths 搜索从指定节点到其他所有节点的路径
func (g *UndirectedGraph) SearchPaths(sourceVertex int) []int {
	marked := make([]bool, g.vertexes)
	edgeTo := make([]int, g.vertexes)
	g.dfs2(sourceVertex, marked, edgeTo)
	return edgeTo
}

func (g *UndirectedGraph) dfs2(currentVertex int, marked []bool, edgeTo []int) {
	marked[currentVertex] = true
	for i := g.adj[currentVertex].Iterator(); i.HasNext(); {
		nextVertex := i.Next().(int)
		if !marked[nextVertex] {
			edgeTo[nextVertex] = currentVertex // 记录路径
			g.dfs(nextVertex, marked)
		}
	}
}

// ShortestPath 最短路径
func (g *UndirectedGraph) ShortestPath(sourceVertex int) []int {
	queue := base.NewQueue()
	marked := make([]bool, g.vertexes)
	edgeTo := make([]int, g.vertexes)

	marked[sourceVertex] = true
	queue.Enqueue(sourceVertex)
	for queue.Size() != 0 {
		currentVertex := queue.Dequeue().(int)
		for i := g.adj[currentVertex].Iterator(); i.HasNext(); {
			nextVertex := i.Next().(int)
			if !marked[nextVertex] {
				edgeTo[nextVertex] = currentVertex // 记录路径
				marked[nextVertex] = true
				queue.Enqueue(nextVertex)
			}
		}
	}

	return edgeTo
}

// SearchConnectedComponents 找到连通子图，为每个节点标记所属连通子图id
func (g *UndirectedGraph) SearchConnectedComponents() []int {
	marked := make([]bool, g.vertexes)
	componentID := make([]int, g.vertexes)
	componentCount := 0

	for v := 0; v < g.vertexes; v++ {
		if !marked[v] {
			g.dfs3(v, componentCount, marked, componentID)
			componentCount++
		}
	}

	return componentID
}

func (g *UndirectedGraph) dfs3(currentVertex, componentCount int, marked []bool, componentID []int) {
	marked[currentVertex] = true
	componentID[currentVertex] = componentCount
	for i := g.adj[currentVertex].Iterator(); i.HasNext(); {
		nextVertex := i.Next().(int)
		if !marked[nextVertex] {
			g.dfs3(nextVertex, componentCount, marked, componentID)
		}
	}
}
