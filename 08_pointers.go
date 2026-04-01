package main

import "fmt"

// ============================================
// Day 2-4: 포인터 (Pointer)
// ============================================
//
// 포인터란?
//   변수의 메모리 주소를 저장하는 변수
//
//   일반 변수: 값 자체를 저장
//   포인터:    값이 있는 메모리 주소를 저장
//
// 기호 2개만 기억하면 됨:
//   &  →  변수의 주소를 가져옴
//   *  →  주소에 있는 값을 가져옴 (역참조)

// ============================================
// 메인 함수
// ============================================

func main() {
	fmt.Println("=== Go 포인터 학습 ===\n")

	// ----------------------------------------
	// 1. 포인터 기본 개념
	// ----------------------------------------
	fmt.Println("1. 포인터 기본 개념")

	x := 10

	ptr := &x // &x = x의 메모리 주소

	fmt.Printf("x의 값:         %d\n", x)
	fmt.Printf("x의 메모리 주소: %p\n", &x)
	fmt.Printf("ptr에 저장된 값: %p  (x의 주소)\n", ptr)
	fmt.Printf("ptr이 가리키는 값: %d  (*ptr = 역참조)\n\n", *ptr)

	// 메모리 그림:
	//
	//  변수명   주소              값
	//  ──────────────────────────────
	//  x     → 0xc0000b4008  →  10
	//  ptr   → 0xc0000b4010  →  0xc0000b4008  (x의 주소를 저장)
	//
	//  *ptr → ptr이 가진 주소(0xc0000b4008)로 찾아가서 값(10)을 꺼냄

	// 포인터로 원본 값 변경
	*ptr = 99 // ptr이 가리키는 주소의 값을 99로 변경
	fmt.Printf("*ptr = 99 실행 후 x의 값: %d\n\n", x) // x도 99로 바뀜!

	// ----------------------------------------
	// 2. 핵심 차이: 값 전달 vs 포인터 전달
	// ----------------------------------------
	fmt.Println("2. 값 전달 vs 포인터 전달 (가장 중요! ⭐)")

	// ── 값 전달 ──────────────────────────────
	// 함수에 값을 넘기면 복사본이 전달됨
	// → 함수 안에서 변경해도 원본은 그대로

	score := 50
	fmt.Printf("[값 전달] 함수 호출 전: %d\n", score)
	addTenByValue(score)
	fmt.Printf("[값 전달] 함수 호출 후: %d  ← 원본 그대로!\n\n", score)

	// ── 포인터 전달 ───────────────────────────
	// 함수에 주소를 넘기면 원본을 직접 변경 가능

	score2 := 50
	fmt.Printf("[포인터 전달] 함수 호출 전: %d\n", score2)
	addTenByPointer(&score2) // &score2 = score2의 주소를 넘김
	fmt.Printf("[포인터 전달] 함수 호출 후: %d  ← 원본이 바뀜!\n\n", score2)

	// ----------------------------------------
	// 3. Struct에서 값 전달 vs 포인터 전달
	// ----------------------------------------
	fmt.Println("3. Struct에서 값 전달 vs 포인터 전달")

	user1 := Person{Name: "홍길동", Age: 20}

	fmt.Printf("[값 전달] 전: %+v\n", user1)
	growUpByValue(user1)
	fmt.Printf("[값 전달] 후: %+v  ← 원본 그대로!\n\n", user1)

	user2 := Person{Name: "홍길동", Age: 20}

	fmt.Printf("[포인터 전달] 전: %+v\n", user2)
	growUpByPointer(&user2)
	fmt.Printf("[포인터 전달] 후: %+v  ← 원본이 바뀜!\n\n", user2)

	// ----------------------------------------
	// 4. 값 리시버 vs 포인터 리시버
	// ----------------------------------------
	fmt.Println("4. 값 리시버 vs 포인터 리시버")

	// 값 리시버: 복사본에서 실행 → 원본 안 바뀜
	p1 := Person{Name: "김철수", Age: 25}
	fmt.Printf("[값 리시버] 전: %+v\n", p1)
	p1.BirthdayByValue()
	fmt.Printf("[값 리시버] 후: %+v  ← 원본 그대로!\n\n", p1)

	// 포인터 리시버: 원본에서 실행 → 원본 바뀜
	p2 := Person{Name: "이영희", Age: 25}
	fmt.Printf("[포인터 리시버] 전: %+v\n", p2)
	p2.BirthdayByPointer()
	fmt.Printf("[포인터 리시버] 후: %+v  ← 원본이 바뀜!\n\n", p2)

	// ----------------------------------------
	// 5. 포인터를 쓰는 이유 2: 복사 비용 절감
	// ----------------------------------------
	fmt.Println("5. 복사 비용 절감")

	big := BigStruct{Data: [100]int{1, 2, 3}} // 800바이트짜리 struct

	fmt.Println("[값 전달]    → BigStruct 전체(800바이트)를 복사해서 전달 (느림)")
	processByValue(big)

	fmt.Println("[포인터 전달] → 주소(8바이트)만 전달 (빠름)")
	processByPointer(&big)
	fmt.Println()

	// ----------------------------------------
	// 6. 포인터와 nil
	// ----------------------------------------
	fmt.Println("6. 포인터와 nil")

	var nilPtr *Person // 선언만 하면 nil (아무 주소도 안 가리킴)
	fmt.Printf("nilPtr: %v\n", nilPtr)
	fmt.Printf("nilPtr == nil: %t\n", nilPtr == nil)

	// ❌ nil 포인터 역참조 → panic!
	// fmt.Println(nilPtr.Name)  // panic: nil pointer dereference

	// ✅ nil 체크 후 사용
	if nilPtr != nil {
		fmt.Println(nilPtr.Name)
	} else {
		fmt.Println("nilPtr가 nil이라 접근 불가 (안전하게 처리)")
	}
	fmt.Println()

	// ----------------------------------------
	// 7. 언제 포인터를 쓰고 언제 값을 쓰나?
	// ----------------------------------------
	fmt.Println("7. 포인터 vs 값 사용 기준")
	fmt.Println(`
✅ 포인터를 써야 하는 경우:
   1. 함수 안에서 원본을 변경해야 할 때
      → func (u *User) ChangeName(name string)

   2. struct가 클 때 (복사 비용 아끼려고)
      → func process(data *BigStruct)

   3. nil 가능성을 표현해야 할 때
      → func findUser(id int) (*User, error)
         없으면 nil, 있으면 *User 반환

✅ 값을 써야 하는 경우:
   1. 원본을 변경할 필요가 없을 때 (읽기 전용)
      → func (u User) GetName() string

   2. int, string, bool 같은 기본 타입
      → 이미 충분히 작아서 복사해도 됨

   3. 불변(immutable)하게 유지하고 싶을 때
      → 값을 넘기면 함수가 원본을 못 바꿈이 보장됨
	`)

	// ----------------------------------------
	// Java vs Go 비교
	// ----------------------------------------
	fmt.Println("=== Java vs Go 비교 ===")
	fmt.Println(`
Java 타입별 원본 변경 여부:
  기본 타입 (int, double, boolean 등) → 값 복사 → 원본 안 바뀜
  객체 타입 (User, List 등)           → 참조 전달 → 원본 바뀜
  → Java는 타입에 따라 자동으로 결정됨 (개발자가 선택 불가)

Go 타입별 원본 변경 여부:
  모든 타입 (int, struct 등) 기본 → 값 복사 → 원본 안 바뀜
  & 붙이면                        → 포인터 전달 → 원본 바뀜
  → Go는 타입 상관없이 개발자가 명시적으로 선택

Java:                               Go:
// ❌ 기본 타입: 원본 안 바뀜        // ❌ & 없음: 원본 안 바뀜
int x = 10;                         x := 10
void addTen(int n) {                func addTen(n int) {
  n += 10; // 복사본만 바뀜             n += 10 // 복사본만 바뀜
}                                   }

// ✅ 객체 타입: 원본 바뀜           // ✅ & 붙임: 원본 바뀜
void changeName(User u) {           func addTen(n *int) {
  u.setName("변경"); // 원본 바뀜       *n += 10 // 원본 바뀜
}                                   }
// Java는 객체면 자동으로 참조 전달  addTen(&x) // 명시적으로 주소 전달
// 개발자가 선택 불가               // Go는 항상 개발자가 선택
	`)

	// ----------------------------------------
	// 실습 문제!
	// ----------------------------------------
	fmt.Println("\n=== 실습 문제 ===")

	// TODO 1: swap 함수를 만드세요
	//   두 int 변수의 값을 서로 바꾸는 함수
	//   포인터를 사용해서 원본이 바뀌어야 함
	//
	//   a, b := 10, 20
	//   swap(&a, &b)
	//   → a=20, b=10

	// TODO 2: Counter struct를 만드세요
	//   필드: count int
	//   메서드:
	//     Increment() → count를 1 증가 (포인터 리시버)
	//     Reset()     → count를 0으로 (포인터 리시버)
	//     Value() int → 현재 count 반환 (값 리시버)
	//   포인터 리시버와 값 리시버를 올바르게 구분해서 사용할 것

	// TODO 3: findEven 함수를 만드세요
	//   []int 슬라이스를 받아서
	//   짝수가 있으면 *int (첫번째 짝수의 포인터)를 반환
	//   짝수가 없으면 nil 반환
	//   반환값을 nil 체크 후 안전하게 출력

	fmt.Println("위의 TODO를 직접 구현해보세요!")
	fmt.Println("\n학습 완료! Day 2 완료! 다음은 Day 3 - 09_goroutine.go 로 가봐요.")
}

