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

const symbols = {
	"|": createTwoWayMap("ns"),
	"-": createTwoWayMap("ew"),
	L: createTwoWayMap("ne"),
	J: createTwoWayMap("nw"),
	7: createTwoWayMap("sw"),
	F: createTwoWayMap("se"),
};

const directionPair = {
	...createTwoWayMap("ns"),
	...createTwoWayMap("ew"),
};

const relativePositionByDirection = {
	n: [-1, 0],
	s: [1, 0],
	e: [0, 1],
	w: [0, -1],
};

function moveToDirection(i, j, direction) {
	const [a, b] = relativePositionByDirection[direction];
	return [i + a, j + b];
}

const nextTiles = [];
for (const direction in relativePositionByDirection) {
	const [x, y] = moveToDirection(s[0], s[1], direction);
	if (x < 0 || x >= inputs.length) continue;
	if (y < 0 || y >= inputs[0].length) continue;

	const swappedDirection = directionPair[direction];

	if (symbols[inputs[x][y]]?.[swappedDirection]) {
		nextTiles.push([x, y, swappedDirection]);
	}
}

function navigateTile(i, j, directionEnteredFrom) {
	const directionToGo = symbols[inputs[i][j]][directionEnteredFrom];
	const [x, y] = moveToDirection(i, j, directionToGo);
	const swappedDirection = directionPair[directionToGo];
	return [x, y, swappedDirection];
}

function buildUniqueId(i, j) {
	return `${i},${j}`;
}

const foundTiles = new Set([buildUniqueId(s[0], s[1])]);
function traversePath() {
	let maxDistance = 0;
	while (true) {
		maxDistance += 1;

		const [tileA, tileB] = nextTiles;
		foundTiles.add(buildUniqueId(tileA[0], tileA[1]));
		foundTiles.add(buildUniqueId(tileB[0], tileB[1]));
		if (tileA[0] === tileB[0] && tileA[1] === tileB[1]) break;

		nextTiles[0] = navigateTile(...tileA);
		nextTiles[1] = navigateTile(...tileB);
	}
	return maxDistance;
}

traversePath();
const trackedSymbols = new Set("|JL");

let totalContainedInsidePath = 0;
for (let i = 0; i < inputs.length; i++) {
	for (let j = 0; j < inputs[i].length; j++) {
		if (foundTiles.has(buildUniqueId(i, j))) continue;

		let crossed = 0;
		for (let k = j - 1; k >= 0; k--) {
			if (!trackedSymbols.has(inputs[i][k])) continue;
			if (foundTiles.has(buildUniqueId(i, k))) crossed += 1;
		}

		if (crossed % 2 !== 0) {
			totalContainedInsidePath += 1;
		}
	}
}

console.log(totalContainedInsidePath);
