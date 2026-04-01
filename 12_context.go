package main

import (
	"context"
	"fmt"
	"time"
)

// 패키지 레벨에서 선언해야 함수 간에 같은 타입으로 공유 가능
type contextKey string

// ============================================
// Day 3-3: Context
// ============================================
//
// Context란?
//   작업의 취소, 타임아웃, 데이터 전달을 위한 개념
//   goroutine 간에 "이 작업을 언제까지 해라, 취소됐으면 멈춰라"
//   를 전달하는 수단
//
// 실무에서 주로 쓰는 경우:
//   ✅ HTTP 요청 처리 (요청이 끊기면 하위 작업도 취소)
//   ✅ DB 쿼리 타임아웃
//   ✅ 외부 API 호출 타임아웃
//   ✅ goroutine 취소 신호 전달
//
// ┌─────────────────────────────────────────────────────┐
// │              Context 종류                            │
// ├──────────────┬──────────────────────────────────────┤
// │ 종류         │ 설명                                  │
// ├──────────────┼──────────────────────────────────────┤
// │ Background() │ 최상위 context, 취소 없음             │
// │ WithCancel   │ 수동으로 취소 가능                    │
// │ WithTimeout  │ 일정 시간 후 자동 취소                │
// │ WithDeadline │ 특정 시각에 자동 취소                 │
// │ WithValue    │ context에 값 저장                     │
// └──────────────┴──────────────────────────────────────┘

