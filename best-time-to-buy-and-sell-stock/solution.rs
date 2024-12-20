fn main() {
    println!("{}", max_profit(vec![7, 1, 5, 3, 6, 4]));
    println!("{}", max_profit(vec![7, 6, 4, 3, 1]));
}

pub fn max_profit(prices: Vec<i32>) -> i32 {
    let mut min_price = i32::MAX;
    let mut max_profit: i32 = 0;

    for price in prices {
        if price < min_price {
            min_price = price;
        }

        let profit = price - min_price;

        if profit > max_profit {
            max_profit = profit;
        }
    }

    return max_profit;
}
