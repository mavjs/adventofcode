use std::collections::HashSet;

const INPUT: &str = include_str!("../data/input.txt");

fn main() {
    let part1 = INPUT.lines().filter_map(|line| line.parse::<isize>().ok()).sum::<isize>();
    println!("[Part 1]: {}", part1);

    let mut cache = HashSet::new();
    let mut sum = 0;
    let part2 = INPUT
                    .lines()
                    .filter_map(|line| line.parse::<i64>().ok())
                    .cycle()
                    .find_map(|c| {
                        sum = sum + c;
                        cache.replace(sum)
                    });
    println!("[Part 2]: {}", part2.unwrap());
    println!("Sum: {}", sum);
}
