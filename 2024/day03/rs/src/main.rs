use std::env;

use regex::Regex;

const EXAMPLE_1_INPUT: &str = include_str!("../example_1.txt");
const EXAMPLE_2_INPUT: &str = include_str!("../example_2.txt");
const INPUT: &str = include_str!("../input.txt");

fn part1(inputs: &str) {
    let mut sum = 0;

    for line in inputs.lines() {
        let re = Regex::new(r"mul\((\d+),(\d+)\)").unwrap();
        for (_, [left_str, right_str]) in re.captures_iter(line).map(|captures| captures.extract())
        {
            let left_num: i32 = left_str.parse::<i32>().unwrap();
            let right_num: i32 = right_str.parse::<i32>().unwrap();
            sum = sum + (left_num * right_num);
        }
    }

    println!("Part 1: {}", sum);
}

fn part2(inputs: &str) {
    let mut sum = 0;

    let mut captured_strings: Vec<&str> = vec![];
    for line in inputs.lines() {
        let re = Regex::new(r"(do\(\)|don\'t\(\)|mul\(\d+,\d+\))").unwrap();
        for (_, [capture_str]) in re.captures_iter(line).map(|captures| captures.extract()) {
            captured_strings.push(capture_str);
        }
    }

    let mut skip: bool = false;
    for captured_str in &captured_strings {
        if skip == false && captured_str.starts_with("mul(") {
            let re = Regex::new(r"mul\((\d+),(\d+)\)").unwrap();
            for (_, [left_str, right_str]) in re
                .captures_iter(captured_str)
                .map(|captures| captures.extract())
            {
                let left_num: i32 = left_str.parse::<i32>().unwrap();
                let right_num: i32 = right_str.parse::<i32>().unwrap();
                sum = sum + (left_num * right_num);
            }
        } else if captured_str.starts_with("don't(") {
            skip = true;
        } else if captured_str.starts_with("do(") {
            skip = false;
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
            part1(EXAMPLE_1_INPUT);
            part2(EXAMPLE_2_INPUT);
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
