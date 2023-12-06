from collections.abc import Iterator
from typing import List

def parse_first_line(line: str):
    return list(map(int, line.rstrip().split(": ")[1].split()))

def process_data_maps(lines: List[str]):
    data_maps: List[List[List[int]]] = []
    for line in lines:
        line = line.rstrip()

        if len(line) == 0:
            data_maps.append([])
            continue

        if line.endswith("map:"):
            continue

        data_maps[-1].append(list(map(int, line.split())))

    return data_maps

def day_05_puzzle_1(inputs: List[str]):
    seeds = parse_first_line(inputs[0])

    data_maps = process_data_maps(inputs[1:])

    locations: List[int] = []
    for seed in seeds:
        curr = seed
        for data_map in data_maps:
            for data in data_map:
                destination, source, length = data
                if source <= curr <= source + length:
                    curr = destination + curr - source
                    break
        locations.append(curr)

    print(min(locations))

def day_05_puzzle_2(inputs: Iterator[str]):
    first_line = parse_first_line(next(inputs))

    seeds: List[List[int]] = []
    for i in range(0, len(first_line), 2):
        seeds.append([first_line[i],first_line[i+1]])

if __name__ == "__main__":
    with open("../inputs/05.txt") as inputs:
        day_05_puzzle_1(inputs.readlines())
