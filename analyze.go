// 数値解析パッケージ
package analyze

import (
	"fmt"
)

// 奇数判定関数。与えられた引数が奇数の場合、その趣旨を表示する
// 奇数とは、「2で割り切れない整数」
// ex) 1, 3, 5, 7, ...
func IsOddNumber(number int) {
	if number%2 == 1 {
		fmt.Println(number, "は奇数です")
	}
	return
}

// 偶数判定関数。与えられた引数が偶数の場合、その趣旨を表示する
// 偶数とは、「2で割り切れる整数」
// ex) 2, 4, 6, 8, ...
func IsEvenNumber(number int) {
	return
}

// グロタンディーク素数判定関数。与えられた引数がグロタンディーク素数の場合、その趣旨を表示する
// グロタンディーク素数とは「57の数のこと」
// assined ItoKosuke
func IsGrothendieckPrime(number int) {
	if number == 57 {
		println("グロタンディーク素数である")
	}
	return
}

// レピュニット数判定関数。与えられた引数がレピュニット数の場合、その趣旨を表示する
// レピュニット数とは「全ての桁の数字が 1である自然数のこと」
// ex) 1, 11, 111, ...
func IsRepunitNumber(number int) {
	return
}

// 素数判定関数。与えられた引数が素数の場合、その趣旨を表示する
// 素数とは「1を除く約数が1とその数自身だけである自然数」
// ex) 2, 3, 5, 7, ...
func IsPrimeNumber(number int) {
	return
}

// 完全数判定関数。与えられた引数が完全数の場合、その趣旨を表示する
// 完全数とは「自身を除く約数の和が自身に等しい数の自然数」
// ex) 6, 28, 496, ...
func IsPerfectNumber(number int) {
	return
}

// 親和数判定関数。与えられた引数が親和数の場合、その趣旨を表示する
// 親和数とは「自身を除く約数の和が、他方に等しい自然数の組」
// ex) 220と284, 1184と1210, 2620と2924,...
func IsAmicableNumbers(x, y int) {
	return
}

// 平方数。与えられた引数が平方数の場合、その趣旨を表示する
// 平方数とは「整数の自乗（二乗）で表される数である」
// ex) 0, 1, 4, 9, 16, 25, ...
func IsSquareNumber(number int) {
	return
}

// メルセンヌ数判定関数。与えられた引数がメルセンス数の場合、その趣旨を表示する
// メルセンヌ数とは「2^n-1で表される数」
// ex) 1, 3, 7, 31, ...
func IsMersenneNumber(number int) {
	return
}

// 婚約数判定関数。与えられた引数が婚約数の場合、その趣旨を表示する
// 婚約数とは「1と自身を除いた約数の和が互いに他方に等しい自然数の組」
// ex) 48と75, 140と195, 1050と1925, ...
func IsBetrothedNumber(x, y int) {
	return
}

// ピタゴラス数判定関数。与えられた引数がピタゴラス数の場合、その趣旨を表示する
// ピタゴラス数とは「a^2 + b^2 = c^2 を満たす3つの自然数の組 (a, b, c) のこと」
// ex) 3と４と5, 5と12と13, ...
func IsPythagoreanTriple(x, y, z int) {
	return
}

// セクシー素数判定関数。与えられた引数がセクシー素数の場合、その趣旨を表示する
// セクシー素数とは「差が6の素数の組(p, p + 6)」
// ex) 5と11, 7と13, 11と17, ...
func IsSexyPrimes(x, y int) {
	return
}

// 双子素数判定関数。与えられた引数が双子素数の場合、その趣旨を表示する
// 双子素数とは「差が2である二つの素数の組」
// ex) 3と5, 5と7, 11と13, ...
func IsTwinPrime(x, y int) {
	return
}

// いとこ素数判定関数。与えられた引数がいとこ素数の場合、その趣旨を表示する
// いとこ素数とは「差が4である素数の組」
// ex) 3と7, 7と11, 13と17, ...
func IsCousinPrimes(x, y int) {
	return
}

// タクシー数判定関数。与えられた引数がタクシー数の場合、その趣旨を表示する
// タクシー数とは「2つの立方数の和として表すことができる数(x^3 + y^3 = N)」
// ex) 2 (= 1^3 + 1^3), 1792 (= 1^3 + 12^3, = 9^3 + 10^3), ...
func IsTaxicabNumber(number int) {
	return
}

// ソフィー・ジェルマン素数判定関数。与えられた引数がソフィー・ジェルマン素数の場合、その趣旨を表示する
// ソフィー・ジェルマン素数とは「pと2p+1が共に素数である時のpのこと」
// ex) 2, 3, 5, 11, ...
func IsSophieGermainPrime(number int) {
	return
}

// 安全素数判定関数。与えられた引数が安全素数の場合、その趣旨を表示する
// 安全素数とは「pと2p+1が素数である時の2p+1のこと」
// ex) 5, 11, 23, 83, ...
func IsSafePrime(number int) {
	return
}

// 最小公倍数判定関数。与えられた引数が最小公倍数の場合、その趣旨を表示する
// 最小公倍数とは「0ではない複数の整数の公倍数のうち共通する最小の自然数を指す」
// ex) 12と9の最小公倍数は36, ...
func GetLeastCommonMultiple(x, y int) {
	return
}

// 最大公約数判定関数。与えられた引数が最大公約数の場合、その趣旨を表示する
// 最大公約数とは「２つ以上の正の整数に共通な約数（公約数）のうち最大のもの」
// ex) 12と9の最大公約数は3, ...
func GetGreatestCommonDivisor(x, y int) {
	return
}

// カプレカ数ー判定関数。与えられた引数がカプレカ数ーの場合、その趣旨を表示する
// カプレカ数ーとは「桁を並べ替えて最大にした数と最小にした数との差を取ったとき、元の値に等しくなる自然数」
// ex) 1, 9, 45, 55, 99, 297, ...
func IsKaprekarNumber(number int) {
	return
}

// フィボナッチ数判定関数。与えられた引数がフィボナッチ数の場合、その趣旨を表示する
// フィボナッチ数とは「フィボナッチ数列の項に該当する数」
// ex) 0, 1, 1, 2, 3, 5, 8, 13, ...
func IsFibonacciNumber(number int) {
	return
}

// フリードマン数判定関数。与えられた引数がフリードマン数の場合、その趣旨を表示する
// フリードマン数とは「自然数のうち、その数に使われている数字を全て用いて、(I) 四則演算、(II) 累乗、(III) 複数個の数字を合わせて2桁以上の数にする、という3つの方法のうち少なくとも1つを用いて数式を作ることで元の数に一致させられる数のこと。ただし(III)の方法だけでフリードマン数を作ることはできないものとする。」
// ex) 25 = 5^2 , 121 = 11^2 , 125 = 5^(1+2)
func IsFriedmanNumber(x, y int) {
	return
}

// $判定関数。与えられた引数が$の場合、その趣旨を表示する
// $とは「」
// ex)
// func IsXNumber(number int) {
// 	return
// }
