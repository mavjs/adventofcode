use std::fs::File;
use std::io::{BufReader, BufRead};

fn part_two(vec: &Vec<i32>) {
    let mut counter = vec.len();

    while counter > 0 {
        counter -= 1;
    }
}

fn part_one(vec: &Vec<i32>) {
    let mut counter = vec.len();


    while counter > 0 {

        let mut prev = 0i32;
        let next = vec[counter-1];

        for item in vec.iter() {

            if prev == next {
                break;
            }

            let total = item + next;

            if total == 2020 {
                println!("Part#1 - answer ({} + {}): {}", item, next, item * next);
                break;
            }

            prev = *item;
        }

        counter -= 1;
    }
}

fn main() {
    let contents = File::open("data/input").unwrap();
    let reader = BufReader::new(contents);
    let mut vec = Vec::new();

    for line in reader.lines() {
        let number = line.unwrap().parse::<i32>().unwrap();
        vec.push(number);
    }

    // Call the function for part 1.
    part_one(&vec);
   
   // Call the function for part 2.
   part_two(&vec);
}
