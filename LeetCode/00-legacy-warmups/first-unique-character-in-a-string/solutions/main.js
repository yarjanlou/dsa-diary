const firstUniqChar = (str) => {
  if (str.length === 1) return str;
  const hashMap = {};

  for (let i = 0; i < str.length; i++) {
    hashMap[str[i]] = (hashMap[str[i]] || 0) + 1;
  }

  for (let i in hashMap) {
    if (hashMap[i] === 1) return i;
  }
};


console.log(firstUniqChar("leetcode"))