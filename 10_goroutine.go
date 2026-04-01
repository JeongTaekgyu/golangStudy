package main

import (
	"fmt"
	"sync"
	"time"
)

// ============================================
// Day 3-1: Goroutine (Go의 동시성 핵심)
// ============================================
//
// Goroutine이란?
//   Go에서 동시에 실행되는 가벼운 실행 단위(경량 스레드)
//   Java의 Thread와 비슷하지만 Go 런타임이 관리해서 훨씬 가벼움
//
// ┌─────────────────────────────────────────────────────┐
// │           Thread vs Goroutine 비교                   │
// ├──────────────┬──────────────────┬───────────────────┤
// │              │ Java Thread      │ Go Goroutine      │
// ├──────────────┼──────────────────┼───────────────────┤
// │ 메모리       │ 1MB+             │ 2KB (훨씬 가벼움) │
// │ 관리 주체    │ OS               │ Go 런타임         │
// │ 동시 실행    │ 수백~수천개      │ 수만개 가능       │
// │ 생성 방법    │ new Thread(...)  │ go 키워드 하나    │
// └──────────────┴──────────────────┴───────────────────┘
//
// Go 동시성 철학:
//   "공유 메모리로 통신하지 말고, 통신으로 메모리를 공유해라"
//   "Do not communicate by sharing memory;
//    share memory by communicating."

// ============================================
// 메인 함수
// ============================================

