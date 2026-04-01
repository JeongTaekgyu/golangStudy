package main

import (
	"fmt"
	"sync"
	"time"
)

// ============================================
// Day 3-2: Channel (goroutine 간 통신)
// ============================================
//
// Channel이란?
//   goroutine 간에 데이터를 주고받는 통로
//   Go의 동시성 철학의 핵심
//
//   "공유 메모리로 통신하지 말고, 통신으로 메모리를 공유해라"
//   → Mutex로 공유 변수 보호하는 대신 channel로 데이터 전달
//
// 기본 문법:
//   ch := make(chan int)  // int를 주고받는 channel 생성
//   ch <- 42             // channel에 값 전송 (send)
//   v := <-ch            // channel에서 값 수신 (receive)
//
// ┌─────────────────────────────────────────────────────┐
// │        Unbuffered vs Buffered Channel               │
// ├──────────────┬──────────────────┬───────────────────┤
// │              │ Unbuffered       │ Buffered          │
// ├──────────────┼──────────────────┼───────────────────┤
// │ 생성         │ make(chan int)    │ make(chan int, 3)  │
// │ 동작         │ 송수신 동시 대기 │ 버퍼 찰 때까지 대기 없음│
// │ 용도         │ 동기화           │ 비동기 처리        │
// └──────────────┴──────────────────┴───────────────────┘

