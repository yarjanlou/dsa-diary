function buildLPS(pattern) {
  const m = pattern.length;
  const lps = new Array(m).fill(0);
  let len = 0;
  let i = 1;

  while (i < m) {
    if (pattern[i] === pattern[len]) {
      len++;
      lps[i] = len;
      i++;
    } else {
      if (len !== 0) {
        len = lps[len - 1];
      } else {
        lps[i] = 0;
        i++;
      }
    }
  }
  return lps;
}

var strStr = function (haystack, needle) {
  const n = haystack.length;
  const m = needle.length;

  if (m === 0) return 0;
  if (n < m) return -1;

  const lps = buildLPS(needle);
  let i = 0;
  let j = 0;

  while (i < n) {
    if (haystack[i] === needle[j]) {
      i++;
      j++;

      if (j === m) {
        return i - j;
      }
    } else {
      if (j !== 0) {
        j = lps[j - 1];
      } else {
        i++;
      }
    }
  }
  return -1;
};