// ============================================
// 함수 정의들
// ============================================

type Person struct {
	Name string
	Age  int
}

// 값 전달: 복사본이 넘어옴 → 원본 변경 불가
func addTenByValue(n int) {
	n += 10
	fmt.Printf("  (함수 안) n = %d  ← 복사본만 바뀜\n", n)
}

// 포인터 전달: 주소가 넘어옴 → 원본 직접 변경
func addTenByPointer(n *int) {
	*n += 10 // *n = 주소로 찾아가서 값을 변경
	fmt.Printf("  (함수 안) *n = %d  ← 원본을 바꿈\n", *n)
}

// Struct 값 전달
func growUpByValue(p Person) {
	p.Age++
	fmt.Printf("  (함수 안) p.Age = %d  ← 복사본만 바뀜\n", p.Age)
}

// Struct 포인터 전달
func growUpByPointer(p *Person) {
	p.Age++
	fmt.Printf("  (함수 안) p.Age = %d  ← 원본을 바꿈\n", p.Age)
}

// 값 리시버: 복사본에서 실행
func (p Person) BirthdayByValue() {
	p.Age++
	fmt.Printf("  (값 리시버 안) p.Age = %d  ← 복사본만 바뀜\n", p.Age)
}

// 포인터 리시버: 원본에서 실행
func (p *Person) BirthdayByPointer() {
	p.Age++
	fmt.Printf("  (포인터 리시버 안) p.Age = %d  ← 원본을 바꿈\n", p.Age)
}

// 복사 비용 비교용 큰 struct
type BigStruct struct {
	Data [100]int // int 100개 = 800바이트
}

func processByValue(b BigStruct) {
	fmt.Printf("  받은 데이터 첫번째 값: %d\n", b.Data[0])
}

func processByPointer(b *BigStruct) {
	fmt.Printf("  받은 데이터 첫번째 값: %d\n", b.Data[0])
}
