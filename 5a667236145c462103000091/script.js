function buildGraph(n) {
  const squares = [];
  for (let i = 2; i * i < 2 * n; i++) squares.push(i * i);
  const graph = [...Array(n + 1)].map(e => new Set());
  for (let i = 0; i <= n; i++) {
    for (let j of squares) {
      if (i < j) {
        let diff = j - i;
        if (diff == i) continue;
        if (diff <= n) graph[i].add(diff);
        else break;
      }
    }
  }
  return graph;
}

function square_sums_row(n) {
  const graph = buildGraph(n);
  const candidates = [...Array(n)].map((_, val) => val + 1);

  const findNext = (currentCandidates, path) => {
    if (path.length == n) return path;
    currentCandidates.sort((a, b) => graph[a].size - graph[b].size);
    for (let candidate of currentCandidates) {
      path.push(candidate);
      graph[candidate].forEach(e => graph[e].delete(candidate));
      const newCandidates = [...graph[candidate]];
      const result = findNext(newCandidates, path);
      if (result) return result;
      path.pop();
      graph[candidate].forEach(e => graph[e].add(candidate));
    }
    return false;
  }
  return findNext(candidates, []);
}

console.log(square_sums_row(15))