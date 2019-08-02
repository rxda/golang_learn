package dfsbfs

func bfs(g map[string][]string, start string) []string {
	queue := []string{start}
	seen := make(map[string]struct{})
	result := []string{}
	seen[start] = struct{}{}

	for len(queue) > 0 {
		vertex := queue[0]
		result = append(result, queue[0])
		queue = queue[1:]
		nodes := g[vertex]
		for _, w := range nodes {
			if _, ok := seen[w]; !ok {
				queue = append(queue, w)
				// result = append(result, w)
				seen[w] = struct{}{}
			}
		}

	}
	return result
}
