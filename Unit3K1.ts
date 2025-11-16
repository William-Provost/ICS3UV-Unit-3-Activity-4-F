/**
 * @author Mr Coxall
 * @version 1.0.0
 * @date 2025-01-01
 * @fileoverview This program shows comparing strings with different cases.
 */

// variables
let string1: string = "Hello";
let string2: string = "HELLO";
let string3: string = "I LOVE CS!";
let string4: string = "I LOVE CS!";

// check if string are the same
// NOTE: using regular "equal" operator
if (string1 == string2) {
  console.log("'" + string1 + "' is the same as '" + string2 + "'.");
} else if (string1.toLowerCase() == string2.toLowerCase()) {
  console.log("'" + string1 + "' is the same as '" + string2 + "' (lower case insensitive).");
} else {
  console.log("'" + string1 + "' is NOT the same as '" + string2 + "'.");
}

// check if string are the same
// NOTE: using strict (also checking type) "equal" operator
if (string3 === string4) {
  console.log("'" + string3 + "' is the same as '" + string4 + "' (upper case insensitive).");
} else if (string3.toUpperCase() === string4.toUpperCase()) {
  console.log("'" + string3 + "' is the same as '" + string4 + "' (upper case insensitive).");
} else {
  console.log("'" + string3 + "' is NOT the same as '" + string4 + "'.");
}

console.log("\nDone.");