# 多个Promise

使用`Node-redis`操作 `redis` 时，redis 命令的执行结果在 nodejs 中都体现为回调。

```javascript
redis.set('foo','bar')
redis.get('foo',(err,reply) => {
    console.log(reply);//bar
})
```

用回调的形式得到命令执行的结果，对于单条命令还可以应对，当需要连续执行多条命令时，便会遇到“callback  hell”，甚至，需要对数组中的每个值都执行一次`redis`命令，并一次性获得所有执行结果用于后续操作时，使用回调不单单会使代码变得难以读懂，执行结果甚至和预期不同，导致运行结果有逻辑错误。

## Promise

代表了未来某个将要发生的事件，`Promise` 对象将异步操作以同步操作的流程表现出来，避免层层嵌套的回调函数。

Promise对象的详细解释，参考[JavaScript标准参考教程](http://javascript.ruanyifeng.com/advanced/promise.html) 或者 [ECMAScript 6 文档](http://www.nodeclass.com/api/ECMAScript6.html#promise)，此处只说明所遇到的问题。

页面中表格数据的最后一列，需要挨个从 redis 中查找出来。

```js
function func (id) {
    return new Promise((resolve, reject) => {
        db.hgetall('courses:' + id, (err, reply) => {
            if (reply) {
                resolve({rest: reply.capcatity - reply.selected})
            }
        })
    })
}
let promises = preDefined.ids.map(id => {
    return func(id)
})
Promise.all(promises).then(results => {
    preDefined.rest = results
    for (let index = 0; index < times; index++) {
        let data = courseItem()
        data.courseId = preDefined.ids[index]
        data.courseName = preDefined.courseNames[index]
        data.position = preDefined.position[index]
        data.rest = preDefined.rest[index].rest
        datas.push(data)
    }
    res.status(200).send({datas:datas})
}).catch(reason => {
    console.log('reason:',reason)
})
```

最开始使用 `for` 循环去多次调用`db.hgetall`,再在其回调函数中将返回值 `push` 到一个数组 `result`，再对此数组进行操作。但结果却是`result`始终为 `undefined`,后续操作也就无从谈起。

```js
let result = []
for(let i = 0;i<preDefined.ids.length;i++){
    db.hgetall('courses:' + preDefined.ids[i], (err, reply) => {
        if (reply) {
            result.push({rest: reply.capcatity - reply.selected})
        }
    })
}
preDefined.rest = result //preDefined.rest.length = 0
```

这个问题的关键在于，就算循环执行`db.hgetall`，在每个回调函数里，也做不到对 `result`  不断添加元素，因为程序使用 `result`  时，回调函数还没有开始执行。

### Promise.all([p1,p2,p3])

`Promise.all` 函数用于将多个 Promise 实例包装成一个实例，其返回值仍然是一个 Promise 实例 `p`，p的状态由p1,p2,p3决定，分成两种情况。

- 只有p1、p2、p3的状态都变成fulfilled，p的状态才会变成fulfilled，此时p1、p2、p3的返回值组成一个数组，传递给p的回调函数。
- 只要p1、p2、p3之中有一个被rejected，p的状态就变成rejected，此时第一个被reject的实例的返回值，会传递给p的回调函数。

### async 函数

async 函数也是用来**取代回调函数**的一种方法。

在函数名前加 `async` 关键字，就表明此函数内部有异步操作，应该返回一个 `Promise` 对象，前面用 `await` 关键字注明，当函数执行时，遇到 `await` 就会先返回该 Promise 对象，等到触发的异步操作完成，再接着执行函数体内后面的语句。

```js
function timeout(ms) {
  return new Promise((resolve) => {
    setTimeout(resolve, ms);
  });
}

async function asyncValue(value) {
  await timeout(50);//先执行timeout函数而不返回，若没有await，则直接返回
  return value;
}
```
