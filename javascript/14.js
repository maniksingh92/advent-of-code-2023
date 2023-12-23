import fs from "fs";

import { transpose, rotateClockwise, pipe } from "./util.js";

const inputs = fs
	.readFileSync("../inputs/14.txt")
	.toString()
	.trim()
	.split("\n")
	.map(line => line.split(""));

function tiltRocks(data) {
	return pipe(
		transpose,
		x => x.map(line => line.join("").split("#")),
		x => x.map(line => line.map(chunk => chunk.split("").toSorted().toReversed().join(""))),
		x => x.map(line => line.join("#")),
		x => x.map(line => line.split("")),
		transpose,
	)(data);
}

// const grid = tiltRocks(inputs);

// let sum = 0;
// for (const [i, line] of grid.entries()) {
// 	sum += (grid.length - i) * line.filter(ch => ch === "O").length;
// }

// console.log(sum);

function cycle(data) {
	return pipe(
		tiltRocks,
		tiltRocks,
		tiltRocks,
		tiltRocks,
	)(data);
}

function convertToKey(data) {
	return data.map(line => line.join("")).join("");
}

const init = convertToKey(inputs);
const seen = new Set([init]);
const arr = [init];

function isSeen(data) {
	const key = convertToKey(data);
	if (seen.has(key)) return true;

	seen.add(key);
	arr.push(key);

	return false;
}

let grid = inputs;
let iter = 0;
while (true) {
	iter += 1;
	grid = cycle(grid);
	if (isSeen(grid)) break;
}

const first = arr.indexOf(convertToKey(grid));
const result = arr[(1_000_000_000 - first) % (iter - first) + first];

console.log(result)