const binarySearchForSquare = (left, right, target) => {
    if (right < left) return false;
  
    const length = right - left + 1;
    const middle = left + Math.floor(length / 2);
  
    const square = middle * middle;
  
    if (square === target) return true;
    if (square < target) return binarySearchForSquare(middle + 1, right, target);
    if (square > target) return binarySearchForSquare(left, middle - 1, target);
  };
  
  function isPerfectSquare(num) {
    if (num < 2) return true;
    return binarySearchForSquare(2, Math.floor(num / 2), num);
  }