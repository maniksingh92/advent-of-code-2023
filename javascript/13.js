import fs from "fs";

const inputs = fs
	.readFileSync("../inputs/13.txt")
	.toString()
	.trim()
	.split("\n\n")
	.map((grid) => grid.split("\n"))
	.map((grid) => grid.map((line) => line.split("")));

function transpose(array) {
	return array[0].map((_, i) => array.map((line) => line[i]));
}

function findMirror(grid) {
	for (let i = 1; i < grid.length; i++) {
		let above = grid.slice(0, i).toReversed();
		let below = grid.slice(i);
		above = above.slice(0, below.length);
		below = below.slice(0, above.length);

		if (above.join("") === below.join("")) {
			return i;
		}
	}

	return 0;
}

let sum = 0;
for (const grid of inputs) {
	sum += findMirror(grid) * 100;
	sum += findMirror(transpose(grid));
}

console.log(sum);
