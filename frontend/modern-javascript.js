"use strict"; // 严格模式

// javascript规范 https://tc39.es/ecma262/

// 每个语句都应该以;结尾
// https://tc39.es/ecma262/#sec-automatic-semicolon-insertion

// 变量
let message;
// 当它没有被初始化的时候，为 undefined

// 多行变量初始化
let x = 1,
  y = 2,
  z = 3;

const CONSTANTS = 1;

// 数据类型:

// 1. Number (包括浮点数)
// 特殊值: 无穷大 Infinity, 不准确 NaN

// 2. BigInt
// 数字以n结尾为BitInt
// 大于54位有状态数字的存储范围为BitInt

// 3. 字符串
// ``类似python的 f'' format的字符串

// 4. Boolean

// 5. null
// 6. undefined
// 需要区别出null和undefined。

// 7. Object
// 类似于Java,所有的数据类型都是一个Object
// 8. Symbol

console.log(typeof 1);

// 类型转换的细节
console.log(Number(null));
console.log(Number(undefined));

// 数字转换
console.log(+"123f"); // NaN

// = 会再返回一个值
console.log((x = y = z = 1));
console.log((x, y + 1));

// 隐形转换，不同类型的比较，会先转换为数字再比较
// null和undefined不太一样。
// null == 0 为 false。这个是因为相等性和其他如>=设计不一样导致的
// null == undefined 为 true
console.log("0" == 0);
console.log(null >= 0); // true

// javascript 支持 三元运算
console.log(true ? 1 : 2);

// 空值判断
let o;
console.log(o ?? 3); // 3

function echo(msg, prefix = "[cndoit18]") {
  console.log(`${prefix}: ${msg}`);
}
echo("hello world!"); // [cndoit18]: hello world!

let sum = (a, b) => a + b;
// let sum = (a, b) => {
//   return a + b;
// };

console.log(sum(1, 2)); // 3
