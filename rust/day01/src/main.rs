use std::collections::{HashMap};
use std::fs::File;
use std::io::{BufRead,BufReader};
use std::time::Instant;

fn main() -> Result<(), Box<dyn std::error::Error>> {
    assert_eq!(11, run_and_measure("A", part_a, "../../data/day01/sample.txt"));
    assert_eq!(31, run_and_measure("B", part_b, "../../data/day01/sample.txt"));

    // 3569916
    run_and_measure("A", part_a, "../../data/day01/input.txt");
    // 26407426
    run_and_measure("B", part_b, "../../data/day01/input.txt");

    Ok(())
}

fn run_and_measure(part: &str, f: fn(BufReader<File>) -> i64, file_name: &str) -> i64 {
    let file = File::open(file_name).unwrap();
    let reader = BufReader::new(file);

    let start = Instant::now();
    let answer = f(reader);
    let duration = start.elapsed();

    println!("Part {} ({}): {} ({} us)", part, file_name, answer, duration.as_micros());

    answer
}

fn part_a(reader: BufReader<File>) -> i64 {
    let (mut left_list, mut right_list): (Vec<i64>, Vec<i64>) = reader.lines()
        .map(|line| line.unwrap())
        .map(|line| {
            let (l_val, r_val) = line.split_once(' ').unwrap();
            (l_val.trim().parse::<i64>().unwrap(), r_val.trim().parse::<i64>().unwrap())
        })
        .unzip();
    left_list.sort();
    right_list.sort();

    left_list.iter()
        .zip(right_list.iter())
        .map(|(a, b)| (b-a).abs())
        .sum()
}

fn part_b(reader: BufReader<File>) -> i64 {
    let mut left_list: Vec<i64> = Vec::new();
    let mut right_map: HashMap<i64, i64> = HashMap::new();

    reader.lines()
        .map(|line| line.unwrap())
        .for_each(|line| {
            let (l_val, r_val) = line.split_once(' ').unwrap();
            left_list.push(l_val.trim().parse::<i64>().unwrap());
            *right_map.entry(r_val.trim().parse::<i64>().unwrap()).or_insert(0) += 1;
        });
        left_list.iter()
            .map(|n| n * right_map.get(n).unwrap_or(&0))
            .sum()
}
