# Crypto

## 分类

### 分组密码（block cipher）

主流应用算法,仅支持处理固定长度数据块,包含:

- DES
- AES

分组密码算法将明文数据序列按照**固定长度**进行分组,让后在同一密钥控制下用同一算法逐组迭代加密,从而将各个明文分组变成一个长度固定的密文分组的密码.
每次仅加密一个明文分组。如果明文因总长度超过分组长度而存在多个分组，那么分组密码算法会被迭代调用以逐个处理明文分组。

迭代的方法称为分组密码算法的模式.

### 流密码（stream cipher）

流密码是对数据流进行连续处理的一类算法

## 非对称加密/公钥密码

私有密钥（private key，简称私钥）和公共密钥（public key，简称公钥）.
RSA是世界上使用最广泛的公钥密码算法，Go语言的crypto/rsa包提供了对RSA算法的实现

```go
// $GOROOT/src/crypto/rsa/rsa.go
// RSA的公钥可以看成数对(E, N)，而私钥可以看出数对(D, N)
// 密文= 明文^E mod N (RSA密文就是对代表明文的数字的E次方求mod N的结果)
// 明文= 密文^D mod N
type PublicKey struct {
    N *big.Int // modulus 模数
    E int      // public exponent 指数
}

type PrivateKey struct {
    PublicKey            // 公钥部分
    D         *big.Int   // 私钥组件
    Primes    []*big.Int // N的质因数，有2个或2个以上元素
    ...
}
```

[RSA算法对待处理的数据长度是有要求的](https://www.cnblogs.com/caidi/p/14794952.html)

## 单向散列函数/one-way hash function

**单向散列函数**是一个接受**不定长输入**但产生**定长输出**的函数，这个定长输出被称为“摘要”（digest）或“指纹”（fingerprint）

Go标准库提供了多种主流单向散列函数标准的实现，包括

- MD5,不推荐,已被破解
- SHA-1,不推荐,已被破解
- SHA-256
- SHA-384
- SHA-512等,

拓展包golang.org/x/crypto/sha3中提供了最新的SHA-3标准的实现

## 消息认证码

单向散列函数虽然能辨别出数据是否被篡改，但却无法辨别出数据是不是伪装的。
在某些场合下，我们还需要对消息进行认证（Authentication），即校验消息的来源是不是我们所期望的。而用于解决这一问题的常见密码技术就是消息认证码（Message Authentication Code，MAC）

消息认证码是一种不仅能确认数据完整性，还能保证消息来自期望来源的密码技术。消息认证码技术是以通信双方共享密钥为前提的。对于任意长度的消息，我们都可以计算出一个固定长度的消息认证码数据，这个数据被称为MAC值。

实现

- crypto/hmac 一种基于单向散列函数的消息认证码实现，它实现了NIST发布的HMAC标准（the keyed-Hash Message Authentication Code）

对数据进行对称加密且携带MAC值的方式被称为“认证加密”,主要有:

- Encrypt-then-MAC：先用对称密码对明文进行加密，然后计算密文的MAC值
- Encrypt-and-MAC： 将明文用对称密码加密，并计算明文的MAC值
- MAC-then-Encrypt：先计算明文的MAC值，然后将明文和MAC值一起用对称密码加密。

## 数字签名

消息认证码虽然解决了消息发送者的身份认证问题，但由于采用消息认证码的通信双方共享密钥，因此对于一条通过了MAC验证的消息，通信双方依旧无法向第三方证明这条消息就是对方发送的.

数字签名用于解决上述问题.

使用公钥密码(公钥加密,私钥解密)的逆过程,即私钥签名,公钥验证.

在实际生产应用中，我们通常对消息的摘要进行签名。这是因为公钥密码加密算法本身很慢，如果对消息全文进行加密将非常耗时

RSA-PSS算法:采用对消息摘要进行签名，并在计算散列值时对消息加盐（salt）的方式来提高安全性（这样对同一条消息进行多次签名，每次得到的签名都不同）

## 随机数生成

Go密码学包crypto/rand提供了密码学级别的随机数生成器实现rand.Reader