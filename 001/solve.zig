const std = @import("std");

pub fn main() void {
    var sum: i32 = 0;

    for (1..1000) |i| {
        if (i%3 == 0 or i%5 == 0) {
            sum += @intCast(i);
        }
    }

    std.debug.print("Sum: {d}\n", .{sum});
}
