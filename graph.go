package main

func bfsShortestPath(adjList map[string][]string, start, end string) []string {
	if start == end {
		return []string{start}
	}

	visited := make(map[string]bool)
	prev := make(map[string]string)
	queue := []string{start}
	visited[start] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, neighbor := range adjList[node] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
				visited[neighbor] = true
				prev[neighbor] = node

				if neighbor == end {
					return reconstructPath(prev, start, end)
				}
			}
		}
	}
	return nil
}

func reconstructPath(prev map[string]string, start, end string) []string {
	var path []string
	for at := end; at != ""; at = prev[at] {
		path = append([]string{at}, path...)
		if at == start {
			break
		}
	}
	return path
}
