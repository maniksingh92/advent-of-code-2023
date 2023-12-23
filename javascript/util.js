export function transpose(array) {
    return array[0].map((_, i) => array.map((line) => line[i]));
}

export function rotateClockwise(array) {
    return array[0].map((_, index) => array.map(row => row[index]).toReversed())
}

export const pipe = (...fns) => (x) => fns.reduce((v, f) => f(v), x);
