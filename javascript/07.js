import fs from "fs";

const cardStrength = {
	2: 2,
	3: 3,
	4: 4,
	5: 5,
	6: 6,
	7: 7,
	8: 8,
	9: 9,
	T: 10,
	J: 1,
	Q: 12,
	K: 13,
	A: 14,
};

function determineHandStrength(hand) {
	const frequency = Array(hand.length + 1).fill(0);
	const count = {};

	for (const c of hand) {
		count[c] = 1 + (count[c] || 0);
	}

	for (const c in count) {
		if (c === "J") continue;
		frequency[count[c]] += 1;
	}

	if (frequency[5]) return 6;
	if (frequency[4]) {
		if (count["J"] === 1) return 6;
		return 5;
	}
	if (frequency[3] && frequency[2]) return 4;
	if (frequency[3]) {
		if (count["J"] === 2) return 6;
		if (count["J"] === 1) return 5;
		return 3;
	}
	if (frequency[2] === 2) {
		if (count["J"] === 1) return 4;
		return 2;
	}
	if (frequency[2]) {
		if (count["J"] === 3) return 6;
		if (count["J"] === 2) return 5;
		if (count["J"] === 1) return 3;
		return 1;
	}

	if (count["J"] === 5) return 6;
	if (count["J"] === 4) return 6;
	if (count["J"] === 3) return 5;
	if (count["J"] === 2) return 3;
	if (count["J"] === 1) return 1;
	return 0;
}

function compareHands(b, a) {
	for (let i = 0; i < b.length; i++) {
		const cmp = cardStrength[b[i]] - cardStrength[a[i]];
		if (cmp !== 0) return cmp;
	}
	return 0;
}

const lines = fs
	.readFileSync("../inputs/07.txt")
	.toString()
	.split("\n")
	.filter(Boolean);

let hands = lines.map((l) => l.split(" "));
hands = hands
	.map((h) => [...h, determineHandStrength(h[0])])
	.toSorted((a, b) => {
		if (a[2] > b[2]) return 1;
		if (a[2] === b[2]) return compareHands(a[0], b[0]);
		return -1;
	});

console.log(hands.filter((hand) => hand[0].includes("J")));

const sum = hands.reduce((acc, hand, index) => acc + hand[1] * (index + 1), 0);
console.log(sum);
