package main

import "fmt"

// ============================================
// Day 1-2: Go의 반복문 (for문만 존재!)
// ============================================

func main() {
	fmt.Println("=== Go for문 학습 ===")
	fmt.Println("⭐ Go에는 for문만 있습니다! (while, do-while 없음)\n")

	// ----------------------------------------
	// 1. 기본 for문 (Java와 동일)
	// ----------------------------------------
	fmt.Println("1. 기본 for문")

	// Java: for (int i = 0; i < 5; i++)
	// Go:   for i := 0; i < 5; i++
	for i := 0; i < 5; i++ {
		fmt.Printf("i = %d\n", i)
	}
	fmt.Println()

	// ----------------------------------------
	// 2. while문처럼 사용 (조건만)
	// ----------------------------------------
	fmt.Println("2. while문처럼 사용")

	count := 0
	// Java의 while (count < 3)과 동일
	for count < 3 {
		fmt.Printf("count = %d\n", count)
		count++
	}
	fmt.Println()

	// ----------------------------------------
	// 3. 무한 루프
	// ----------------------------------------
	fmt.Println("3. 무한 루프 (break로 탈출)")

	num := 0
	// Java의 while(true)와 동일
	for {
		fmt.Printf("num = %d\n", num)
		num++

		if num >= 3 {
			break // 루프 탈출
		}
	}
	fmt.Println()

	// ----------------------------------------
	// 4. continue 사용
	// ----------------------------------------
	fmt.Println("4. continue (홀수만 출력)")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 { // 짝수면
			continue // 다음 반복으로
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println("\n")

	// ----------------------------------------
	// 5. 중첩 반복문
	// ----------------------------------------
	fmt.Println("5. 중첩 for문 (구구단 2단)")

	dan := 2
	for i := 1; i <= 9; i++ {
		fmt.Printf("%d x %d = %d\n", dan, i, dan*i)
	}
	fmt.Println()

	// ----------------------------------------
	// 6. range를 사용한 for문 (배열/슬라이스)
	// ----------------------------------------
	fmt.Println("6. range를 사용한 for문")

	// 배열 순회
	numbers := []int{10, 20, 30, 40, 50}

	// 인덱스와 값 모두 사용
	fmt.Println("인덱스와 값:")
	for index, value := range numbers {
		fmt.Printf("numbers[%d] = %d\n", index, value)
	}
	fmt.Println()

	// 값만 사용 (인덱스 무시)
	fmt.Println("값만 사용:")
	for _, value := range numbers {
		fmt.Printf("%d ", value)
	}
	fmt.Println("\n")

	// 인덱스만 사용
	fmt.Println("인덱스만 사용:")
	for index := range numbers {
		fmt.Printf("인덱스: %d\n", index)
	}
	fmt.Println()

	// ----------------------------------------
	// 7. range를 사용한 for문 (맵)
	// ----------------------------------------
	fmt.Println("7. map 순회")

	// map 생성 (Java의 HashMap 같은 것)
	userScores := map[string]int{
		"홍길동": 95,
		"김철수": 87,
		"이영희": 92,
	}

	for name, score := range userScores {
		fmt.Printf("%s: %d점\n", name, score)
	}
	fmt.Println()

	// ----------------------------------------
	// 8. range를 사용한 for문 (문자열)
	// ----------------------------------------
	fmt.Println("8. 문자열 순회")

	text := "Hello"
	for index, char := range text {
		// char는 rune 타입 (유니코드 코드포인트)
		fmt.Printf("text[%d] = %c (코드: %d)\n", index, char, char)
	}
	fmt.Println()

	// 한글도 가능
	korean := "안녕"
	for index, char := range korean {
		fmt.Printf("%d: %c\n", index, char)
	}
	fmt.Println()

	// ----------------------------------------
	// 9. 라벨(Label)을 사용한 break
	// ----------------------------------------
	fmt.Println("9. 라벨을 사용한 break (중첩 루프 탈출)")

OuterLoop: // 라벨 선언
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("i=%d, j=%d\n", i, j)
			if i == 2 && j == 2 {
				break OuterLoop // 바깥 루프까지 탈출
			}
		}
	}
	fmt.Println()

	// ----------------------------------------
	// 10. 실전 예제: 합계 구하기
	// ----------------------------------------
	fmt.Println("10. 실전 예제")

	// 1부터 100까지의 합
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Printf("1부터 100까지의 합: %d\n", sum)

	// 짝수만 더하기
	evenSum := 0
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			evenSum += i
		}
	}
	fmt.Printf("1부터 100까지 짝수의 합: %d\n", sum)
	fmt.Println()

	// ----------------------------------------
	// Java vs Go 비교
	// ----------------------------------------
	fmt.Println("=== Java vs Go 비교 ===")
	fmt.Println(`
Java:
  for (int i = 0; i < 5; i++) { }
  while (condition) { }
  do { } while (condition);
  for (int num : numbers) { }

Go:
  for i := 0; i < 5; i++ { }        // 기본 for
  for condition { }                  // while처럼
  for { }                            // 무한 루프
  for _, num := range numbers { }    // range로 순회
	`)

	// ----------------------------------------
	// 실습 문제!
	// ----------------------------------------
	fmt.Println("\n=== 실습 문제 ===")
	fmt.Println("아래 주석을 풀고 직접 코드를 작성해보세요!")

	// TODO 1: 1부터 10까지 출력하는 for문 작성

	// TODO 2: 1부터 50까지 중 홀수의 합을 구하는 코드 작성

	// TODO 3: 다음 배열의 모든 요소를 출력하세요
	// fruits := []string{"사과", "바나나", "오렌지", "포도"}

	// TODO 4: 구구단 3단을 출력하세요 (3 x 1 = 3 ~ 3 x 9 = 27)

	// TODO 5: 다음 map의 모든 키와 값을 출력하세요
	// ages := map[string]int{"Alice": 25, "Bob": 30, "Charlie": 35}

	fmt.Println("\n학습 완료! 다음은 03_structs.go를 만들어볼까요?")
}
