package main

import (
	"errors"
	"fmt"
)

// ============================================
// Day 2-2: 에러 처리 (Go의 철학 ⭐⭐⭐)
// ============================================
//
// Java와 가장 큰 사고방식 차이!
//
//   Java: 에러 = 예외(Exception) → try/catch로 분리
//   Go:   에러 = 값(Value)       → 반환값으로 처리
//
// 📌 핵심 마인드셋:
//   "에러는 특별한 게 아니다, 그냥 반환값 중 하나다"
//   → 숨기거나 무시하면 안 됨, 반드시 명시적으로 처리

// ============================================
// 1. 기본 에러 생성 방법
// ============================================

func basicErrorExamples() {
	fmt.Println("1. 에러 생성 방법")

	// 방법 1: errors.New() - 단순 에러 메시지
	err1 := errors.New("something went wrong")
	fmt.Printf("errors.New: %v\n", err1)

	// 방법 2: fmt.Errorf() - 포맷 지정 가능
	id := 42
	err2 := fmt.Errorf("ID %d인 사용자를 찾을 수 없습니다", id)
	fmt.Printf("fmt.Errorf: %v\n\n", err2)
}

// ============================================
// 2. 에러 반환 패턴 (실무 기본 패턴)
// ============================================

type User struct {
	ID   int
	Name string
}

// Java: User findById(int id) throws UserNotFoundException
// Go:   func findUser(id int) (*User, error)
func findUser(id int) (*User, error) {
	if id <= 0 {
		return nil, fmt.Errorf("유효하지 않은 ID: %d", id)
	}
	if id == 999 {
		return nil, fmt.Errorf("ID %d인 사용자 없음", id)
	}
	return &User{ID: id, Name: "홍길동"}, nil
}

// ============================================
// 3. Sentinel Error (특정 에러 구분)
// ============================================

// 패키지 레벨에서 미리 정의해두는 에러 값
// Java의 특정 Exception 클래스와 비슷한 역할
var (
	ErrNotFound    = errors.New("not found")
	ErrUnauthorized = errors.New("unauthorized")
	ErrInvalidInput = errors.New("invalid input")
)

func getUser(id int) (*User, error) {
	if id <= 0 {
		return nil, ErrInvalidInput
	}
	if id == 999 {
		return nil, ErrNotFound
	}
	if id == 403 {
		return nil, ErrUnauthorized
	}
	return &User{ID: id, Name: "홍길동"}, nil
}

// ============================================
// 4. 에러 래핑 (Error Wrapping) ⭐
// ============================================

// %w 로 에러를 감싸면 원인 추적 가능
// Java의 new RuntimeException("msg", cause) 와 비슷

func validateUser(id int) error {
	if id <= 0 {
		return fmt.Errorf("validateUser 실패: %w", ErrInvalidInput)
	}
	return nil
}

func createOrder(userID int, amount float64) error {
	if err := validateUser(userID); err != nil {
		return fmt.Errorf("createOrder 실패: %w", err) // 에러 체인
	}
	if amount <= 0 {
		return fmt.Errorf("createOrder 실패: 금액은 0보다 커야 합니다")
	}
	return nil
}

// ============================================
// 5. 커스텀 에러 타입
// ============================================

// Java: class ValidationException extends RuntimeException { ... }
// Go:   error 인터페이스를 구현하는 struct

type ValidationError struct {
	Field   string
	Message string
}

// error 인터페이스 구현 (Error() string 메서드만 있으면 됨)
func (e *ValidationError) Error() string {
	return fmt.Sprintf("유효성 검사 실패 - %s: %s", e.Field, e.Message)
}

type DatabaseError struct {
	Code    int
	Message string
}

func (e *DatabaseError) Error() string {
	return fmt.Sprintf("DB 에러 [%d]: %s", e.Code, e.Message)
}

func saveUser(user *User) error {
	if user.Name == "" {
		return &ValidationError{Field: "Name", Message: "이름은 필수입니다"}
	}
	if user.ID == 0 {
		return &DatabaseError{Code: 1001, Message: "ID가 없으면 저장 불가"}
	}
	return nil
}

// ============================================
// 메인 함수
// ============================================

