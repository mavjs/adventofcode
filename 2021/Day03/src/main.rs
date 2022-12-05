use std::fs::File;
use std::io::{BufReader, BufRead};

fn part1(vec: &Vec<String>) {
    let mut gamma: Vec<String> = Vec::new();
    let mut epsilon: Vec<String> = Vec::new();
    let bit_length = vec[0].len();

    for bit in 0..bit_length {
        let mut bit_one = 0;
        let mut bit_zero = 0;

        for string_num in vec {
            if string_num.get(bit..bit+1) == Some("0") {
                bit_zero = bit_zero + 1;
            } else {
                bit_one = bit_one + 1;
            }
        }
        if bit_one > bit_zero {
            gamma.push(String::from("1"));
            epsilon.push(String::from("0"));
        } else {
            gamma.push(String::from("0"));
            epsilon.push(String::from("1"));
        }
    }

    let value_gamma = isize::from_str_radix(&gamma.join(""), 2).unwrap();
    let value_epsilon = isize::from_str_radix(&epsilon.join(""), 2).unwrap();

    println!("Part 1 - gamma: {} ({}), epsilon: {} ({}) - power consumption: {}", gamma.join(""), value_gamma, epsilon.join(""), value_epsilon, value_gamma * value_epsilon);
}
// TODO: actually implement part 2 logic here.
fn part2(vec: &Vec<String>) {
    let mut oxygen_generator = vec.clone();
    let mut co2_scrubber = vec.clone();

    let mut index = 0;

    while oxygen_generator.len() > 1 {
        let bit_one = oxygen_generator.clone().into_iter().filter(|x| {
            &x[index..index+1] == "1"
        }).count();
        let bit_zero  = oxygen_generator.len() - bit_one;

        if bit_one >= bit_zero {
            oxygen_generator = oxygen_generator.into_iter().filter(|x| &x[index..index+1] == "1").collect();
        } else {
            oxygen_generator = oxygen_generator.into_iter().filter(|x| &x[index..index+1] == "0").collect();
        }
        index = index + 1;
    }

    index = 0;
    while co2_scrubber.len() > 1 {
        let bit_one = co2_scrubber.clone().into_iter().filter(|x| {
            &x[index..index+1] == "1"
        }).count();
        let bit_zero  = co2_scrubber.len() - bit_one;

        if bit_one >= bit_zero {
            co2_scrubber = co2_scrubber.into_iter().filter(|x| &x[index..index+1] == "0").collect();
        } else {
            co2_scrubber = co2_scrubber.into_iter().filter(|x| &x[index..index+1] == "1").collect();
        }
        index = index + 1;
    }

    let value_oxygen_generator = isize::from_str_radix(&oxygen_generator.join(""), 2).unwrap();
    let value_co2_scrubber = isize::from_str_radix(&co2_scrubber.join(""), 2).unwrap();

    println!("Part 2 - oxygen generator: {} ({}), epsilon: {} ({}) - power consumption: {}", oxygen_generator.join(""), value_oxygen_generator, co2_scrubber.join(""), value_co2_scrubber, value_oxygen_generator * value_co2_scrubber);

}

fn main() {
    let contents = File::open("data/input.txt").unwrap();
    let reader = BufReader::new(contents);

    let vec: Vec<String> = reader.lines().map(|line| line.expect("Error reading line")).collect::<Vec<String>>().clone();
    
    part1(&vec);

    part2(&vec);
}