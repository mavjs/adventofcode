use std::fs::File;
use std::io::{BufReader, BufRead};

fn part1(vec: &Vec<i32>) {
    let mut prev_num = vec[0];
    let mut sum = 0;

    for index in 1..vec.len() {
        let cur_num = vec[index];
        if cur_num > prev_num {
            sum += 1
        }
        prev_num = cur_num;
    }
    println!("Part 1 sum: {}", sum)
}

fn part2(vec: &Vec<i32>) {
    let mut prev_window = vec[0] + vec[1] + vec[2];
    let mut sum = 0;
    let mut cur_window;

    for index in 1..vec.len() {
        if vec.len() - index < 2 {
            cur_window = vec[index];
        } else if vec.len() - index < 3 {
            cur_window = vec[index] + vec[index+1];
        } else {
            cur_window = vec[index] + vec[index+1] + vec[index+2];
        }


        if cur_window > prev_window {
            sum += 1
        }

        prev_window = cur_window;
    }
    println!("Part 2 sum: {}", sum)
}

fn main() {
    let contents = File::open("data/input.txt").unwrap();
    let reader = BufReader::new(contents);
    let mut vec = Vec::new();

    for line in reader.lines() {
        let number = line.unwrap().parse::<i32>().unwrap();
        vec.push(number);
    }
    
    part1(&vec);

    part2(&vec);
}