func main() {
	fmt.Println("=== Go Goroutine 학습 ===\n")

	// ----------------------------------------
	// 1. Goroutine 기본
	// ----------------------------------------
	fmt.Println("1. Goroutine 기본")

	// Java: new Thread(() -> { ... }).start();
	// Go:   go 키워드 하나로 끝
	go sayHello("goroutine1") // 별도 goroutine에서 실행
	go sayHello("goroutine2")
	go sayHello("goroutine3")
	sayHello("main") // main goroutine에서 실행

	// 📌 주의: main이 끝나면 모든 goroutine도 종료됨
	// goroutine이 실행될 시간을 주기 위해 잠깐 대기
	time.Sleep(100 * time.Millisecond)
	fmt.Println()

	// ----------------------------------------
	// 2. WaitGroup (goroutine 완료 대기)
	// ----------------------------------------
	fmt.Println("2. WaitGroup - goroutine 완료 대기")
	// Java의 CountDownLatch 와 같은 역할
	// time.Sleep 으로 대기하는 건 좋지 않음 → WaitGroup 사용

	var wg sync.WaitGroup
	// sync.WaitGroup: goroutine 완료를 추적하는 카운터 (sync 패키지의 struct)

	for i := 1; i <= 5; i++ {
		wg.Add(1)         // 카운터 +1: goroutine 시작 전에 "goroutine 하나 추가됐다"고 알림
		go func(id int) { // go 키워드: 새 goroutine 생성 및 실행 시작
			defer wg.Done() // 예약: 이 익명 함수(goroutine)가 끝날 때 카운터 -1 (종료시키는 게 아님!)
			doWork(id)      // 실제 작업 실행 (goroutine 생성 아님, 그냥 함수 호출)
		}(i) // ← 이 } 에 도달하면 익명 함수 끝 → goroutine 자동 종료 → defer wg.Done() 실행
	}

	wg.Wait() // 카운터가 0이 될 때까지 대기 → 모든 goroutine 완료 보장
	fmt.Println("모든 작업 완료!")
	fmt.Println()

	// ----------------------------------------
	// 3. 동시 실행 vs 순차 실행 비교
	// ----------------------------------------
	fmt.Println("3. 순차 실행 vs 동시 실행 비교")

	// 순차 실행
	fmt.Println("[순차 실행]")
	start := time.Now()
	for i := 1; i <= 3; i++ {
		slowWork(i) // 하나씩 순서대로 실행
	}
	fmt.Printf("순차 실행 시간: %v\n\n", time.Since(start))

	// 동시 실행
	fmt.Println("[동시 실행]")
	var wg2 sync.WaitGroup
	start2 := time.Now()
	for i := 1; i <= 3; i++ {
		wg2.Add(1)
		go func(id int) {
			defer wg2.Done()
			slowWork(id) // 동시에 실행
		}(i)
	}
	wg2.Wait()
	fmt.Printf("동시 실행 시간: %v\n\n", time.Since(start2))

	// ----------------------------------------
	// 4. Mutex (공유 데이터 보호)
	// ----------------------------------------
	fmt.Println("4. Mutex - 공유 데이터 보호")
	// Mutex = 공유 데이터를 한 번에 하나의 goroutine만 수정하도록 보호하는 잠금
	// Mutual Exclusion (상호 배제) 의 약자
	//
	// 왜 필요하냐면?
	//   여러 goroutine이 같은 변수를 동시에 수정하면 race condition 발생
	//   예) counter=0 일 때 goroutine1, goroutine2 동시에 counter++ 하면
	//       둘 다 0을 읽고 1을 저장 → 2가 돼야 하는데 1이 됨!
	//
	// mu.Lock()   → 잠금: 나 쓰는 동안 다른 goroutine 못 들어옴
	// mu.Unlock() → 잠금 해제: 이제 다른 goroutine 들어와도 됨
	//
	// Java의 synchronized 와 같은 역할

	var mu sync.Mutex
	counter := 0
	var wg3 sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg3.Add(1)
		go func() {
			defer wg3.Done()
			mu.Lock()   // 잠금: 다른 goroutine 접근 차단
			counter++   // 안전하게 변경
			mu.Unlock() // 잠금 해제
		}()
	}

	wg3.Wait()
	fmt.Printf("counter = %d (100이어야 정상)\n\n", counter)

	// ----------------------------------------
	// Java vs Go 비교
	// ----------------------------------------
	fmt.Println("=== Java vs Go 비교 ===")
	fmt.Println(`
Java:                               Go:
// Thread 생성                      // goroutine 생성
new Thread(() -> {                  go func() {
    doWork();                           doWork()
}).start();                         }()

// CountDownLatch                   // WaitGroup
CountDownLatch latch =              var wg sync.WaitGroup
    new CountDownLatch(5);
latch.countDown();                  wg.Done()
latch.await();                      wg.Wait()

// synchronized                     // Mutex
synchronized(lock) {                mu.Lock()
    count++;                        count++
}                                   mu.Unlock()
	`)

	// ----------------------------------------
	// 실습 문제!
	// ----------------------------------------
	fmt.Println("\n=== 실습 문제 ===")

	// TODO 1: 아래 작업을 goroutine으로 동시에 실행하세요
	//   작업 목록: ["이미지 처리", "이메일 전송", "DB 저장", "캐시 업데이트"]
	//   각 작업은 100ms 걸린다고 가정 (time.Sleep 사용)
	//   WaitGroup으로 모든 작업 완료 대기
	//   예상 출력:
	//     이미지 처리 시작
	//     이메일 전송 시작
	//     DB 저장 시작
	//     캐시 업데이트 시작
	//     모든 작업 완료! (약 100ms 만에)

	// TODO 2: 안전한 카운터를 만드세요
	//   50개의 goroutine이 동시에 counter를 2씩 증가
	//   Mutex로 race condition 방지
	//   최종 결과: counter = 100

	fmt.Println("위의 TODO를 직접 구현해보세요!")
	fmt.Println("\n학습 완료! 다음은 11_channel.go 로 가봐요.")
}

// ============================================
// 함수 정의들
// ============================================

func sayHello(name string) {
	fmt.Printf("안녕하세요! 저는 %s입니다.\n", name)
}

func doWork(id int) {
	fmt.Printf("작업 %d 시작\n", id)
	time.Sleep(50 * time.Millisecond) // 작업 시뮬레이션
	fmt.Printf("작업 %d 완료\n", id)
}

func slowWork(id int) {
	fmt.Printf("  slowWork %d 시작\n", id)
	time.Sleep(100 * time.Millisecond) // 100ms 걸리는 작업
	fmt.Printf("  slowWork %d 완료\n", id)
}
