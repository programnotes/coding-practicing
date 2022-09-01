package jzoffer

import "math"

// 当被除数大于除数时，继续比较判断被除数是否大于除数的2倍，如果是，
// 则继续判断被除数是否大于除数的4倍、8倍等。如果被除数最多大于除数的2k倍，
// 那么将被除数减去除数的2k倍，然后将剩余的被除数重复前面的步骤。
// 由于每次将除数翻倍，因此优化后的时间复杂度是O（logn）。
// 下面以15/2为例讨论计算的过程。15大于2，也大于2的2倍（即4），还大于2的4倍（即8），
// 但小于2的8倍（即16）。于是先将15减去8，还剩余7。由于减去的是除数的4倍，减去这部分对应的商是4。
// 接下来对剩余的7和除数2进行比较，7大于2，大于2的2倍（即4），但小于2的4倍（即8），于是将7减去4，
// 还剩余3。这一次减去的是除数2的2倍，对应的商是2。然后对剩余的3和除数2进行比较，3大于2，
// 但小于2的2倍（即4），于是将3减去2，还剩余1。这一次减去的是除数的1倍，对应的商是1。
// 最后剩余的数字是1，比除数小，不能再减去除数了。于是15/2的商是4+2+1，即7。
// 上述讨论假设被除数和除数都是正整数。如果有负数则可以将它们先转换成正数，
// 计算正数的除法之后再根据需要调整商的正负号。例如，如果计算-15/2，则可以先计算15/2的商是4+2+1，即7。

// 上述讨论假设被除数和除数都是正整数。如果有负数则可以将它们先转换成正数，
// 计算正数的除法之后再根据需要调整商的正负号。例如，如果计算-15/2，则可以先计算15/2，得到的商是7。
// 由于被除数和除数中有一个负数，因此商应该是负数，于是商应该是-7。

// 将负数转换成正数存在一个小问题。
// 对于32位的整数而言，最小的整数是-2^31，最大的整数是2^31-1。因此，如果将-231转换为正数则会导致溢出。
// 由于将任意正数转换为负数都不会溢出，因此可以先将正数都转换成负数，用前面优化之后的减法计算两个负数
// 的除法，然后根据需要调整商的正负号。最后讨论可能的溢出。由于是整数的除法并且除数不等于0，因此商的
// 绝对值一定小于或等于被除数的绝对值。
// 因此，int型整数的除法只有一种情况会导致溢出，即（-2^31）/（-1）。这是因为最大的正数为2^31-1，2^31超出了正数的范围
func _divideCore(dividend, divisor int32) int32 {
	halfMin := (math.MinInt32 >> 1)
	var res int32
	for dividend <= divisor {
		val := divisor
		quotient := int32(1)                               /*商*/
		for val >= int32(halfMin) && dividend <= val+val { //负数不要越界溢出
			val += val
			quotient += quotient
		}
		res += quotient
		dividend -= val
	}
	return res
}

func _divide(dividend, divisor int32) int32 {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	nagative := 2
	if dividend > 0 {
		dividend = -dividend
		nagative--
	}
	if divisor > 0 {
		divisor = -divisor
		nagative--
	}
	res := _divideCore(dividend, divisor)
	if nagative == 1 {
		return -res
	}
	return res
}