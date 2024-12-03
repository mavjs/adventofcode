use std::env;

const EXAMPLE_INPUT: &str = include_str!("../example.txt");
const INPUT: &str = include_str!("../input.txt");

fn part1(inputs: &str) {
    let mut sum = 0;

    for line in inputs.lines() {
        let numbers: Vec<i32> = line
            .split_whitespace()
            .map(|s| s.parse().unwrap())
            .collect();
        if (numbers.windows(2).all(|items| items[1] > items[0])
            || numbers.windows(2).all(|items| items[1] < items[0]))
            && (numbers
                .windows(2)
                .all(|items| (items[1] - items[0]).abs() >= 1)
                && numbers
                    .windows(2)
                    .all(|items| (items[1] - items[0]).abs() <= 3))
        {
            sum += 1;
        }
    }

    println!("Part 1: {}", sum);
}

fn part2(inputs: &str) {
    let mut sum = 0;

    for line in inputs.lines() {
        let numbers: Vec<i32> = line
            .split_whitespace()
            .map(|s| s.parse().unwrap())
            .collect();

        if (0..numbers.len()).any(|idx| {
            let temp_diff = &numbers[..idx]
                .iter()
                .chain(&numbers[idx + 1..])
                .collect::<Vec<&i32>>();
            if (temp_diff.windows(2).all(|items| *items[1] > *items[0])
                || temp_diff.windows(2).all(|items| *items[1] < *items[0]))
                && (temp_diff.windows(2).all(|items| (*items[1] - *items[0]).abs() >= 1)
                    && temp_diff.windows(2).all(|items| (*items[1] - *items[0]).abs() <= 3)) {
                        true
                    } else {
                        false
                    }
        }) {
            sum += 1;
        }
    }

    println!("Part 2: {}", sum);
}

fn main() {
    // Read command-line arguments
    let args: Vec<String> = env::args().collect();

    if args.len() != 2 {
        eprintln!("Usage: {} <example|true>", args[0]);
        std::process::exit(1);
    }

    match &args[1] as &str {
        "example" => {
            part1(EXAMPLE_INPUT);
            part2(EXAMPLE_INPUT);
        }
        "true" => {
            part1(INPUT);
            part2(INPUT);
        }
        _ => {
            eprintln!("Invalid argument. Use 'example' or 'true'.");
            std::process::exit(1);
        }
    }
}
