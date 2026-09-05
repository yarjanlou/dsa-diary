const SingleNumber = (nums) => {
  const numsObj = {};

  for (let i = 0; i < nums.length; i++) {
    numsObj[nums[i].toString()] = (numsObj[nums[i].toString()] || 0) + 1;
  }

  for (let key in numsObj) {
    if (numsObj[key] === 1) return +key;
  }
};
