use std::env;

const EXAMPLE_INPUT: &str = include_str!("../example.txt");
const INPUT: &str = include_str!("../input.txt");

fn part1(inputs: &str) {
    let mut col1: Vec<i32> = Vec::new();
    let mut col2: Vec<i32> = Vec::new();

    for line in inputs.lines() {
        // Split each line by whitespace and collect numbers into columns
        let parts: Vec<&str> = line.split_whitespace().collect();
        if parts.len() == 2 {
            if let Ok(num1) = parts[0].parse::<i32>() {
                col1.push(num1);
            }
            if let Ok(num2) = parts[1].parse::<i32>() {
                col2.push(num2);
            }
        } else {
            println!("Skipping malformed line: {}", line);
        }
    }

    col1.sort();
    col2.sort();

    // Ensure both vectors have the same length before zipping them
    if col1.len() != col2.len() {
    panic!("Column lengths are not equal after sorting.");
    }

    let mut sum = 0;

    // Iterate through both columns simultaneously and compute the sum of absolute differences
    for (num1, num2) in col1.iter().zip(col2.iter()) {
        let diff = (*num1 - *num2).abs();
        sum += diff;
    }

    println!("Part 1: {}", sum);

    
}


fn part2(inputs: &str) {
    let mut col1: Vec<i32> = Vec::new();
    let mut col2: Vec<i32> = Vec::new();

    for line in inputs.lines() {
        // Split each line by whitespace and collect numbers into columns
        let parts: Vec<&str> = line.split_whitespace().collect();
        if parts.len() == 2 {
            if let Ok(num1) = parts[0].parse::<i32>() {
                col1.push(num1);
            }
            if let Ok(num2) = parts[1].parse::<i32>() {
                col2.push(num2);
            }
        } else {
            println!("Skipping malformed line: {}", line);
        }
    }

    col1.sort();
    col2.sort();

     // Ensure both vectors have the same length before zipping them
    if col1.len() != col2.len() {
        panic!("Column lengths are not equal after sorting.");
    }

    let mut sum = 0;

    for num in &col1 {
        let count_in_col2 = col2.iter().filter(|&&x| x == *num).count();
        sum += num * count_in_col2 as i32;
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
        },
        "true" => {
            part1(INPUT);
            part2(INPUT);
        },
        _ => {
            eprintln!("Invalid argument. Use 'example' or 'true'.");
            std::process::exit(1);
        }
    }
}