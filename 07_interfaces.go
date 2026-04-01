package main

import (
	"fmt"
	"math"
)

// ============================================
// Day 2-3: 인터페이스 (Go의 진짜 핵심 ⭐⭐⭐)
// ============================================
//
// Java와 가장 큰 차이:
//   Java: implements 키워드로 명시적 선언 필수
//   Go:   메서드만 있으면 자동으로 인터페이스 만족 (Duck Typing)
//
// 📌 핵심:
//   "이 타입이 인터페이스를 구현한다"고 어디에도 선언 안 함
//   → 메서드 시그니처가 맞으면 그냥 자동으로 만족

// ============================================
// 1. 기본 인터페이스
// ============================================

// Java: public interface Shape { double area(); }
// Go:   type Shape interface { Area() float64 }
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle - Shape 인터페이스를 구현 (implements 선언 없음!)
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Rectangle - 마찬가지로 Shape 인터페이스를 구현
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Triangle
type Triangle struct {
	A, B, C float64 // 세 변의 길이
}

func (t Triangle) Area() float64 {
	// 헤론의 공식
	s := (t.A + t.B + t.C) / 2
	return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}

func (t Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

// Shape 인터페이스를 인자로 받는 함수
// Circle이든 Rectangle이든 Triangle이든 Shape이면 다 받을 수 있음
func printShapeInfo(s Shape) {
	fmt.Printf("넓이: %.2f, 둘레: %.2f\n", s.Area(), s.Perimeter())
}

// ============================================
// 2. 인터페이스를 작게 만들기 (Go 철학)
// ============================================

// ❌ Java스러운 방식: 크고 뚱뚱한 인터페이스
type UserServiceBad interface {
	FindByID(id int) (*UserI, error)
	Create(user *UserI) error
	Update(user *UserI) error
	Delete(id int) error
	List() ([]*UserI, error)
	// 너무 많음 → 테스트하기 어렵고 의존성 강함
}

// ✅ Go다운 방식: 소비자가 필요한 것만 담은 작은 인터페이스
type UserFinder interface {
	FindByID(id int) (*UserI, error)
}

type UserCreator interface {
	Create(user *UserI) error
}

// 필요하면 조합도 가능
type UserRepository interface {
	UserFinder
	UserCreator
}

type UserI struct {
	ID   int
	Name string
}

// ============================================
// 3. 인터페이스로 의존성 주입 (실무 패턴)
// ============================================

// Notifier 인터페이스: 알림 전송 방식을 추상화
type Notifier interface {
	Send(message string) error
}

// EmailNotifier - Notifier 구현
type EmailNotifier struct {
	Address string
}

func (e EmailNotifier) Send(message string) error {
	fmt.Printf("[이메일 → %s] %s\n", e.Address, message)
	return nil
}

// SlackNotifier - Notifier 구현
type SlackNotifier struct {
	Channel string
}

func (s SlackNotifier) Send(message string) error {
	fmt.Printf("[슬랙 → #%s] %s\n", s.Channel, message)
	return nil
}

// SMSNotifier - Notifier 구현
type SMSNotifier struct {
	PhoneNumber string
}

func (s SMSNotifier) Send(message string) error {
	fmt.Printf("[SMS → %s] %s\n", s.PhoneNumber, message)
	return nil
}

// OrderService: Notifier 인터페이스에만 의존
// → EmailNotifier든 SlackNotifier든 뭐가 들어오든 상관없음
type OrderService struct {
	notifier Notifier // 인터페이스 타입으로 저장
}

func NewOrderService(notifier Notifier) *OrderService {
	return &OrderService{notifier: notifier}
}

func (s *OrderService) PlaceOrder(productName string) error {
	// 주문 처리 로직...
	message := fmt.Sprintf("주문 완료: %s", productName)
	return s.notifier.Send(message) // 어떤 Notifier인지 몰라도 됨
}

// ============================================
// 4. 빈 인터페이스 (interface{} / any)
// ============================================

// interface{}: 모든 타입을 받을 수 있음
// Java의 Object와 비슷
// Go 1.18부터 any 라는 별칭 사용 가능

func printAnything(v any) {
	fmt.Printf("값: %v, 타입: %T\n", v, v)
}

// ============================================
// 5. 타입 어서션 & 타입 스위치
// ============================================

// 타입 어서션: 인터페이스에서 실제 타입 꺼내기
// Java의 instanceof + 캐스팅과 비슷

func describeShape(s Shape) {
	// 타입 스위치: 실제 타입에 따라 다른 처리
	switch v := s.(type) {
	case Circle:
		fmt.Printf("원 (반지름: %.1f)\n", v.Radius)
	case Rectangle:
		fmt.Printf("직사각형 (가로: %.1f, 세로: %.1f)\n", v.Width, v.Height)
	case Triangle:
		fmt.Printf("삼각형 (변: %.1f, %.1f, %.1f)\n", v.A, v.B, v.C)
	default:
		fmt.Printf("알 수 없는 도형: %T\n", v)
	}
}

// ============================================
// 메인 함수
// ============================================

func main() {
	fmt.Println("=== Go 인터페이스 학습 ===\n")

	// ----------------------------------------
	// 1. 기본 인터페이스 사용
	// ----------------------------------------
	fmt.Println("1. 기본 인터페이스")

	circle := Circle{Radius: 5}
	rect := Rectangle{Width: 4, Height: 6}
	tri := Triangle{A: 3, B: 4, C: 5}

	// 각 타입의 메서드 직접 호출
	fmt.Printf("원 넓이: %.2f\n", circle.Area())
	fmt.Printf("직사각형 넓이: %.2f\n", rect.Area())

	// Shape 인터페이스로 통일해서 처리
	fmt.Println("\nShape 인터페이스로 통일:")
	shapes := []Shape{circle, rect, tri} // 서로 다른 타입이지만 Shape으로 묶임
	for _, s := range shapes {
		printShapeInfo(s)
	}
	fmt.Println()

	// ----------------------------------------
	// 2. 인터페이스로 의존성 주입
	// ----------------------------------------
	fmt.Println("2. 인터페이스로 의존성 주입")

	// 이메일로 알림
	emailService := NewOrderService(EmailNotifier{Address: "user@example.com"})
	emailService.PlaceOrder("맥북 프로")

	// 슬랙으로 알림 (OrderService 코드 변경 없이 교체!)
	slackService := NewOrderService(SlackNotifier{Channel: "orders"})
	slackService.PlaceOrder("아이패드")

	// SMS로 알림
	smsService := NewOrderService(SMSNotifier{PhoneNumber: "010-1234-5678"})
	smsService.PlaceOrder("애플워치")
	fmt.Println()

	// ----------------------------------------
	// 3. 빈 인터페이스
	// ----------------------------------------
	fmt.Println("3. 빈 인터페이스 (any)")

	printAnything(42)
	printAnything("hello")
	printAnything(true)
	printAnything(circle)
	printAnything([]int{1, 2, 3})
	fmt.Println()

	// ----------------------------------------
	// 4. 타입 어서션
	// ----------------------------------------
	fmt.Println("4. 타입 어서션")

	var s Shape = Circle{Radius: 3}

	// 안전한 타입 어서션 (ok 패턴)
	c, ok := s.(Circle)
	if ok {
		fmt.Printf("Circle 맞음, 반지름: %.1f\n", c.Radius)
	}

	// 잘못된 타입 어서션 시도
	_, ok2 := s.(Rectangle)
	fmt.Printf("Rectangle 맞음: %t\n\n", ok2) // false

	// ----------------------------------------
	// 5. 타입 스위치
	// ----------------------------------------
	fmt.Println("5. 타입 스위치")

	shapes2 := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 3, Height: 4},
		Triangle{A: 3, B: 4, C: 5},
	}

	for _, s := range shapes2 {
		describeShape(s)
	}
	fmt.Println()

	// ----------------------------------------
	// 6. 인터페이스 nil 주의사항
	// ----------------------------------------
	fmt.Println("6. 인터페이스 nil 주의사항")

	var n Notifier = nil                    // 인터페이스 자체가 nil
	fmt.Printf("nil 인터페이스: %v\n", n == nil) // true

	// 📌 주의: 인터페이스에 nil 포인터를 담으면 nil이 아님!
	var email *EmailNotifier = nil                           // 포인터가 nil
	var notifier Notifier = email                            // 인터페이스에 nil 포인터 담기
	fmt.Printf("nil 포인터를 담은 인터페이스: %v\n\n", notifier == nil) // false! 함정!

	// ----------------------------------------
	// Java vs Go 비교
	// ----------------------------------------
	fmt.Println("=== Java vs Go 비교 ===")
	fmt.Println(`
Java:                               Go:
interface Shape {                   type Shape interface {
  double area();                        Area() float64
}                                   }

class Circle implements Shape {     type Circle struct { Radius float64 }
  public double area() { ... }      func (c Circle) Area() float64 { ... }
}                                   // implements 선언 없음!

Shape s = new Circle();             var s Shape = Circle{Radius: 5}

if (s instanceof Circle c) {        if c, ok := s.(Circle); ok {
  c.getRadius()                         c.Radius
}                                   }

switch (s) {                        switch v := s.(type) {
  case Circle c: ...                case Circle: ...
  case Rectangle r: ...             case Rectangle: ...
}                                   }

Object obj = anything;              var obj any = anything
	`)

	// ----------------------------------------
	// 실습 문제!
	// ----------------------------------------
	fmt.Println("\n=== 실습 문제 ===")

	// TODO 1: Animal 인터페이스를 만드세요
	//   메서드: Sound() string, Name() string
	//   구현체: Dog, Cat, Cow
	//   함수: makeSound(a Animal) → "{이름}: {소리}" 출력

	// TODO 2: Logger 인터페이스를 만드세요
	//   메서드: Log(level, message string)
	//   구현체: ConsoleLogger (콘솔 출력), FileLogger (파일명 저장)
	//   Service struct에 Logger를 주입해서 사용

	fmt.Println("위의 TODO를 직접 구현해보세요!")
	fmt.Println("\n학습 완료! 다음은 08_pointers.go 로 가봐요.")
}
