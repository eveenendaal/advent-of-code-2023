fn part1() {
    let mut total: i128 = 0;
    // Read in the test.txt file
    let input = std::fs::read_to_string("data1.txt").unwrap();
    // Iterate over each line
    for line in input.lines() {
        // replace all letters with a |
        let line = line.replace(|c: char| c.is_alphabetic(), "|");
        // split on |
        let line = line.split("|");
        let mut first = "";
        let mut last = "";
        // iterate over all the elements and print them
        for element in line.clone() {
            // skip empty elements
            if element == "" {
                continue;
            }
            // set if first is empty
            if first == "" {
                // get first character
                first = &element[0..1];
            }
            // set last to the current element's last character
            last = &element[element.len() - 1..element.len()];
        }
        // concat first and last and convert to int
        let next: i128 = format!("{}{}", first, last).parse::<i128>().unwrap();
        // print next
        println!("next: {}", next);
        // add to total
        total += next;
    }
    // print total
    println!("total: {}", total);
}

fn part2() {
    let mut total: i128 = 0;
    // Read in the test.txt file
    let input = std::fs::read_to_string("test2.txt").unwrap();
    // list of number words
    let numbers = vec!["zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"];

    // Iterate over each line
    for line in input.lines() {
        // find the first number in the line
        let mut first_position = line.find(char::is_numeric).unwrap();
        // get character
        let mut first = &line[first_position..first_position + 1];

        // find the last number in the line
        let mut last_position = line.rfind(char::is_numeric).unwrap();
        // get character
        let mut last = &line[last_position..last_position + 1];

        let mut value: String = "".to_string();

        // iterate over each number words
        for number in numbers.clone() {
            // get position of number
            let value = line.find(number).unwrap();
            // convert usize to string slice
            let value = format!("{}", value);
            
            // find first word in line
            let next = line.find(number);
            // if next is found
            if next.is_some() {
                let next = next.unwrap();
                // if next before first_position, set first_position to next
                if next < first_position {
                    first_position = next;
                    // copy value to first
                    first = value;
                }
            }
            

            // find last instance in line
            let last = line.rfind(number);
        }

        // concat first and last and convert to int
        let next: i128 = format!("{}{}", first, last).parse::<i128>().unwrap();
        // print next
        println!("next: {}", next);
        // add to total
        total += next;
    }
    // print total
    println!("total: {}", total);
}

fn main() {
    part2();
}