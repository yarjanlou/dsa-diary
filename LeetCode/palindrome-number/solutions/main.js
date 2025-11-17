const isPalindrome = function(x) {
    if (x < 0) return false

    const number = x
    let reverseNumber = 0

    while (x > 0) {
        reverseNumber = reverseNumber * 10 + (x % 10)
        
        x = Math.floor(x/10)
    }
    
    
    return number === reverseNumber
};
