package dfsbfs

func dfs(g map[string][]string, start string) []string {
	stack := []string{start}
	seen := make(map[string]struct{})
	result := []string{}
	seen[start] = struct{}{}

	for len(stack) > 0 {
		vertex := stack[len(stack)-1]
		result = append(result, vertex)
		stack = stack[:len(stack)-1]
		nodes := g[vertex]
		for _, w := range nodes {
			if _, ok := seen[w]; !ok {
				stack = append(stack, w)
				// result = append(result, w)
				seen[w] = struct{}{}
			}
		}

	}
	return result
}
