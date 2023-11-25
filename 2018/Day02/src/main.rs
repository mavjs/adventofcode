
use itertools::Itertools;

const INPUT: &str = include_str!("../data/input.txt");

// Copied shamelessly from: https://github.com/BurntSushi/advent-of-code/blob/master/aoc02/src/main.rs
fn part1(input: &str) -> i64 {
    let mut ids = [0u8; 256];

    let (mut twos, mut threes) = (0, 0);

    for line in input.lines() {

        for f in ids.iter_mut() {
            *f = 0;
        }

        for b in line.as_bytes().iter().map(|&b| b as usize) {
            ids[b] = ids[b].saturating_add(1);
        }

        if ids.iter().any(|&f| f == 2) {
            twos = twos + 1
        }

        if ids.iter().any(|&f| f == 3) {
            threes = threes + 1
        }
    };
    println!("Twos: {}, Threes: {}", twos, threes);
    let part1 = twos * threes;

    return part1;
}

// Copied shamelessly from: https://github.com/angch/adventofcode/blob/master/2018-02/opcode0x90/src/lib.rs
fn part2(input: &Vec<String>) -> String {
    let part2 = input
                            .into_iter()
                            .filter_map(|line| line.parse::<String>().ok())
                            .tuple_combinations()
                            .find_map(|(a, b)| {
                                // find both common and distinct chars
                                let mut same = String::with_capacity(a.len());
                                let mut diff = 0;

                                for (a, b) in a.chars().zip(b.chars()) {
                                    if a == b {
                                        same.push(a);
                                    } else {
                                        diff += 1;

                                        // early skipping if diff > 1
                                        if diff > 1 {
                                            break;
                                        }
                                    }
                                }

                                match diff {
                                    1 => Some(same),
                                    _ => None,
                                }
                            }).expect("there is no candiate for difference of one");
    return part2
}

fn main() {
    let part1 = part1(INPUT);
    println!("[Part 1]: {}", part1);

    let part2 = part2(&INPUT.lines().filter_map(|line| line.parse::<String>().ok()).collect::<Vec<String>>());
    println!("[Part 2]: {}", part2);
}

#[cfg(test)]
mod tests {
    use crate::{part1, part2};

    const INPUT: &str = include_str!("../data/test-input.txt");

    #[test]
    fn test_input_part1(){
        let test_parts = part1(INPUT);
        assert_eq!(test_parts, 12);
    }

    #[test]
    fn test_input_part2() {
        let common = part2(
            &vec!["abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"]
                    .into_iter()
                    .map(|s| String::from(s))
                    .collect::<Vec<String>>()
        );
        assert_eq!(common, "fgij");
    }
}