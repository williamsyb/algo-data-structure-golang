### 项目简介

介绍数据结构与算法，涵盖字符串、数组、链表、队列、栈、树、图等一系列的数据结构，同时也会注重算法的应用。在理论知识的基础上，注重算法的go语言实现。

本教程注重数据结构与算法基础，目的不为刷题，但为了照顾部分同学，会加入经典的数据结构与算法面试题目，很多题目都是大厂喜欢问的面试题目。

### 适宜人群
 - 非计算机专业，但是对计算机感兴趣的同学
 - 数据结构与算法基础不扎实的同学
 - 希望拿到满意offer，有面试需求的同学
 - 了解算法理论，但是不知道如何用go语言实现的同学

### 你会得到的收获
 - 熟悉数据结构与算法
 - 掌握go编程语言
 - 一份满意的offer


### 大话Go语言

2009年，Go语言诞生了！正如历史上其他一切新事物一样，Go也有它自己的梦想与使命。2009年，我自认为不管是在计算机软件发展史、还是在人类发展史上，都是一个相当重要的年份。2008年金融危机爆发，经济下行全球产生信任危机。从此绝大多数行业开始走向转折，互联网行业跟随时代的步伐，异军突起，开始了草莽式的发展。

正是在这样的历史时刻，Go语言作为Google的金宝贝（Golden son）在一批大牛的怀抱中诞生了，诞生之初就自诩为互联网时代的C语言。然而，也正是这一年，在另一边的芬兰赫尔辛基，中本村用C++语言实现了比特币的代码，挖出了世界上第一枚比特币。 历史总是惊人的巧合，或许在当时没人能够想到，这两个完全不同的事情，对人类发展史有如此重大的影响。 现如今，Go跟C++作为区块链底层技术领域的双雄，正在改变历史。 或许今天，这种改变你可能还没意识到，或者没体会到，但我要告诉你，Time is changed！

当然，Go语言诞生的本意并不是成为区块链底层技术开发的主流语言，就像我前面提到的，它有自己的梦想与使命。08年前后，随着互联网的高速发展，对网站的性能提出了越来越多的挑战，所以并发与分布式成为了工程师们关注的焦点。 此外不同与传统的大型服务器，互联网行业一方面为了适应高速发展的需求，同时降低资源成本，多核化与集群化是互联网时代的典型特点。Go语言作为互联网时代诞生的语言，一开始就顺应时代，天然支持高并发。


经过这么多年的发展之后，技术工程师们希望在保持高性能的基础上，在开发使用方面也能够保持足够的轻量、简单，同时又能适应大规模的软件工程开发工作。所以工程师们带着这些要求，重新出发，在借鉴前人的基础上，又推翻了之前存在的很多东西，另辟蹊径，终于在2009年，Go语言出现在了我们面前。

不同于其他绝大多数语言，由于Go语言本身就出身在大户人家，所以天生就发育良好。到2013年，Docker的出现吹响了go语言开疆扩土的号角。此后随着k8s、 docker-swarm 等一批基于pass平台的技术出现之后，Go语言正式成为了云计算领域的绝对主角。好巧不巧，同样也是在2013年，ethereum基金会成立，融资开发ethereum平台，2015年随着ethereum的发布上线，ethereum成为了比特币之后第二影响力的公链。而让人兴奋的是，ethereum是基于Go语言开发的，从此go语言跟C++一样成为了区块链底层开发技术的不二之选。

今年是2018年，Go语言已经应用在了很多产品当中，同时也有很多杀手级的应用。

谁也无法阻止历史的车轮，未来的事情会是怎么样，我们谁都无法预测。但，在当下，Go语言作为高并发与分布式系统领域最炙手可热的语言，值得我们去学习，去钻研，去探讨！

### 代码目录
- examples
    - [RPC实战](https://github.com/csunny/argo/tree/master/examples/rpc_example)
    - [GRPC实战](https://github.com/csunny/argo/tree/master/examples/grpc)

- src
    - 数据结构与算法
        - [快速排序](https://github.com/csunny/argo/tree/master/src/argothrim/qsort)
        - [切片操作与反转](https://github.com/csunny/argo/blob/master/src/common/array.go)
        - [递归实现](https://github.com/csunny/argo/blob/master/src/common/recursion.go)
        - [二叉搜索树](https://github.com/csunny/argo/tree/master/src/tree)
        - [链表](https://github.com/csunny/argo/tree/master/src/linklist)
        - [图](https://github.com/csunny/argo/tree/master/src/graph)
        - [队列](https://github.com/csunny/argo/tree/master/src/queue)
        - [栈](https://github.com/csunny/argo/tree/master/src/stack)
        - [迭代器实现](https://github.com/csunny/argo/tree/master/src/iterator)
        - [哈希表](https://github.com/csunny/argo/tree/master/src/hashtable)
    - 分布式算法实现
        - [POS算法实现](https://github.com/csunny/argo/tree/master/src/pos)
        - [DPOS算法实现](https://github.com/csunny/argo/tree/master/src/dpos)
        - [DHT(分布式hash表实现)](https://github.com/csunny/argo/tree/master/src/libs/kademlia)
    - 分布式网络(P2P)
        - [P2P对等网络]()
        - [一个基于P2P对等网络的区块链实现](https://github.com/csunny/argo/tree/master/src/p2p)
### 文档地址
- https://xiaozhuanlan.com/argo


### Licence
argo code is licensed under the Apache
