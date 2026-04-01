package main

import "fmt"

// ============================================
// 함수 vs 메서드 vs 리시버
// ============================================
//
// 핵심 정리:
//   함수   → 리시버 없음, 독립적으로 존재
//   메서드 → 리시버 있음, struct에 종속
//   리시버 → 메서드를 struct에 연결해주는 부분
//
// 구분 기준: 리시버가 있냐 없냐
//   리시버 없음 → 함수
//   리시버 있음 → 메서드
//
// ┌─────────────────────────────────────────────────────┐
// │           함수 vs 메서드 핵심 차이                   │
// ├──────────────┬──────────────────┬───────────────────┤
// │              │ 함수             │ 메서드            │
// ├──────────────┼──────────────────┼───────────────────┤
// │ 리시버       │ 없음             │ 있음              │
// │ 소속         │ 없음 (독립적)    │ struct에 종속     │
// │ 호출 방법    │ 함수명(인자)     │ 변수.메서드명()   │
// │ Java 대응    │ static 메서드    │ 인스턴스 메서드   │
// └──────────────┴──────────────────┴───────────────────┘

// ============================================
// 사용할 Struct
// ============================================

type Account struct {
	Owner   string
	Balance int
}

// ============================================
// 1. 함수 (Function)
// ============================================
//
// 리시버 없음 → 독립적으로 존재
// 호출: 함수명(인자)

func createAccount(owner string, balance int) Account {
	return Account{Owner: owner, Balance: balance}
}

func printAccount(a Account) {
	fmt.Printf("계좌주: %s, 잔액: %d원\n", a.Owner, a.Balance)
}

// ============================================
// 2. 메서드 (Method)
// ============================================
//
// 리시버 있음 → Account에 종속
// 호출: 변수.메서드명()
//
// 메서드 = 리시버 + 함수 본체
// func (리시버) 메서드명(파라미터) 반환타입 { }
//       ↑ 이 부분이 리시버

// 값 리시버: 읽기 전용, 원본 변경 안 됨
func (a Account) GetInfo() string {
	// a는 Account의 복사본
	return fmt.Sprintf("계좌주: %s, 잔액: %d원", a.Owner, a.Balance)
}

// 포인터 리시버: 원본 변경 가능
func (a *Account) Deposit(amount int) {
	// a는 Account의 포인터 → 원본 직접 변경
	a.Balance += amount
}

func (a *Account) Withdraw(amount int) error {
	if amount > a.Balance {
		return fmt.Errorf("잔액 부족: 잔액=%d, 출금요청=%d", a.Balance, amount)
	}
	a.Balance -= amount
	return nil
}

// ============================================
// 리시버 설명
// ============================================
//
// func (a Account) GetInfo() string { }
//       ↑
//       리시버: "이 메서드는 Account 소속이다" 라고 선언하는 부분
//              Java의 this + 클래스 소속 선언을 한번에 하는 것
//
// 값 리시버  (a Account)  → 복사본, 원본 변경 불가 (읽기 전용)
// 포인터 리시버 (a *Account) → 원본, 원본 변경 가능

// ============================================
// 메인 함수
// ============================================

func main() {
	fmt.Println("=== 함수 vs 메서드 vs 리시버 ===\n")

	// ----------------------------------------
	// 1. 함수 호출
	// ----------------------------------------
	fmt.Println("1. 함수 호출 (리시버 없음)")

	// 함수: 그냥 함수명(인자) 으로 호출
	account := createAccount("홍길동", 10000)
	printAccount(account)
	fmt.Println()

	// ----------------------------------------
	// 2. 메서드 호출
	// ----------------------------------------
	fmt.Println("2. 메서드 호출 (리시버 있음)")

	// 메서드: 변수.메서드명() 으로 호출
	fmt.Println(account.GetInfo()) // 값 리시버 → account가 a(복사본)로 들어감

	account.Deposit(5000)          // 포인터 리시버 → account가 a(원본)로 들어감
	fmt.Println(account.GetInfo()) // 잔액이 15000으로 바뀜

	err := account.Withdraw(3000)
	if err != nil {
		fmt.Println("에러:", err)
	} else {
		fmt.Println(account.GetInfo()) // 잔액이 12000으로 바뀜
	}

	// 잔액 부족 케이스
	err = account.Withdraw(99999)
	if err != nil {
		fmt.Println("에러:", err)
	}
	fmt.Println()

	// ----------------------------------------
	// 3. 함수 vs 메서드 차이 비교
	// ----------------------------------------
	fmt.Println("3. 함수 vs 메서드 차이")

	a := Account{Owner: "김철수", Balance: 5000}

	// 함수: 인자로 넘김
	printAccount(a)         // 함수 호출
	createAccount("이영희", 0) // 함수 호출

	// 메서드: 변수.메서드명()
	a.Deposit(1000) // 메서드 호출
	a.GetInfo()     // 메서드 호출
	fmt.Println(a.GetInfo())
	fmt.Println()

	// ----------------------------------------
	// Java vs Go 비교
	// ----------------------------------------
	fmt.Println("=== Java vs Go 비교 ===")
	fmt.Println(`
Java:                               Go:
// 클래스 안에 있으면 자동으로 소속   // 리시버로 명시적으로 소속 선언
public class Account {
    public String getInfo() {       func (a Account) GetInfo() string {
        return this.owner;              return a.Owner
    }                               }
    public void deposit(int n) {    func (a *Account) Deposit(n int) {
        this.balance += n;              a.Balance += n
    }                               }
}

// static 메서드 (클래스 소속)      // 함수 (리시버 없음)
Account.create(...)                 createAccount(...)

// 인스턴스 메서드                  // 메서드 (리시버 있음)
account.deposit(1000)               account.Deposit(1000)
	`)

	// ----------------------------------------
	// 실습 문제!
	// ----------------------------------------
	fmt.Println("\n=== 실습 문제 ===")

	// TODO 1: Calculator struct를 만드세요
	//   필드: result float64
	//   메서드 (포인터 리시버):
	//     Add(n float64)      → result에 n을 더함
	//     Subtract(n float64) → result에서 n을 뺌
	//     Multiply(n float64) → result에 n을 곱함
	//   메서드 (값 리시버):
	//     GetResult() float64 → result 반환
	//
	//   함수:
	//     NewCalculator(initial float64) Calculator → 초기값으로 생성

	// TODO 2: 아래 출력이 나오도록 코드를 작성하세요
	//   초기값: 10
	//   +5 후: 15
	//   -3 후: 12
	//   *2 후: 24

	fmt.Println("위의 TODO를 직접 구현해보세요!")
}
