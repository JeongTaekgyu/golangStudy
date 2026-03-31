package main

import "fmt"

// ============================================
// Day 1-3: Struct (Go의 클래스 대체 개념)
// ============================================
//
// Java와 가장 큰 차이:
//   Java: class + 상속(extends) + 캡슐화
//   Go:   struct + 컴포지션(embedding) + 메서드
//
// ❌ Go에는 없는 것: class, extends, implements, super, this
// ✅ Go에 있는 것:   struct, 메서드, 인터페이스, 임베딩

// ============================================
// 1. 기본 Struct 선언
// ============================================

// Java: public class User { ... }
// Go:   type User struct { ... }
type User struct {
	ID    int    // 대문자 = public (패키지 외부에서 접근 가능)
	Name  string // 대문자 = public
	Email string // 대문자 = public
	age   int    // 소문자 = private (이 패키지 내부에서만 접근 가능)
}

// ============================================
// 2. 생성자 패턴 (Go에는 new 키워드 없음)
// ============================================

// Java: new User(id, name, email)
// Go:   관례적으로 New함수명() 함수를 만들어서 사용
func NewUser(id int, name string, email string) *User {
	// 유효성 검사도 여기서!
	return &User{
		ID:    id,
		Name:  name,
		Email: email,
		age:   0,
	}
}

// ============================================
// 3. 메서드 (값 리시버 vs 포인터 리시버)
// ============================================

// 값 리시버 (Value Receiver): 복사본에서 실행 → 원본 변경 안 됨
// Java의 getter 같은 읽기 전용 메서드에 사용
func (u User) Greet() string {
	// 여기서 u를 변경해도 원본 User는 그대로
	return fmt.Sprintf("안녕하세요, 저는 %s입니다!", u.Name)
}

// 포인터 리시버 (Pointer Receiver): 원본을 직접 변경
// Java의 setter 같은 상태 변경 메서드에 사용
// 📌 규칙: 상태 변경이 필요하면 항상 포인터 리시버 사용!
func (u *User) ChangeName(name string) {
	u.Name = name // 원본 User.Name이 실제로 바뀜
}

func (u *User) SetAge(age int) {
	u.age = age
}

func (u User) GetAge() int {
	return u.age
}

// ============================================
// 4. Struct 임베딩 (Go의 컴포지션 = 상속 대체)
// ============================================

// Java의 상속: class Admin extends User { ... }
// Go의 컴포지션: Admin에 User를 embed

type Address struct {
	City    string
	Country string
}

// Address 메서드
func (a Address) FullAddress() string {
	return fmt.Sprintf("%s, %s", a.City, a.Country)
}

type Admin struct {
	User    // 임베딩: User의 필드와 메서드를 모두 가져옴
	Address // 임베딩: Address의 필드와 메서드도 가져옴
	Level   int
}

// ============================================
// 5. 중첩 Struct
// ============================================

type Order struct {
	ID       int
	Customer User // 중첩 (임베딩이 아닌 일반 필드로 포함)
	Amount   float64
}

// ============================================
// 메인 함수
// ============================================

