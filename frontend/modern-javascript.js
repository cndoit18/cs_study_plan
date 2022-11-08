"use strict"; // 严格模式

// javascript规范 https://tc39.es/ecma262/

// 每个语句都应该以;结尾
// https://tc39.es/ecma262/#sec-automatic-semicolon-insertion
// 或者
// https://standardjs.com/

// 变量
let message;
// 当它没有被初始化的时候，为 undefined

// 多行变量初始化
let x = 1,
  y = 2,
  z = 3;

// https://developer.chrome.com/docs/devtools/
debugger;

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

// https://en.wikipedia.org/wiki/JSDoc
/**
 * 返回 x 的 n 次幂的指。
 *
 * @param {number} x
 * @param {number} n
 * @return {number} 返回 x 的 n 次幂
 */
function pow(x, n) {
  return x ** n;
}

let obj = {
  name: "object",
  Welcome() {
    // this表示是对象的上下文环境
    console.log(this); // this 是 object
    console.log(this.name);
  },
  Echo: () => console.log(this),
};

obj.Welcome();
delete obj.Welcome; // 可以删除方法

console.log(obj.Welcome === undefined);

// Object.assign 浅拷贝

obj.Echo(); // this 是 windows

function User(name) {
  this.name = name;
}

let user = new User("name");

// 类似于 groovy 的 ?.
console.log(obj.test?.());

console.log(Symbol(1) === Symbol(1)); // false

// Symbol 会被 for ... in 跳过
let sl = Symbol.for("name"); // 注册到全局中
console.log(Symbol.for("name") === sl);
console.log(Symbol.keyFor(sl)); // name

// 类型转换
let h = {
  [Symbol.toPrimitive](hint) {
    console.log(hint);
    return "[Symbol.toPrimitive]";
  },
  toString() {
    return "toString";
  },
  valueOf() {
    return "valueOf";
  },
};

console.log(String(h)); // [Symbol.toPrimitive]

// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects

// WeakMap WeakSet

// 解构
let [first, , ...second] = ["a", "b", "c", "d"];
console.log(first, second); // a [c, d]

let options = {
  title: "Menu",
};

let { width: ow = 100, height: oh = 200, title } = options;

// 与 golang 的不同，js的Spread语法和Rest参数都是在变量前面加 ...

let sayHi = function func(who) {
  console.log(`${who} ${sayHi.name}`);
};

sayHi("test"); // test func

// 对象属性配置
Object.defineProperties(user, {
  name: { value: "John", writable: false },
  surname: { value: "Smith", writable: false },
});

// user.name = "abc"; Error

let prop = {
  n: "prop",
  get Name() {
    return this.n;
  },

  set Name(value) {
    this.n = value;
  },
};

console.log(prop.Name); // prop

let newProp = {
  __proto__: prop,
};

newProp.Name = "newProp";
console.log(newProp.Name); // newProp

// prototype

class Animal {
  constructor(name) {
    this.name = name;
  }
}

class Rabbit extends Animal {
  // 私有属性
  #waterAmount = 0;
  hide() {
    alert(`${this.name} hides!`);
  }
}

// 集成相当于把prototype指向了父类
