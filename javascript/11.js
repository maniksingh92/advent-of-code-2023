import fs from "fs";
import { argv0 } from "process";

let inputs = fs
	.readFileSync("../inputs/11.txt")
	.toString()
	.trim()
	.split("\n")
	.map((line) => line.split(""));

const emptyRows = inputs.map((line) => line.every((ch) => ch === "."));

const emptyColumns = inputs[0].map((_, column) =>
	inputs.every((line) => line[column] === "."),
);

function findDistance(a, b) {
	const [ax, ay] = a;
	const [bx, by] = b;

	const [x1, x2] = [Math.min(ax, bx), Math.max(ax, bx)];
	const [y1, y2] = [Math.min(ay, by), Math.max(ay, by)];

	let distance = 0;
	for (let n = x1; n < x2; n++) {
		distance += 1;
		if (emptyRows[n]) distance += 10 ** 6 - 1;
	}
	for (let n = y1; n < y2; n++) {
		distance += 1;
		if (emptyColumns[n]) distance += 10 ** 6 - 1;
	}

	return distance;
}

const galaxyCoordinates = [];
for (let i = 0; i < inputs.length; i++) {
	for (let j = 0; j < inputs[i].length; j++) {
		if (inputs[i][j] === "#") galaxyCoordinates.push([i, j]);
	}
}

let total = 0;
for (const g1 of galaxyCoordinates) {
	for (const g2 of galaxyCoordinates) {
		total += findDistance(g1, g2);
	}
}

console.log(total / 2);
