package main

import "fmt"

// ============================================
// Day 1-1: Go 기본 문법
// ============================================

// 전역 변수 선언 (함수 밖에서는 var 키워드 필수)
var globalVar = "전역 변수"

// const 선언 (상수)
const PI = 3.14
const AppName = "Go학습앱"

func main() {
	fmt.Println("=== Go 기본 문법 학습 ===\n")

	// ----------------------------------------
	// 1. 변수 선언 방법
	// ----------------------------------------
	fmt.Println("1. 변수 선언 방법")

	// 방법 1: var 키워드 사용 (타입 명시)
	// var 변수명 타입 = 값
	var name string = "홍길동"
	var age int = 30
	fmt.Printf("이름: %s, 나이: %d\n", name, age)

	// 방법 2: var 키워드 사용 (타입 추론)
	// var 변수명 = 값 (타입을 Go가 알아서 추론)
	var city = "서울" // string으로 자동 추론
	var count = 100 // int로 자동 추론
	fmt.Printf("도시: %s, 카운트: %d\n", city, count)

	// 방법 3: := 단축 선언 ⭐⭐⭐ 가장 많이 씀!
	// := 는 "선언 + 할당"을 한번에!
	// 변수명 := 값 (타입도 자동 추론)
	email := "test@example.com" // 이건 var email string = "test@example.com" 와 같음
	isActive := true
	fmt.Printf("이메일: %s, 활성화: %t\n", email, isActive)

	// := 와 = 의 차이
	email = "changed@example.com" // = 는 이미 선언된 변수에 값만 재할당
	// email := "error@example.com" // ❌ 에러! email은 이미 선언됨 (중복 선언 불가)

	// ❌ 주의: 함수 밖에서는 := 사용 불가!
	// globalEmail := "error@example.com"  // 이건 에러!
	// 함수 밖에서는 var 키워드 필수!

	// 여러 변수 한번에 선언
	var x, y int = 10, 20
	a, b, c := 1, "hello", true // 타입이 달라도 OK
	fmt.Printf("x=%d, y=%d\n", x, y)
	fmt.Printf("a=%d, b=%s, c=%t\n\n", a, b, c)

	// ----------------------------------------
	// 2. 기본 타입
	// ----------------------------------------
	fmt.Println("2. 기본 타입")

	// 정수형
	var num1 int = 42           // 기본 정수 타입 (Java의 int)
	var num2 int64 = 9999999999 // 64비트 정수 (Java의 long)
	var num3 int32 = 100        // 32비트 정수
	fmt.Printf("int: %d, int64: %d, int32: %d\n", num1, num2, num3)

	// 실수형
	var price float64 = 99.99 // 64비트 실수 (Java의 double)
	var rate float32 = 3.14   // 32비트 실수 (Java의 float)
	fmt.Printf("float64: %.2f, float32: %.2f\n", price, rate)

	// 문자열
	var message string = "안녕하세요"
	var multiLine string = `여러 줄
	문자열도
	가능합니다` // 백틱(`)으로 여러 줄 문자열
	fmt.Printf("문자열: %s\n", message)
	fmt.Printf("여러줄:\n%s\n", multiLine)

	// 불린
	var isGoFun bool = true
	var isJavaOld bool = false
	fmt.Printf("Go 재밌나요? %t, Java 오래됐나요? %t\n\n", isGoFun, isJavaOld)

	// ----------------------------------------
	// 3. 상수 (const)
	// ----------------------------------------
	fmt.Println("3. 상수")

	const MaxUsers = 1000
	const ServerURL = "https://api.example.com"

	fmt.Printf("PI: %.2f\n", PI)
	fmt.Printf("최대 사용자: %d\n", MaxUsers)
	fmt.Printf("서버 URL: %s\n\n", ServerURL)

	// ❌ const는 변경 불가
	// MaxUsers = 2000  // 에러!

	// ----------------------------------------
	// 4. 제로값 (Zero Value)
	// ----------------------------------------
	fmt.Println("4. 제로값 (초기화 안 하면 자동으로 들어가는 값)")

	var zeroInt int       // 0
	var zeroFloat float64 // 0.0
	var zeroString string // "" (빈 문자열)
	var zeroBool bool     // false
	var zeroPointer *int  // nil

	fmt.Printf("int 제로값: %d\n", zeroInt)
	fmt.Printf("float64 제로값: %.1f\n", zeroFloat)
	fmt.Printf("string 제로값: '%s'\n", zeroString)
	fmt.Printf("bool 제로값: %t\n", zeroBool)
	fmt.Printf("포인터 제로값: %v\n\n", zeroPointer)

	// ----------------------------------------
	// 5. 타입 변환 (Casting)
	// ----------------------------------------
	fmt.Println("5. 타입 변환")

	// Go는 자동 형변환이 없음! 명시적으로 해야 함
	var i int = 42
	var f float64 = float64(i) // int -> float64
	var u uint = uint(f)       // float64 -> uint

	fmt.Printf("int: %d, float64: %.1f, uint: %d\n\n", i, f, u)

	// ❌ 주의: 자동 변환 안 됨!
	// var wrong float64 = i  // 에러!

	// ----------------------------------------
	// 6. if 문 (괄호 없음!)
	// ----------------------------------------
	fmt.Println("6. if 문")

	score := 85

	// Java: if (score >= 90)
	// Go:   if score >= 90  ← 괄호 없음!
	if score >= 90 {
		fmt.Println("A 학점")
	} else if score >= 80 {
		fmt.Println("B 학점")
	} else {
		fmt.Println("C 학점")
	}

	// if문에서 변수 선언 가능 (scope는 if 블록 안에만)
	if value := score * 2; value > 150 {
		fmt.Printf("값이 큽니다: %d\n\n", value)
	}
	// fmt.Println(value)  // 에러! value는 if 블록 밖에서 사용 불가

	// ----------------------------------------
	// 7. 간단한 출력 정리
	// ----------------------------------------
	fmt.Println("7. 출력 함수들")

	fmt.Println("Println: 자동 줄바꿈")
	fmt.Print("Print: 줄바꿈 없음 ")
	fmt.Print("이어서 출력\n")
	fmt.Printf("Printf: 포맷 지정 가능 - 이름: %s, 나이: %d\n", "김철수", 25)

	// ----------------------------------------
	// 실습 문제!
	// ----------------------------------------
	fmt.Println("\n=== 실습 문제 ===")
	fmt.Println("아래 주석을 풀고 직접 코드를 작성해보세요!")

	// TODO 1: 본인의 이름, 나이, 이메일을 변수로 선언하고 출력하세요 (:= 사용)
	printMyInfo() // 함수 호출

	// TODO 2: 상수로 본인이 좋아하는 숫자를 선언하고 출력하세요
	printFavoriteNumber()

	// TODO 3: 점수가 60 이상이면 "합격", 아니면 "불합격" 출력하는 if문 작성
	checkPassOrFail(75)

	fmt.Println("\n학습 완료! 다음은 02_loops.go 파일을 실행해보세요.")
}

// ============================================
// 실습 문제 함수들 (main 밖에 선언)
// ============================================
// 주의: 이건 "함수"입니다! (메서드 아님)
// 메서드는 struct에 붙은 것, 이건 독립적인 함수

// TODO 1 구현: 개인 정보 출력 함수
func printMyInfo() {
	// := 사용해서 변수 선언 + 할당
	myName := "김택규"
	myAge := 28
	myEmail := "taekgyu@example.com"

	fmt.Println("\n[TODO 1 실행]")
	fmt.Printf("이름: %s, 나이: %d, 이메일: %s\n", myName, myAge, myEmail)
}

// TODO 2 구현: 좋아하는 숫자 출력 함수
func printFavoriteNumber() {
	const FavoriteNumber = 7 // 상수 선언

	fmt.Println("\n[TODO 2 실행]")
	fmt.Printf("내가 좋아하는 숫자: %d\n", FavoriteNumber)
}

// TODO 3 구현: 합격/불합격 체크 함수
func checkPassOrFail(score int) {
	fmt.Println("\n[TODO 3 실행]")
	fmt.Printf("점수: %d점 -> ", score)

	if score >= 60 {
		fmt.Println("합격!")
	} else {
		fmt.Println("불합격!")
	}
}
