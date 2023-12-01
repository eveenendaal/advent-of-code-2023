fn main() {
    let mut total: i128 = 0;
    // Read in the test.txt file
    let input = std::fs::read_to_string("data.txt").unwrap();
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
