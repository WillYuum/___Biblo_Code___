const grid = [
  [1, 1, 0, 1, 0],
  [1, 1, 1, 1, 0],
  [1, 1, 1, 0, 1],
  [1, 1, 0, 1, 0]
];

function numOffices(grid) {
  let counter = 0;
  //Put your code here.
  for (let x = 0; x < grid.length; x++) {
    for (let y = 0; y < grid[x].length; y++) {
      if (grid[x][y] === 1) {
        if (grid[x][y] === grid[x][y + 1]) {
          if (grid[x + 1][y] === grid[x + 1][y + 1]) {
            grid[x][y] = 0;
            grid[x][y + 1] = 0;
            grid[x + 1][y] = 0;
            grid[x + 1][y + 1] = 0;
            counter += 1;
          } else {
            grid[x][y] = 0;
            grid[x][y + 1] = 0;
            counter += 1;
          }
        } else if ((grid[x][y] = grid[x + 1][y])) {
          grid[x + 1][y] = 0;
          grid[x + 1][y] = 0;
          counter += 1;
        } else {
          grid[x][y] = 0;
          counter += 1;
        }
      }
    }
  }
  return counter;
}

console.log(numOffices(grid));
