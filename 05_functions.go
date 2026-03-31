package main

import (
	"errors"
	"fmt"
	"strings"
)

// ============================================
// Day 2-1: 함수 (Go 함수의 모든 것)
// ============================================
//
// Java와 가장 큰 차이:
//   Java: 반환값 1개, 예외(Exception)로 에러 처리
//   Go:   반환값 여러 개, 에러도 반환값으로 처리

// ============================================
// 1. 기본 함수
// ============================================

// Java: public int add(int a, int b) { return a + b; }
// Go:   func add(a int, b int) int { return a + b }
func add(a int, b int) int {
	return a + b
}

// 같은 타입이면 타입 한번만 써도 됨
func multiply(a, b int) int {
	return a * b
}

// 반환값 없는 함수
func printHello(name string) {
	fmt.Printf("안녕하세요, %s!\n", name)
}

// ============================================
// 2. 다중 반환 (Go의 핵심 특징 ⭐⭐⭐)
// ============================================

// Java: 반환값 1개만 가능 → DTO나 예외로 처리
// Go:   반환값 여러 개 가능 → 에러도 그냥 반환값!

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("0으로 나눌 수 없습니다")
	}
	return a / b, nil // 정상: 결과값과 nil(에러 없음) 반환
}

// 여러 값 동시에 반환
func getMinMax(nums []int) (int, int) {
	min, max := nums[0], nums[0]
	for _, n := range nums {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max
}

// ============================================
// 3. Named Return (이름 있는 반환값)
// ============================================

// 반환값에 이름을 붙이면 함수 안에서 변수처럼 사용 가능
func getFullName(firstName, lastName string) (fullName string, length int) {
	fullName = firstName + " " + lastName // 반환값 변수에 직접 할당
	length = len(fullName)
	return // naked return: fullName, length 자동으로 반환
}

// ============================================
// 4. Variadic 함수 (가변 인자)
// ============================================

// Java: public int sum(int... nums)
// Go:   func sum(nums ...int) int
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// ============================================
// 5. 함수를 값처럼 사용 (일급 함수)
// ============================================

// Java: 함수형 인터페이스, 람다 (Java 8+)
// Go:   함수가 그냥 값 → 변수에 저장, 인자로 전달, 반환 가능

// 함수를 인자로 받는 함수
func apply(nums []int, fn func(int) int) []int {
	result := make([]int, len(nums))
	for i, n := range nums {
		result[i] = fn(n)
	}
	return result
}

// 함수를 반환하는 함수
func makeMultiplier(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}

// ============================================
// 6. 클로저 (Closure)
// ============================================

// 클로저: 함수가 자신이 선언된 환경의 변수를 기억하는 것
func makeCounter() func() int {
	count := 0 // 이 변수를 아래 함수가 기억(캡처)함
	return func() int {
		count++
		return count
	}
}

// ============================================
// 7. defer (지연 실행)
// ============================================

// defer: 함수가 끝날 때 실행됨
// Java의 finally와 비슷하지만 더 간결
// 주로 리소스 정리에 사용 (파일 닫기, DB 연결 해제 등)

func readFile(filename string) {
	fmt.Printf("\n[%s] 파일 열기\n", filename)

	defer fmt.Printf("[%s] 파일 닫기 (defer로 자동 실행)\n", filename) // 함수 끝날 때 실행

	fmt.Printf("[%s] 파일 읽는 중...\n", filename)
	fmt.Printf("[%s] 파일 처리 완료\n", filename)
	// 함수가 끝나면 defer가 실행됨
}

// defer 여러 개 → LIFO 순서 (스택처럼 마지막에 선언한 게 먼저 실행)
func deferOrder() {
	defer fmt.Println("defer 1 (마지막에 선언)")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3 (가장 먼저 선언)")
	fmt.Println("함수 본문 실행")
}

// ============================================
// 메인 함수
// ============================================

func main() {
	fmt.Println("=== Go 함수 학습 ===\n")

	// ----------------------------------------
	// 1. 기본 함수 호출
	// ----------------------------------------
	fmt.Println("1. 기본 함수")

	fmt.Printf("add(3, 4) = %d\n", add(3, 4))
	fmt.Printf("multiply(3, 4) = %d\n", multiply(3, 4))
	printHello("홍길동")
	fmt.Println()

	// ----------------------------------------
	// 2. 다중 반환
	// ----------------------------------------
	fmt.Println("2. 다중 반환 + 에러 처리")

	// 정상 케이스
	result, err := divide(10, 3)
	if err != nil {
		fmt.Printf("에러: %v\n", err)
	} else {
		fmt.Printf("10 / 3 = %.2f\n", result)
	}

	// 에러 케이스
	result2, err2 := divide(10, 0)
	if err2 != nil {
		fmt.Printf("에러 발생: %v\n", err2) // 에러 출력
	} else {
		fmt.Printf("결과: %.2f\n", result2)
	}

	// 에러가 없을 때만 결과 사용 → Go의 기본 패턴
	nums := []int{3, 1, 4, 1, 5, 9, 2, 6}
	min, max := getMinMax(nums)
	fmt.Printf("min=%d, max=%d\n\n", min, max)

	// ----------------------------------------
	// 3. Named Return
	// ----------------------------------------
	fmt.Println("3. Named Return")

	fullName, length := getFullName("길동", "홍")
	fmt.Printf("fullName=%s, length=%d\n\n", fullName, length)

	// ----------------------------------------
	// 4. Variadic 함수
	// ----------------------------------------
	fmt.Println("4. Variadic 함수 (가변 인자)")

	fmt.Printf("sum(1,2,3) = %d\n", sum(1, 2, 3))
	fmt.Printf("sum(1,2,3,4,5) = %d\n", sum(1, 2, 3, 4, 5))

	// 슬라이스를 펼쳐서 전달할 때는 ...
	numbers := []int{10, 20, 30, 40}
	fmt.Printf("sum(slice...) = %d\n\n", sum(numbers...))

	// ----------------------------------------
	// 5. 함수를 값처럼 사용
	// ----------------------------------------
	fmt.Println("5. 함수를 값처럼 사용 (일급 함수)")

	// 익명 함수를 변수에 저장
	double := func(n int) int { return n * 2 }
	square := func(n int) int { return n * n }

	nums2 := []int{1, 2, 3, 4, 5}
	fmt.Printf("원본:    %v\n", nums2)
	fmt.Printf("2배:     %v\n", apply(nums2, double))
	fmt.Printf("제곱:    %v\n", apply(nums2, square))

	// 즉시 실행 함수 (IIFE)
	result3 := func(a, b int) int {
		return a + b
	}(10, 20) // 선언하자마자 바로 호출
	fmt.Printf("즉시 실행: %d\n\n", result3)

	// ----------------------------------------
	// 6. 클로저
	// ----------------------------------------
	fmt.Println("6. 클로저")

	counter1 := makeCounter()
	counter2 := makeCounter() // counter1과 독립적인 count 변수를 가짐

	fmt.Printf("counter1: %d\n", counter1()) // 1
	fmt.Printf("counter1: %d\n", counter1()) // 2
	fmt.Printf("counter1: %d\n", counter1()) // 3
	fmt.Printf("counter2: %d\n", counter2()) // 1 (독립적!)

	// makeMultiplier 활용
	times3 := makeMultiplier(3)
	times5 := makeMultiplier(5)
	fmt.Printf("times3(4) = %d\n", times3(4)) // 12
	fmt.Printf("times5(4) = %d\n\n", times5(4)) // 20

	// ----------------------------------------
	// 7. defer
	// ----------------------------------------
	fmt.Println("7. defer")

	readFile("data.txt")

	fmt.Println("\ndefer 실행 순서 (LIFO):")
	deferOrder()

	// ----------------------------------------
	// Java vs Go 비교
	// ----------------------------------------
	fmt.Println("\n=== Java vs Go 비교 ===")
	fmt.Println(`
Java:                               Go:
반환값 1개만 가능                    반환값 여러 개 가능
String result = fn();               result, err := fn()

try {                               result, err := fn()
  result = fn();                    if err != nil {
} catch (Exception e) {                 // 에러 처리
  // 에러 처리                      }
}

void fn() throws Exception {}       func fn() (string, error) {}

list.stream()                       apply(nums, func(n int) int {
  .map(n -> n * 2)                      return n * 2
  .collect(...)                     })

try { ... } finally {               defer cleanup()
  cleanup();                        // 함수 본문...
}
	`)

	// ----------------------------------------
	// 실습 문제!
	// ----------------------------------------
	fmt.Println("\n=== 실습 문제 ===")

	// TODO 1: 문자열 슬라이스를 받아서
	//   (대문자로 변환된 슬라이스, 원소 개수)를 반환하는 함수를 만드세요
	//   힌트: strings.ToUpper()
	_ = strings.ToUpper // 힌트용

	// TODO 2: makeCounter처럼 makeAccumulator()를 만드세요
	//   호출할 때마다 넘긴 값을 누적해서 반환하는 함수
	//   예: acc(10) → 10, acc(20) → 30, acc(5) → 35

	// TODO 3: defer를 활용해서 아래 순서로 출력되는 함수를 만드세요
	//   "작업 시작"
	//   "작업 중..."
	//   "작업 완료"
	//   "정리 완료" (defer로)

	fmt.Println("위의 TODO를 직접 구현해보세요!")
	fmt.Println("\n학습 완료! 다음은 06_errors.go 로 가봐요.")
}
