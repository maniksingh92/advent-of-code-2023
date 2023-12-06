from typing import List, Tuple

def parse_first_line(line: str):
    return list(map(int, line.rstrip().split(": ")[1].split()))

def process_data_maps(lines: List[str]):
    data_maps: List[List[Tuple[int, ...]]] = []
    for line in lines:
        line = line.rstrip()

        if len(line) == 0:
            data_maps.append([])
            continue

        if line.endswith("map:"):
            continue

        data_maps[-1].append(tuple(map(int, line.split())))

    return data_maps

def day_05_puzzle_1(inputs: List[str]):
    seeds = parse_first_line(inputs[0])
    data_maps = process_data_maps(inputs[1:])

    for data_map in data_maps:
        mapped = []
        for curr in seeds:
            for destination, source, length in data_map:
                if source <= curr <= source + length:
                    mapped.append(destination + curr - source)
                    break
            else:
                mapped.append(curr)
        seeds = mapped

    print(min(seeds))

def day_05_puzzle_2(inputs: List[str]):
    first_line = parse_first_line(inputs[0])
    data_maps = process_data_maps(inputs[1:])

    seeds: List[Tuple[int, int]] = []
    for i in range(0, len(first_line), 2):
        seeds.append((first_line[i],first_line[i] + first_line[i+1]))

    for data_map in data_maps:
        mapped: List[Tuple[int, int]] = []
        while len(seeds) > 0:
            seed_start, seed_end = seeds.pop()
            for destination_start, source_start, range_length in data_map:
                overlap_start = max(source_start, seed_start)
                overlap_end = min(source_start+range_length, seed_end)
                if overlap_start < overlap_end:
                    mapped.append((destination_start + overlap_start - source_start, destination_start + overlap_end - source_start))
                    if overlap_start > seed_start:
                        seeds.append((seed_start, overlap_start))
                    if overlap_end < seed_end: 
                        seeds.append((overlap_end, seed_end))
                    break
            else:
                mapped.append((seed_start, seed_end))
        seeds = mapped
    print(min(seeds))

if __name__ == "__main__":
    with open("../inputs/05.txt") as inputs:
        day_05_puzzle_2(inputs.readlines())