func main() {
	fmt.Println("=== Go 에러 처리 학습 ===\n")

	// ----------------------------------------
	// 1. 기본 에러 생성
	// ----------------------------------------
	basicErrorExamples()

	// ----------------------------------------
	// 2. 기본 에러 처리 패턴
	// ----------------------------------------
	fmt.Println("2. 기본 에러 처리 패턴")

	// ✅ Go의 기본 패턴: 에러를 바로 체크
	user, err := findUser(1)
	if err != nil {
		fmt.Printf("에러: %v\n", err)
	} else {
		fmt.Printf("찾은 사용자: %+v\n", *user)
	}

	// 에러 케이스
	_, err = findUser(-1)
	if err != nil {
		fmt.Printf("에러: %v\n", err)
	}

	// 에러만 필요할 때 _ 로 값 무시
	_, err = findUser(999)
	if err != nil {
		fmt.Printf("에러: %v\n\n", err)
	}

	// ----------------------------------------
	// 3. Sentinel Error + errors.Is()
	// ----------------------------------------
	fmt.Println("3. Sentinel Error 구분 (errors.Is)")

	// errors.Is(): 에러가 특정 sentinel error인지 확인
	// Java: e instanceof UserNotFoundException
	ids := []int{1, -1, 999, 403}
	for _, id := range ids {
		user, err := getUser(id)
		if err != nil {
			switch {
			case errors.Is(err, ErrNotFound):
				fmt.Printf("ID %d → 404 Not Found\n", id)
			case errors.Is(err, ErrUnauthorized):
				fmt.Printf("ID %d → 401 Unauthorized\n", id)
			case errors.Is(err, ErrInvalidInput):
				fmt.Printf("ID %d → 400 Bad Request\n", id)
			default:
				fmt.Printf("ID %d → 500 Internal Error: %v\n", id, err)
			}
		} else {
			fmt.Printf("ID %d → 200 OK: %s\n", id, user.Name)
		}
	}
	fmt.Println()

	// ----------------------------------------
	// 4. 에러 래핑 + errors.Is() / errors.As()
	// ----------------------------------------
	fmt.Println("4. 에러 래핑 (체인 추적)")

	err = createOrder(-1, 1000)
	if err != nil {
		fmt.Printf("에러 메시지: %v\n", err)

		// errors.Is(): 래핑된 에러 체인 안에 ErrInvalidInput이 있는지 확인
		if errors.Is(err, ErrInvalidInput) {
			fmt.Println("→ 근본 원인: 잘못된 입력값")
		}
	}
	fmt.Println()

	// ----------------------------------------
	// 5. 커스텀 에러 타입 + errors.As()
	// ----------------------------------------
	fmt.Println("5. 커스텀 에러 타입 구분 (errors.As)")

	// errors.As(): 에러가 특정 타입인지 확인 + 해당 타입으로 변환
	// Java: catch (ValidationException e) { e.getField() }
	users := []*User{
		{ID: 1, Name: ""},  // ValidationError 발생
		{ID: 0, Name: "김철수"}, // DatabaseError 발생
		{ID: 2, Name: "이영희"}, // 정상
	}

	for _, u := range users {
		err := saveUser(u)
		if err != nil {
			var validErr *ValidationError
			var dbErr *DatabaseError

			switch {
			case errors.As(err, &validErr):
				fmt.Printf("유효성 에러 → 필드: %s, 메시지: %s\n", validErr.Field, validErr.Message)
			case errors.As(err, &dbErr):
				fmt.Printf("DB 에러 → 코드: %d, 메시지: %s\n", dbErr.Code, dbErr.Message)
			default:
				fmt.Printf("알 수 없는 에러: %v\n", err)
			}
		} else {
			fmt.Printf("저장 성공: %s\n", u.Name)
		}
	}
	fmt.Println()

	// ----------------------------------------
	// 6. 에러 처리 안티패턴 (하면 안 되는 것)
	// ----------------------------------------
	fmt.Println("6. 에러 처리 안티패턴")

	fmt.Println(`
❌ 에러 무시 (절대 하면 안 됨!)
   result, _ := findUser(id)   // _ 로 에러 무시
   → 에러가 발생해도 모름, 나중에 nil pointer panic 위험

❌ 에러 로그만 찍고 계속 진행
   if err != nil { log.Println(err) }  // 처리 안 하고 계속
   result.DoSomething()                // nil일 수 있음 → panic!

✅ 올바른 패턴
   result, err := findUser(id)
   if err != nil {
       return fmt.Errorf("처리 실패: %w", err)  // 상위로 전달
   }
   result.DoSomething()  // 여기서 result는 nil이 아님이 보장됨
	`)

	// ----------------------------------------
	// Java vs Go 비교
	// ----------------------------------------
	fmt.Println("=== Java vs Go 비교 ===")
	fmt.Println(`
Java:                               Go:
try {                               user, err := findUser(id)
  User user = findUser(id);         if err != nil {
} catch (NotFoundException e) {         // 에러 처리
  // 처리                           }
}

throw new NotFoundException(...)    return nil, ErrNotFound

catch (NotFoundException e) {}      errors.Is(err, ErrNotFound)

catch (ValidationException e) {     var ve *ValidationError
  e.getField()                      if errors.As(err, &ve) {
}                                       ve.Field
                                    }

new RuntimeException("msg", cause)  fmt.Errorf("msg: %w", cause)
	`)

	// ----------------------------------------
	// 실습 문제!
	// ----------------------------------------
	fmt.Println("\n=== 실습 문제 ===")

	// TODO 1: BankAccount struct를 만드세요
	//   필드: balance float64
	//   메서드: Deposit(amount float64) error
	//           Withdraw(amount float64) error
	//   조건: 입금/출금 금액이 0 이하면 ErrInvalidInput 반환
	//         잔액 부족이면 커스텀 InsufficientFundsError 반환

	// TODO 2: 아래 시나리오를 에러 처리와 함께 구현하세요
	//   1. 계좌 생성 (잔액 1000)
	//   2. 500 출금 → 성공
	//   3. 800 출금 → InsufficientFundsError 발생, 잔액 출력
	//   4. -100 입금 → ErrInvalidInput 발생

	fmt.Println("위의 TODO를 직접 구현해보세요!")
	fmt.Println("\n학습 완료! 다음은 07_interfaces.go 로 가봐요.")
}
