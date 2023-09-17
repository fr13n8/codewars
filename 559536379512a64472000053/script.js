function playPass(s, n) {
  let result = '';
  for (let i = 0; i < s.length; i++) {
    if (/\d/.test(s[i])) {
      let num = s[i] - '0';
      result += (9 - num);
    } else if (/[a-z]/.test(s[i])) {
      let num = s[i].charCodeAt() - 'a'.charCodeAt();
      let newNum = (num + n) % 26;
      result += String.fromCharCode(newNum + 'a'.charCodeAt());
    } else if (/[A-Z]/.test(s[i])) {
      let num = s[i].charCodeAt() - 'A'.charCodeAt();
      let newNum = (num + n) % 26;
      result += String.fromCharCode(newNum + 'A'.charCodeAt());
    } else {
      result += s[i];
    }
  }
  let newResult = result.split('');
  for (let i = 0; i < newResult.length; i++) {
    if (i % 2 === 0) {
      newResult[i] = newResult[i].toUpperCase();
    } else {
      newResult[i] = newResult[i].toLowerCase();
    }
  }
  newResult.reverse();
  return newResult.join('');
}

console.log(playPass("BORN IN 2015!", 1))