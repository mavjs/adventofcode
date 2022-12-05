use std::fs;

fn part1(contents: String)   {
    let mut horizontal = 0;
    let mut depth = 0;

    for line in contents.lines() {
        let command_pair: Vec<&str> = line.split_whitespace().collect();
        let command: String = command_pair[0].parse().unwrap();
        let unit = command_pair[1].parse::<i32>().unwrap();
        if command.as_str() == "forward" {
            horizontal = horizontal + unit;
        } else if command.as_str() == "down" {
            depth = depth + unit;
        } else if command.as_str() == "up" {
            depth = depth - unit;
        }
    }

    let product = horizontal * depth;

    println!("Part 1 planned course: {}", product)
}

fn part2(contents: String) {
    let mut horizontal = 0;
    let mut depth = 0;
    let mut aim = 0;

    for line in contents.lines() {
        let command_pair: Vec<&str> = line.split_whitespace().collect();
        let command: String = command_pair[0].parse().unwrap();
        let unit = command_pair[1].parse::<i32>().unwrap();

        if command.as_str() == "forward" {
            horizontal = horizontal + unit;
            if aim != 0 {
                depth = depth + (unit * aim);
            }
        } else if command.as_str() == "down" {
            aim = aim + unit;
        } else if command.as_str() == "up" {
            aim = aim - unit;
        }
    }

    let product = horizontal * depth;
    println!("Part 2 planned course: {}", product);
}

fn main() {
    let contents = fs::read_to_string("data/input.txt").expect("Something went wrong while reading file.");

    // Because I'm still noob at Rust.
    let contents2 = fs::read_to_string("data/input.txt").expect("Something went wrong while reading file.");

    part1(contents);
    part2(contents2);
}