package main

import "fmt"

// ============================================
// Day 1-4: Slice & Map
// ============================================
//
// Java와 비교:
//   Java: ArrayList  → Go: Slice
//   Java: HashMap    → Go: Map
//
// 실무에서 가장 많이 쓰는 자료구조 2가지!

// ============================================
// SLICE
// ============================================

func main() {
	fmt.Println("=== Slice & Map 학습 ===\n")

	// ----------------------------------------
	// 1. Slice 선언 방법
	// ----------------------------------------
	fmt.Println("1. Slice 선언 방법")

	// 방법 1: 리터럴로 선언 (값이 있을 때)
	fruits := []string{"사과", "바나나", "오렌지"}
	fmt.Printf("fruits: %v\n", fruits)

	// 방법 2: make로 선언 (크기를 미리 알 때)
	// make([]타입, 길이, 용량)
	scores := make([]int, 3)      // [0, 0, 0] 길이 3짜리 슬라이스
	scores2 := make([]int, 3, 10) // 길이 3, 용량 10 (내부 배열은 10칸 확보)
	fmt.Printf("scores: %v\n", scores)
	fmt.Printf("scores2: %v (len=%d, cap=%d)\n", scores2, len(scores2), cap(scores2))

	// 방법 3: var로 선언 (nil slice)
	var nilSlice []int  // nil (아직 메모리 할당 안 됨)
	emptySlice := []int{} // empty (메모리는 할당됨, 길이 0)

	fmt.Printf("nilSlice == nil: %t\n", nilSlice == nil)   // true
	fmt.Printf("emptySlice == nil: %t\n\n", emptySlice == nil) // false

	// ----------------------------------------
	// 2. nil slice vs empty slice (중요! ⭐)
	// ----------------------------------------
	fmt.Println("2. nil slice vs empty slice")

	// 📌 JSON 응답에서 차이 발생!
	// nil slice   → JSON: null
	// empty slice → JSON: []
	//
	// API 응답에서 데이터가 없을 때 null 대신 []을 보내려면
	// 반드시 empty slice를 써야 함!

	var result []string   // null 로 직렬화됨 ← API 응답에서 주의!
	result2 := []string{} // [] 로 직렬화됨  ← 실무에서 이걸 씀

	fmt.Printf("nil slice:   %v\n", result)
	fmt.Printf("empty slice: %v\n\n", result2)

	// ----------------------------------------
	// 3. append (요소 추가)
	// ----------------------------------------
	fmt.Println("3. append로 요소 추가")

	// Java: list.add("포도")
	// Go:   fruits = append(fruits, "포도")
	// 📌 반드시 자기 자신에 재할당해야 함!
	fruits = append(fruits, "포도")
	fruits = append(fruits, "딸기", "수박") // 여러 개 한번에 추가
	fmt.Printf("fruits: %v\n\n", fruits)

	// ----------------------------------------
	// 4. len, cap (길이, 용량)
	// ----------------------------------------
	fmt.Println("4. len vs cap")

	s := make([]int, 3, 5)
	fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
	// len: 실제 사용 중인 길이
	// cap: 내부 배열의 총 크기 (재할당 없이 쓸 수 있는 최대)

	s = append(s, 1)
	fmt.Printf("append 1개 후 → len=%d, cap=%d\n", len(s), cap(s))

	s = append(s, 2)
	fmt.Printf("append 1개 후 → len=%d, cap=%d\n", len(s), cap(s))

	s = append(s, 3) // cap 초과 → 내부적으로 새 배열 할당 (cap 자동 증가)
	fmt.Printf("cap 초과 후  → len=%d, cap=%d\n\n", len(s), cap(s))

	// ----------------------------------------
	// 5. 슬라이싱 (부분 추출)
	// ----------------------------------------
	fmt.Println("5. 슬라이싱")

	nums := []int{0, 1, 2, 3, 4, 5}

	fmt.Printf("nums:       %v\n", nums)
	fmt.Printf("nums[1:4]:  %v\n", nums[1:4]) // 인덱스 1~3
	fmt.Printf("nums[:3]:   %v\n", nums[:3])  // 처음~인덱스 2
	fmt.Printf("nums[3:]:   %v\n\n", nums[3:]) // 인덱스 3~끝

	// ----------------------------------------
	// 6. range로 순회
	// ----------------------------------------
	fmt.Println("6. Slice 순회")

	names := []string{"Alice", "Bob", "Charlie"}

	for i, name := range names {
		fmt.Printf("[%d] %s\n", i, name)
	}
	fmt.Println()

	// ----------------------------------------
	// 7. Slice 복사 주의사항 (중요! ⭐)
	// ----------------------------------------
	fmt.Println("7. Slice 복사 주의사항")

	original := []int{1, 2, 3}
	shared := original    // ❌ 주소 공유! 같은 배열을 가리킴
	shared[0] = 999

	fmt.Printf("original: %v\n", original) // [999 2 3] ← 같이 바뀜!
	fmt.Printf("shared:   %v\n", shared)

	// ✅ 진짜 복사하려면 copy() 사용
	original2 := []int{1, 2, 3}
	copied := make([]int, len(original2))
	copy(copied, original2)
	copied[0] = 999

	fmt.Printf("\noriginal2: %v\n", original2) // [1 2 3] ← 안 바뀜!
	fmt.Printf("copied:    %v\n\n", copied)

	// ============================================
	// MAP
	// ============================================

	// ----------------------------------------
	// 8. Map 선언 방법
	// ----------------------------------------
	fmt.Println("8. Map 선언 방법")

	// 방법 1: 리터럴로 선언
	// Java: Map<String, Int> = new HashMap<>()
	userScores := map[string]int{
		"홍길동": 95,
		"김철수": 87,
		"이영희": 92,
	}
	fmt.Printf("userScores: %v\n", userScores)

	// 방법 2: make로 선언
	config := make(map[string]string)
	config["host"] = "localhost"
	config["port"] = "8080"
	fmt.Printf("config: %v\n\n", config)

	// ----------------------------------------
	// 9. Map CRUD
	// ----------------------------------------
	fmt.Println("9. Map CRUD")

	m := map[string]int{}

	// Create / Update
	m["apple"] = 1
	m["banana"] = 2
	m["cherry"] = 3
	fmt.Printf("추가 후: %v\n", m)

	// Read
	fmt.Printf("apple: %d\n", m["apple"])

	// 📌 없는 키 조회 → 에러 아님! 제로값 반환
	fmt.Printf("없는키: %d\n", m["없는키"]) // 0 반환 (int 제로값)

	// ✅ 키 존재 여부 확인 (실무 필수 패턴!)
	// Java: map.containsKey("apple")
	// Go:   value, ok := map[key]
	value, ok := m["apple"]
	fmt.Printf("apple 존재: %t, 값: %d\n", ok, value)

	value2, ok2 := m["없는키"]
	fmt.Printf("없는키 존재: %t, 값: %d\n\n", ok2, value2)

	// Delete
	delete(m, "banana")
	fmt.Printf("banana 삭제 후: %v\n\n", m)

	// ----------------------------------------
	// 10. Map 순회
	// ----------------------------------------
	fmt.Println("10. Map 순회")

	// 📌 Map은 순서 보장 안 됨! 매번 순서가 다를 수 있음
	// Java의 HashMap과 동일
	for key, value := range userScores {
		fmt.Printf("%s: %d점\n", key, value)
	}
	fmt.Println()

	// ----------------------------------------
	// 11. Struct와 함께 쓰는 실전 패턴
	// ----------------------------------------
	fmt.Println("11. 실전 패턴: map[string]struct")

	type Product struct {
		Name  string
		Price int
	}

	products := map[string]Product{
		"p001": {Name: "맥북", Price: 2000000},
		"p002": {Name: "아이폰", Price: 1200000},
	}

	for id, product := range products {
		fmt.Printf("[%s] %s: %d원\n", id, product.Name, product.Price)
	}
	fmt.Println()

	// ----------------------------------------
	// Java vs Go 비교
	// ----------------------------------------
	fmt.Println("=== Java vs Go 비교 ===")
	fmt.Println(`
Java:                               Go:
new ArrayList<>()                   []string{}
list.add("a")                       slice = append(slice, "a")
list.get(0)                         slice[0]
list.size()                         len(slice)
list.subList(1, 3)                  slice[1:3]

new HashMap<>()                     map[string]int{}
map.put("key", 1)                   m["key"] = 1
map.get("key")                      m["key"]
map.containsKey("key")              _, ok := m["key"]
map.remove("key")                   delete(m, "key")
	`)

	// ----------------------------------------
	// 실습 문제!
	// ----------------------------------------
	fmt.Println("\n=== 실습 문제 ===")

	// TODO 1: 아래 학생 점수를 map으로 만들고
	//   "Alice":90, "Bob":75, "Charlie":85, "Dave":60
	//   평균 점수를 구해서 출력하세요

	// TODO 2: 위 map에서 점수가 80 이상인 학생만
	//   []string 슬라이스에 담아서 출력하세요

	// TODO 3: nil slice와 empty slice를 각각 만들고
	//   둘 다 append로 요소를 추가해보세요
	//   (둘 다 append가 되는지 확인!)

	fmt.Println("위의 TODO를 직접 구현해보세요!")
	fmt.Println("\n학습 완료! 다음은 Day 2 - 05_functions.go 로 가봐요.")
}