func main() {
	fmt.Println("=== Go Channel 학습 ===\n")

	// ----------------------------------------
	// 1. Channel 기본
	// ----------------------------------------
	fmt.Println("1. Channel 기본")

	// channel 생성
	ch := make(chan int) // int 타입 channel

	// goroutine에서 channel로 값 전송
	go func() {
		fmt.Println("goroutine: 값 전송 중...")
		ch <- 42 // channel에 42 전송
		fmt.Println("goroutine: 전송 완료")
	}()

	// main에서 channel로부터 값 수신
	v := <-ch // channel에서 값 받을 때까지 대기
	fmt.Printf("main: channel에서 받은 값 = %d\n\n", v)

	// ----------------------------------------
	// 2. Channel로 goroutine 결과 받기
	// ----------------------------------------
	fmt.Println("2. Channel로 goroutine 결과 받기")
	// WaitGroup 없이 channel로 완료 신호 받기

	resultCh := make(chan string)

	go func() {
		time.Sleep(50 * time.Millisecond) // 작업 시뮬레이션
		resultCh <- "작업 완료!"              // 결과를 channel로 전송
	}()

	result := <-resultCh // 결과 받을 때까지 대기
	fmt.Printf("결과: %s\n\n", result)

	// ----------------------------------------
	// 3. 여러 goroutine 결과 수집
	// ----------------------------------------
	fmt.Println("3. 여러 goroutine 결과 수집")

	jobs := []int{1, 2, 3, 4, 5}
	resultsCh := make(chan int, len(jobs)) // 버퍼드 channel

	for _, job := range jobs {
		go func(n int) {
			resultsCh <- n * n // 제곱값 전송
		}(job)
	}

	// 모든 결과 수집
	for range jobs {
		fmt.Printf("결과: %d\n", <-resultsCh)
	}
	fmt.Println()

	// ----------------------------------------
	// 4. Buffered Channel
	// ----------------------------------------
	fmt.Println("4. Buffered Channel")

	// Unbuffered: 수신자 없으면 송신자가 블로킹됨
	// Buffered: 버퍼가 찰 때까지 블로킹 없이 전송 가능
	bufferedCh := make(chan string, 3) // 버퍼 크기 3

	// 수신자 없어도 버퍼가 있으니까 바로 전송 가능
	bufferedCh <- "첫번째"
	bufferedCh <- "두번째"
	bufferedCh <- "세번째"
	// bufferedCh <- "네번째"  // ← 버퍼 꽉 참 → 블로킹!

	fmt.Println(<-bufferedCh) // "첫번째"
	fmt.Println(<-bufferedCh) // "두번째"
	fmt.Println(<-bufferedCh) // "세번째"
	fmt.Println()

	// ----------------------------------------
	// 5. Channel 방향 지정
	// ----------------------------------------
	fmt.Println("5. Channel 방향 지정")
	// 함수 파라미터에서 channel 방향을 제한할 수 있음
	// chan<- : 송신 전용
	// <-chan : 수신 전용

	dirCh := make(chan int)

	go sendOnly(dirCh) // 새 goroutine으로 실행: 송신/수신이 동시에 있어야 하기 때문 (없으면 deadlock)
	receiveOnly(dirCh) // main goroutine에서 실행: 채널에서 값 수신
	fmt.Println()

	// ----------------------------------------
	// 6. select (여러 channel 동시 처리)
	// ----------------------------------------
	fmt.Println("6. select - 여러 channel 동시 처리")
	// Java의 switch와 비슷하지만 channel용
	// 여러 channel 중 준비된 것부터 처리

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch1 <- "ch1 완료"
	}()

	go func() {
		time.Sleep(30 * time.Millisecond)
		ch2 <- "ch2 완료"
	}()

	// 2개의 channel 중 먼저 도착하는 것부터 처리
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("ch1에서 받음: %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("ch2에서 받음: %s\n", msg2)
		}
	}
	fmt.Println()

	// ----------------------------------------
	// 7. select + default (non-blocking)
	// ----------------------------------------
	fmt.Println("7. select + default (non-blocking)")
	// blocking     = 채널에 값이 올 때까지 무조건 대기 (그 동안 다른 작업 못 함)
	// non-blocking = 채널에 값이 없으면 대기 없이 바로 다른 작업 실행
	//
	// default가 있으면 → non-blocking (채널에 값 없으면 바로 default 실행)
	// default가 없으면 → blocking    (채널에 값 올 때까지 무조건 대기)
	//
	// 실무 사용 예:
	//   캐시 조회 시 캐시 있으면 바로 반환, 없으면 DB 조회
	//   작업 큐가 비어있으면 대기 없이 다른 작업 처리

	nonBlockCh := make(chan int, 1) // 체날 생성

	// channel이 비어있을 때
	select {
	case v := <-nonBlockCh:
		fmt.Printf("받은 값: %d\n", v)
	default:
		fmt.Println("channel 비어있음 → default 실행")
	}

	nonBlockCh <- 99

	// channel에 값이 있을 때
	select {
	case v := <-nonBlockCh:
		fmt.Printf("받은 값: %d\n", v)
	default:
		fmt.Println("channel 비어있음 → default 실행")
	}
	fmt.Println()

	// ----------------------------------------
	// 8. Channel로 done 신호 보내기 (실무 패턴)
	// ----------------------------------------
	fmt.Println("8. done channel 패턴")
	// goroutine 종료 신호를 channel로 전달하는 패턴

	done := make(chan struct{}) // 값 없이 신호만 보낼 때 struct{} 사용

	go func() {
		fmt.Println("goroutine: 작업 시작")
		time.Sleep(50 * time.Millisecond)
		fmt.Println("goroutine: 작업 완료")
		done <- struct{}{} // 완료 신호 전송
	}()

	<-done // 완료 신호 받을 때까지 대기
	fmt.Println("main: goroutine 완료 확인")
	fmt.Println()

	// ----------------------------------------
	// 9. Pipeline 패턴 (실무 패턴)
	// ----------------------------------------
	fmt.Println("9. Pipeline 패턴")
	// 여러 goroutine을 channel로 연결해서 데이터 처리
	// [생성] → channel → [처리] → channel → [출력]

	numbers := generate(1, 2, 3, 4, 5) // 숫자 생성
	squared := square(numbers)         // 제곱 처리
	printResults(squared)              // 출력
	fmt.Println()

	// ----------------------------------------
	// 10. Mutex vs Channel 비교
	// ----------------------------------------
	fmt.Println("10. Mutex vs Channel 선택 기준")
	fmt.Println(`
Mutex를 쓰는 경우:
  → 공유 변수를 여러 goroutine이 읽고 써야 할 때
  → 캐시, 카운터 등 상태 관리
  예) mu.Lock(); counter++; mu.Unlock()

Channel을 쓰는 경우:
  → goroutine 간에 데이터를 전달할 때
  → 작업 결과를 수집할 때
  → goroutine 완료 신호를 보낼 때
  예) resultCh <- result
	`)

	// ----------------------------------------
	// Java vs Go 비교
	// ----------------------------------------
	fmt.Println("=== Java vs Go 비교 ===")
	fmt.Println(`
Java:                               Go:
// BlockingQueue로 데이터 전달       // Channel로 데이터 전달
BlockingQueue<Integer> q =          ch := make(chan int)
    new LinkedBlockingQueue<>();
q.put(42);    // 전송               ch <- 42   // 전송
q.take();     // 수신               <-ch        // 수신

// 결과 없음, 직접 구현 필요         // select로 여러 channel 처리
if (q1.peek() != null) {            select {
    process(q1.take());             case v := <-ch1:
} else if (q2.peek() != null) {         process(v)
    process(q2.take());             case v := <-ch2:
}                                       process(v)
                                    }
	`)

	// ----------------------------------------
	// 실습 문제!
	// ----------------------------------------
	fmt.Println("\n=== 실습 문제 ===")

	// TODO 1: channel을 이용한 계산기를 만드세요
	//   두 숫자를 channel로 받아서 합계를 반환하는 goroutine
	//   numCh := make(chan int, 2)
	//   resultCh := make(chan int)
	//   numCh <- 10
	//   numCh <- 20
	//   → resultCh 에서 30 받기

	// TODO 2: timeout 패턴을 구현하세요
	//   작업이 200ms 안에 완료되면 결과 출력
	//   200ms 초과하면 "타임아웃!" 출력
	//   힌트: select + time.After(200 * time.Millisecond)

	fmt.Println("위의 TODO를 직접 구현해보세요!")
	fmt.Println("\n학습 완료! 다음은 12_context.go 로 가봐요.")
}

// ============================================
// 함수 정의들
// ============================================

// 송신 전용 channel 파라미터 (chan<-)
func sendOnly(ch chan<- int) {
	ch <- 100
	fmt.Println("sendOnly: 100 전송")
}

// 수신 전용 channel 파라미터 (<-chan)
func receiveOnly(ch <-chan int) {
	v := <-ch
	fmt.Printf("receiveOnly: %d 수신\n", v)
}

// Pipeline 패턴 함수들
func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out) // 전송 완료 후 channel 닫기
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in { // channel이 닫힐 때까지 수신
			out <- n * n
		}
		close(out)
	}()
	return out
}

func printResults(in <-chan int) {
	var wg sync.WaitGroup
	for n := range in {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			fmt.Printf("결과: %d\n", v)
		}(n)
	}
	wg.Wait()
}