func main() {
	fmt.Println("=== Go Struct 학습 ===\n")

	// ----------------------------------------
	// 1. Struct 생성 방법
	// ----------------------------------------
	fmt.Println("1. Struct 생성 방법")

	// 방법 1: 필드명 지정 (권장 ✅)
	user1 := User{
		ID:    1,
		Name:  "홍길동",
		Email: "hong@example.com",
	}
	fmt.Printf("user1: %+v\n", user1) // %+v 는 필드명까지 출력

	// 방법 2: 생성자 함수 사용 (실무 권장 ✅✅)
	user2 := NewUser(2, "김철수", "kim@example.com")
	fmt.Printf("user2: %+v\n", *user2) // *user2 로 역참조

	// 방법 3: 빈 struct 후 필드 채우기
	var user3 User
	user3.ID = 3
	user3.Name = "이영희"
	fmt.Printf("user3: %+v\n\n", user3)

	// ----------------------------------------
	// 2. 값 리시버 vs 포인터 리시버 차이
	// ----------------------------------------
	fmt.Println("2. 값 리시버 vs 포인터 리시버")

	user := NewUser(1, "원래이름", "test@example.com")
	fmt.Printf("변경 전: %s\n", user.Name)

	// 포인터 리시버: 원본이 변경됨
	user.ChangeName("바뀐이름")
	fmt.Printf("변경 후: %s\n", user.Name)

	// 메서드 호출 (값 리시버)
	fmt.Println(user.Greet())

	// private 필드는 메서드로만 접근
	user.SetAge(28)
	fmt.Printf("나이: %d\n\n", user.GetAge())

	// ----------------------------------------
	// 3. Struct 복사 주의사항
	// ----------------------------------------
	fmt.Println("3. Struct 복사 주의사항")

	original := User{ID: 1, Name: "원본"}
	copied := original // 값 복사 → 완전히 독립된 복사본
	copied.Name = "복사본"

	fmt.Printf("원본: %s\n", original.Name) // "원본" 그대로
	fmt.Printf("복사본: %s\n\n", copied.Name)

	// 포인터로 복사하면? → 같은 주소를 가리킴
	ptr1 := &User{ID: 2, Name: "포인터원본"}
	ptr2 := ptr1 // 주소 복사 → 같은 User를 가리킴!
	ptr2.Name = "포인터변경"

	fmt.Printf("ptr1: %s\n", ptr1.Name) // "포인터변경"으로 바뀜!
	fmt.Printf("ptr2: %s\n\n", ptr2.Name)

	// ----------------------------------------
	// 4. 임베딩 (컴포지션)
	// ----------------------------------------
	fmt.Println("4. 임베딩 (컴포지션)")

	admin := Admin{
		User: User{
			ID:    10,
			Name:  "관리자",
			Email: "admin@example.com",
		},
		Address: Address{
			City:    "서울",
			Country: "한국",
		},
		Level: 5,
	}

	// 임베딩된 필드에 직접 접근 가능 (Java의 상속처럼)
	fmt.Printf("Admin 이름: %s\n", admin.Name)          // admin.User.Name 과 동일
	fmt.Printf("Admin 이메일: %s\n", admin.Email)        // admin.User.Email 과 동일
	fmt.Printf("Admin 도시: %s\n", admin.City)           // admin.Address.City 와 동일
	fmt.Printf("Admin 레벨: %d\n", admin.Level)

	// 임베딩된 메서드도 직접 호출 가능!
	fmt.Println(admin.Greet())            // User의 메서드
	fmt.Println(admin.FullAddress())      // Address의 메서드
	fmt.Println()

	// ----------------------------------------
	// 5. 중첩 Struct
	// ----------------------------------------
	fmt.Println("5. 중첩 Struct")

	order := Order{
		ID: 1001,
		Customer: User{
			ID:   1,
			Name: "구매자",
		},
		Amount: 99000,
	}

	// 중첩 필드 접근은 명시적으로
	fmt.Printf("주문번호: %d, 구매자: %s, 금액: %.0f원\n\n",
		order.ID, order.Customer.Name, order.Amount)

	// ----------------------------------------
	// 6. Struct 비교
	// ----------------------------------------
	fmt.Println("6. Struct 비교")

	a := User{ID: 1, Name: "테스트"}
	b := User{ID: 1, Name: "테스트"}
	c := User{ID: 2, Name: "다른사람"}

	fmt.Printf("a == b: %t\n", a == b) // true (모든 필드가 같으면 equal)
	fmt.Printf("a == c: %t\n\n", a == c) // false

	// ----------------------------------------
	// Java vs Go 비교 정리
	// ----------------------------------------
	fmt.Println("=== Java vs Go 비교 ===")
	fmt.Println(`
Java:                               Go:
class User {                        type User struct {
    private int id;                     ID    int
    public String name;                 Name  string
}                                   }

new User(1, "홍길동")                NewUser(1, "홍길동")

class Admin extends User {          type Admin struct {
    int level;                          User       // 임베딩 (컴포지션)
}                                       Level int
                                    }

user.setName("변경")                 user.ChangeName("변경")  // 포인터 리시버
user.getName()                      user.Name                // 직접 접근 (public이면)
	`)

	// ----------------------------------------
	// 실습 문제!
	// ----------------------------------------
	fmt.Println("\n=== 실습 문제 ===")

	// TODO 1: Product struct를 만드세요
	//   필드: ID(int), Name(string), Price(float64), stock(int, private)
	//   메서드: NewProduct 생성자, GetStock(), Restock(amount int)

	// TODO 2: Cart struct를 만드세요
	//   필드: Items([]Product), TotalPrice(float64)
	//   메서드: AddItem(product Product), PrintCart()

	// TODO 3: 아래 출력이 나오도록 코드를 작성하세요
	//   상품명: 맥북, 가격: 2000000원, 재고: 5개
	//   상품명: 아이패드, 가격: 800000원, 재고: 10개
	//   장바구니 총액: 2800000원

	fmt.Println("위의 TODO를 직접 구현해보세요!")
	fmt.Println("\n학습 완료! 다음은 04_slice_map.go 를 학습해봐요.")
}