func main() {
	fmt.Println("=== Go Context 학습 ===\n")

	// ----------------------------------------
	// 1. context.Background()
	// ----------------------------------------
	fmt.Println("1. context.Background()")
	// 모든 context의 시작점
	// 취소도 없고 타임아웃도 없는 기본 context
	// main 함수나 최상위에서 사용

	ctx := context.Background()
	fmt.Printf("Background context: %v\n\n", ctx)

	// ----------------------------------------
	// 2. WithCancel (수동 취소)
	// ----------------------------------------
	fmt.Println("2. WithCancel - 수동으로 취소")

	// ctx: 취소 가능한 context
	// cancel: 취소 함수 (호출하면 ctx가 취소됨)
	ctx2, cancel := context.WithCancel(context.Background())
	defer cancel() // 📌 항상 defer cancel() 해줘야 메모리 누수 방지

	go func() {
		select {
		// cancel()이 호출되어야 ctx2.Done() 채널이 닫힘
		// cancel() 미호출 시 이 case는 영원히 대기 상태 (goroutine 누수!)
		case <-ctx2.Done():
			fmt.Printf("goroutine 취소됨: %v\n", ctx2.Err())
		}
	}()

	time.Sleep(50 * time.Millisecond)
	cancel() // 수동으로 취소 → ctx2.Done() 채널 닫힘 → goroutine 깨어남
	time.Sleep(50 * time.Millisecond)
	fmt.Println()

	// ----------------------------------------
	// 3. WithTimeout (타임아웃)
	// ----------------------------------------
	fmt.Println("3. WithTimeout - 일정 시간 후 자동 취소")
	// 가장 많이 쓰는 패턴
	// DB 쿼리, 외부 API 호출 등에 타임아웃 설정

	ctx3, cancel3 := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel3()

	// 타임아웃 안에 완료되는 경우
	fmt.Println("[타임아웃 전 완료]")
	err := doWithTimeout(ctx3, 50*time.Millisecond) // 50ms 작업 → 타임아웃 전 완료
	if err != nil {
		fmt.Printf("에러: %v\n", err)
	} else {
		fmt.Println("작업 성공!")
	}

	// 타임아웃 초과하는 경우
	fmt.Println("\n[타임아웃 초과]")
	ctx4, cancel4 := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel4()

	err = doWithTimeout(ctx4, 200*time.Millisecond) // 200ms 작업 → 타임아웃 초과
	if err != nil {
		fmt.Printf("에러: %v\n", err) // context deadline exceeded
	} else {
		fmt.Println("작업 성공!")
	}
	fmt.Println()

	// ----------------------------------------
	// 4. WithDeadline (특정 시각에 취소)
	// ----------------------------------------
	fmt.Println("4. WithDeadline - 특정 시각에 취소")
	// WithTimeout과 비슷하지만 "몇 초 후" 대신 "몇 시 몇 분에" 취소

	deadline := time.Now().Add(100 * time.Millisecond) // 지금으로부터 100ms 후
	ctx5, cancel5 := context.WithDeadline(context.Background(), deadline)
	defer cancel5()

	select {
	case <-time.After(50 * time.Millisecond):
		fmt.Println("50ms 후 작업 완료 (deadline 전)")
	case <-ctx5.Done():
		fmt.Printf("deadline 초과: %v\n", ctx5.Err())
	}
	fmt.Println()

	// ----------------------------------------
	// 5. WithValue (context에 값 저장)
	// ----------------------------------------
	fmt.Println("5. WithValue - context에 값 저장")
	// HTTP 요청의 사용자 ID, 트레이싱 ID 등을 전달할 때 사용
	//
	// 왜 쓰는가?
	//   userID, requestID 같은 값을 A→B→C 함수 체인에 전달할 때
	//   파라미터로 일일이 넘기면 중간 함수들도 다 파라미터 추가해야 함
	//   ctx 하나만 넘기면 필요한 곳에서 꺼내 쓸 수 있음
	//
	// 뭘 보여주는가?
	//   ctx에 값 담고(WithValue) → 함수에 ctx만 넘기고 → 함수 안에서 꺼내는(ctx.Value) 흐름
	//
	// 📌 주의: 비즈니스 로직 데이터는 파라미터로 전달할 것
	//          context.Value는 인증(userID), 트레이싱(requestID) 같은 요청 범위 데이터에만 사용

	// ctx에 값을 레이어로 쌓음 (덮어쓰기 X)
	// Background() → +userID=12345 → +requestID="req-abc-123"
	ctx6 := context.WithValue(context.Background(), contextKey("userID"), 12345)
	ctx6 = context.WithValue(ctx6, contextKey("requestID"), "req-abc-123")

	// ctx 하나만 넘겨도 안에서 userID, requestID 둘 다 꺼낼 수 있음
	processRequest(ctx6)
	fmt.Println()

	// ----------------------------------------
	// 6. context 전파 (실무 핵심 패턴)
	// ----------------------------------------
	fmt.Println("6. context 전파 - 취소 신호가 하위로 전달됨")
	// HTTP 요청 → DB 쿼리 → 외부 API 처럼
	// 상위 context가 취소되면 하위 context도 자동으로 취소

	parentCtx, parentCancel := context.WithCancel(context.Background())

	// 하위 context 생성 (parent에서 파생)
	childCtx, childCancel := context.WithTimeout(parentCtx, 1*time.Second)
	defer childCancel()

	go func() {
		select {
		case <-childCtx.Done():
			fmt.Printf("child context 취소됨: %v\n", childCtx.Err())
		}
	}()

	time.Sleep(50 * time.Millisecond)
	parentCancel() // 부모 취소 → 자식도 자동 취소
	time.Sleep(50 * time.Millisecond)
	fmt.Println()

	// ----------------------------------------
	// 7. 실무 패턴: HTTP 핸들러에서 context 사용
	// ----------------------------------------
	fmt.Println("7. 실무 패턴 - HTTP 핸들러")
	fmt.Println(`
// 실무에서 이런 식으로 사용
func handleGetUser(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()  // HTTP 요청의 context 가져옴

    // DB 쿼리에 context 전달 → 요청 취소되면 쿼리도 취소
    user, err := db.QueryContext(ctx, "SELECT * FROM users WHERE id = ?", id)

    // 외부 API 호출에도 context 전달
    req, _ := http.NewRequestWithContext(ctx, "GET", apiURL, nil)
    resp, err := http.DefaultClient.Do(req)
}
	`)

	// ----------------------------------------
	// Java vs Go 비교
	// ----------------------------------------
	fmt.Println("=== Java vs Go 비교 ===")
	fmt.Println(`
Java:                               Go:
// Java에 직접 대응 개념 없음        // context로 통일
// 각각 따로 구현해야 함

// 타임아웃                          ctx, cancel := context.WithTimeout(
Future.get(3, TimeUnit.SECONDS)         context.Background(), 3*time.Second)
                                    defer cancel()

// 취소                              ctx, cancel := context.WithCancel(
future.cancel(true)                     context.Background())
                                    cancel()

// 요청 데이터 전달                  ctx := context.WithValue(ctx,
ThreadLocal<String> requestId           "requestID", "abc")
requestId.set("abc")
requestId.get()                     ctx.Value("requestID")
	`)

	// ----------------------------------------
	// 실습 문제!
	// ----------------------------------------
	fmt.Println("\n=== 실습 문제 ===")

	// TODO 1: 3초 타임아웃이 있는 함수를 만드세요
	//   작업이 2초 걸리면 → 성공
	//   작업이 4초 걸리면 → "타임아웃!" 출력

	// TODO 2: context.WithValue로 사용자 정보를 전달하세요
	//   ctx에 userID=999, role="admin" 저장
	//   함수에서 꺼내서 출력
	//   예상 출력: "userID: 999, role: admin"

	fmt.Println("위의 TODO를 직접 구현해보세요!")
	fmt.Println("\n학습 완료! 다음은 13_http.go 로 가봐요.")
}

// ============================================
// 함수 정의들
// ============================================

// context를 받아서 타임아웃/취소를 처리하는 함수
func doWithTimeout(ctx context.Context, workDuration time.Duration) error {
	// "작업이 끝났다"는 신호를 보내기 위한 채널 (데이터 전달 X, 신호 전달 O)
	done := make(chan struct{})

	go func() {
		time.Sleep(workDuration) // 실제 작업 시뮬레이션 (DB 쿼리, API 호출 등)
		close(done)              // 작업 끝! → done 채널 닫아서 신호 보냄
	}()

	// <- 연산: 채널이 열려있으면 값 올 때까지 대기, 닫히면 즉시 실행
	// done과 ctx.Done() 중 먼저 닫히는 쪽 실행
	select {
	case <-done: // close(done) 호출 → 채널 닫힘 → 즉시 실행 → 성공
		return nil
	case <-ctx.Done(): // 타임아웃 되면 → 채널 닫힘 → 즉시 실행 → 에러 반환
		return ctx.Err()
	}
}

// context에서 값을 꺼내서 사용하는 함수
func processRequest(ctx context.Context) {
	userID := ctx.Value(contextKey("userID"))
	requestID := ctx.Value(contextKey("requestID"))

	fmt.Printf("userID: %v, requestID: %v\n", userID, requestID)
}
