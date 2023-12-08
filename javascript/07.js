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
	const freq = Array(hand.length + 1).fill(0);
	const count = { J: 0 };

	for (const c of hand) {
		count[c] = 1 + (count[c] || 0);
	}

	for (const c in count) {
		if (c === "J") continue;
		freq[count[c]] += 1;
	}

	if (freq[5]) return 5;
	if (freq[4]) return 4 + count.J;
	if ((freq[3] && freq[2]) || (freq[2] === 2 && count.J === 1)) return 3.5;
	if (freq[3]) return 3 + count.J;
	if (freq[2] === 2) return 2.5;
	if (freq[2]) return 2 + count.J;
	return Math.min(1 + count.J, 5);
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
