import fs from "fs";

import { transpose } from "./util.js";

const data = fs
	.readFileSync("../inputs/14.txt")
	.toString()
	.trim()
	.split("\n")
	.map(line => line.split(""));

const grid = transpose(
	transpose(data)
		.map(line => line.join("").split("#"))
		.map(line => line.map(chunk => chunk.split("").toSorted().toReversed().join("")))
		.map(line => line.join("#"))
		.map(line => line.split(""))
);

let sum = 0;
for (const [i, line] of grid.entries()) {
	sum += (grid.length - i) * line.filter(ch => ch === "O").length;
}

console.log(sum);