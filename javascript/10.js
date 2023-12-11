import fs from "fs";

const inputs = fs
	.readFileSync("../inputs/10.txt")
	.toString()
	.trim()
	.split("\n");

const s = (() => {
	for (let i = 0; i < inputs.length; i++) {
		for (let j = 0; j < inputs[i].length; j++) {
			if (inputs[i][j] === "S") return [i, j];
		}
	}
})();

function createTwoWayMap(directions) {
	return {
		[directions[0]]: directions[1],
		[directions[1]]: directions[0],
	};
}

const tiles = {
	"|": createTwoWayMap("ns"),
	"-": createTwoWayMap("ew"),
	L: createTwoWayMap("ne"),
	J: createTwoWayMap("nw"),
	7: createTwoWayMap("sw"),
	F: createTwoWayMap("se"),
};

const swapDirection = {
	...createTwoWayMap("ns"),
	...createTwoWayMap("ew"),
};

const relativePositionByDirection = {
	n: [-1, 0],
	s: [1, 0],
	e: [0, 1],
	w: [0, -1],
};

function navigateTile(i, j, direction) {
	const [a, b] = relativePositionByDirection[direction];
	return [i + a, j + b];
}

const nextTiles = [];
for (const direction in relativePositionByDirection) {
	const [x, y] = navigateTile(s[0], s[1], direction);
	if (x < 0 || x >= inputs.length) continue;
	if (y < 0 || y >= inputs[0].length) continue;

	const swappedDirection = swapDirection[direction];

	if (tiles[inputs[x][y]]?.[swappedDirection]) {
		nextTiles.push([x, y, swappedDirection]);
	}
}

function processTile(i, j, directionEnteredFrom) {
	const directionToGo = tiles[inputs[i][j]][directionEnteredFrom];
	const [x, y] = navigateTile(i, j, directionToGo);
	const swappedDirection = swapDirection[directionToGo];
	return [x, y, swappedDirection];
}

let maxDistance = 0;
while (true) {
	maxDistance += 1;

	const [tileA, tileB] = nextTiles;
	if (tileA[0] === tileB[0] && tileA[1] === tileB[1]) break;

	nextTiles[0] = processTile(...tileA);
	nextTiles[1] = processTile(...tileB);
}

console.log(maxDistance);
