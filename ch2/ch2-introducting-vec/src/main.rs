fn main() {
    let ctx_lines = 2;
    let needle = "oo";
    let haystack = "\
Every face, every shop,
bedroom window, public-house, and
dark square is a picture
feverishly turned--in search of what?
It is the same with books.
What do we seek
through millions of pages?";

    // matches
    let mut tags: Vec<usize> = vec![];

    // vector per match to hold context lines
    let mut ctx: Vec<Vec<(usize, String)>> = vec![];

    // iterate through the lines, recording line numbers where matches encountered
    for (i, line) in haystack.lines().enumerate() {
        if line.contains(needle) {
            tags.push(i);

            // reserve space for n items (2 lines top + match + 2 lines bottom)
            let v = Vec::with_capacity(2*ctx_lines + 1);
            ctx.push(v);
        }
    }

    // When there are no matches, exit early
    if tags.is_empty() {
        return;
    }

    // For each tag, at every line, check to see if we are near a match
    // When we are, add that line to the relevant Vec<T> within ctx.
    for (i, line) in haystack.lines().enumerate() {
        for (j, tag) in tags.iter().enumerate() {

            // returns 0 on integer underflow rather than crashing
            let lower_bound = tag.saturating_sub(ctx_lines);
            let upper_bound = tag + ctx_lines;

            if (i >= lower_bound) && (i <= upper_bound) {
                let line_as_string = String::from(line);
                let local_ctx = (i, line_as_string);
                ctx[j].push(local_ctx);
            }
        }
    }

    for local_ctx in ctx.iter() {
        for &(i, ref line) in local_ctx.iter() {
            let line_num = i + 1;
            println!("{}: {}", line_num, line);
        }
    }
}
