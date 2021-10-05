# Set

Set 是集合，其中实现类 `TreeSet` 内元素是有序的，底层基于 `TreeMap` 实现。

## Set 中的元素不可重复的实现

这主要是由 Map 的 put 机制来保证的。HashSet 和 TreeSet 底层的存储都是基于 Map 的 key 。Map 的 `put(E key,T val)` 方法特点是，若 key 存在，则旧的 val 被替换，同时返回修改前的值。HashSet 和 TreeSet 的 add 方法实现相同：

```java
public boolean add(E e) {
    return map.put(e, PRESENT)==null;// PRESENT = new Object();所以值被替换并无大碍
}
/**
 * If the map previously contained a mapping for
 * the key, the old value is replaced by the specified value.  
 */
V put(K key, V value);
```

# Map



## TreeMap

key 值有序，基于红黑树实现。

## HashMap



## ConcurrentHashMap

