function square_sums_row(n) {
  let found = false;
  let result = [];
  let used = new Array(n + 1).fill(false)
  let isSquare = n => n > 0 && Math.sqrt(n) % 1 === 0
  function dfs(k, step) {
    if (step == k + 1) {
      found = true;
      return;
    }
    for (let i = 1; i <= k; i++) {
      if (step == 1 || (!used[i] && isSquare(result[result.length - 1] + i))) {
        result.push(i);
        used[i] = true;
        dfs(k, step + 1)
        if (found) {
          return
        }
        result.pop();
        used[i] = false;
      }
    }
  }

  dfs(n, 1)
  if (result.length != n) {
    return false;
  }
  return result;
}

console.log(square_sums_row(15))