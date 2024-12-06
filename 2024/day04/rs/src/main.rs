use std::env;

const EXAMPLE_INPUT: &str = include_str!("../example.txt");
const INPUT: &str = include_str!("../input.txt");

fn part1(inputs: &str) {
    let mut sum = 0;

    let grid = inputs
        .lines()
        .map(|line| line.chars().collect())
        .collect::<Vec<Vec<char>>>();
    let word = "XMAS";
    let reverse_word = word.chars().rev().collect::<String>();

    let rows = grid.len();
    let cols = if rows > 0 { grid[0].len() } else { 0 };

    // check horizontally and vertically
    // main loops
    let mut search_words = vec![];

    for first_idx in 0..rows {
        for second_idx in 0..cols {
            // check for word horizontally (both forward and reverse)
            if second_idx + word.len() <= cols {
                let search_str: String = grid[first_idx][second_idx..second_idx + word.len()]
                    .iter()
                    .collect();

                if search_str == word || search_str == reverse_word {
                    search_words.push(search_str);
                    sum = sum + 1;
                }
            }

            // check for word vertically (both forward and reverse)
            if first_idx + word.len() <= rows {
                let mut search_str = String::new();
                for idx in 0..word.len() {
                    search_str.push(grid[first_idx + idx][second_idx]);

                    if search_str == word || search_str == reverse_word {
                        search_words.push(search_str.clone());
                        sum = sum + 1;
                    }
                }
            }

            // check main diagonal (left-to-right)
            // Main diagonal
            if first_idx + word.len() <= rows && second_idx + word.len() <= cols {
                let mut search_str = String::new();
                for k in 0..word.len() {
                    search_str.push(grid[first_idx + k][second_idx + k]);
                }
                if search_str == word || search_str == reverse_word {
                    search_words.push(search_str);
                    sum = sum + 1;
                }
            }

            // Anti-diagonal
            if first_idx + word.len() <= rows && second_idx as isize - word.len() as isize >= -1 {
                let mut search_str = String::new();
                for k in 0..word.len() {
                    search_str.push(grid[first_idx + k][second_idx - k]);
                }
                if search_str == word || search_str == reverse_word {
                    search_words.push(search_str);
                    sum = sum + 1;
                }
            }
        }
    }
    println!("Part 1: {}", sum);
}

fn part2(inputs: &str) {
    let mut sum = 0;

    let grid = inputs
        .lines()
        .map(|line| line.chars().collect())
        .collect::<Vec<Vec<char>>>();

    let rows = grid.len();
    let cols = if rows > 0 { grid[0].len() } else { 0 };

    for first_idx in 0..(rows - 2) {
        for second_idx in 0..(cols - 2) {
            if (grid[second_idx][first_idx] == 'M'
                && grid[second_idx + 1][first_idx + 1] == 'A'
                && grid[second_idx + 2][first_idx + 2] == 'S'
                || grid[second_idx][first_idx] == 'S'
                    && grid[second_idx + 1][first_idx + 1] == 'A'
                    && grid[second_idx + 2][first_idx + 2] == 'M')
                && (grid[second_idx + 2][first_idx] == 'M'
                    && grid[second_idx + 1][first_idx + 1] == 'A'
                    && grid[second_idx][first_idx + 2] == 'S'
                    || grid[second_idx + 2][first_idx] == 'S'
                        && grid[second_idx + 1][first_idx + 1] == 'A'
                        && grid[second_idx][first_idx + 2] == 'M')
            {
                sum = sum + 1;
            }
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
