# Go vs Java 비교 정리

> 3년차 Java 백엔드 개발자를 위한 Go 사고방식 전환 가이드

---

## 목차
1. [언어 철학](#1-언어-철학)
2. [메모리 관리 & 런타임](#2-메모리-관리--런타임)
3. [타입 시스템](#3-타입-시스템)
4. [객체지향 vs 컴포지션](#4-객체지향-vs-컴포지션)
5. [에러 처리](#5-에러-처리)
6. [동시성](#6-동시성)
7. [인터페이스](#7-인터페이스)
8. [패키지 & 모듈](#8-패키지--모듈)
9. [빌드 & 배포](#9-빌드--배포)
10. [실무 체감 차이 요약](#10-실무-체감-차이-요약)

---

## 1. 언어 철학

| | Java | Go |
|---|---|---|
| 핵심 가치 | 객체지향, 확장성, 유연성 | 단순함, 명확함, 실용성 |
| 설계 방향 | "어떻게 확장할까?" | "어떻게 단순하게 유지할까?" |
| 문법 키워드 수 | 50개+ | **25개** (매우 적음) |
| 학습 곡선 | 높음 (JVM, 제네릭, 어노테이션 등) | 낮음 (문법 자체는 단순) |
| 마법(Magic) | 많음 (어노테이션, AOP, 리플렉션) | **거의 없음** (명시적) |

**Go의 핵심 철학:**
```
"There is only one way to do it."
하나의 올바른 방법만 존재한다 → 코드 일관성이 높아짐
```

---

## 2. 메모리 관리 & 런타임

### GC (가비지 컬렉터)

| | Java | Go |
|---|---|---|
| GC 방식 | Generational GC | Concurrent Mark & Sweep |
| Stop-The-World | 있음 (수십ms ~ 수초) | **거의 없음** (목표: 1ms 이하) |
| Full GC 문제 | 트래픽 많을 때 응답 튈 수 있음 | 거의 없음 |
| GC 튜닝 | `-Xmx`, `-Xms`, GC 로그 분석 필요 | 옵션 거의 없음, 신경 안 써도 됨 |

### 스택 vs 힙

```java
// Java: 객체는 거의 무조건 힙
User user = new User(); // 힙에 올라감
```

```go
// Go: 컴파일러가 escape analysis로 자동 판단
func createUser() *User {
    u := User{}   // 힙에 올라감 (함수 밖으로 나가니까)
    return &u
}

func process() {
    u := User{}   // 스택에 올라감 (함수 안에서만 씀)
    fmt.Println(u.Name)
}
// → 스택 사용이 많을수록 GC 부담 감소
```

### 런타임 & 바이너리

| | Java | Go |
|---|---|---|
| 실행 방식 | JVM 위에서 바이트코드 실행 | **네이티브 바이너리**로 컴파일 |
| 시작 시간 | 느림 (JVM 웜업) | **매우 빠름** (즉시 시작) |
| 메모리 사용 | 높음 (JVM 자체가 수백MB) | 낮음 |
| 도커 이미지 크기 | Spring Boot: 300~500MB | Go 서버: **10~30MB** |

---

## 3. 타입 시스템

### 변수 선언

```java
// Java
String name = "홍길동";
int age = 30;
var city = "서울";  // Java 10+
```

```go
// Go
var name string = "홍길동"  // 명시적
name := "홍길동"            // 타입 추론 (가장 많이 씀)
var age int = 30
age := 30

// ❌ 전역에서 := 불가
var globalVar = "전역변수"   // 전역은 var만 가능
```

### 타입 변환

```java
// Java: 자동 형변환 있음 (암묵적)
int i = 42;
double d = i;  // 자동 변환 OK
```

```go
// Go: 자동 형변환 없음 (항상 명시적)
var i int = 42
var d float64 = float64(i)  // 명시적 변환 필수
// var d float64 = i  → ❌ 컴파일 에러
```

### 포인터

```java
// Java: 포인터 개념 없음 (참조만 있음)
// 개발자가 메모리 주소를 직접 다루지 않음
User user = new User();
```

```go
// Go: 포인터 직접 사용
// & : 변수의 메모리 주소를 가져옴
// * : 그 주소에 있는 값을 가져옴 (역참조)

x := 10
ptr := &x          // ptr = x의 메모리 주소
fmt.Println(*ptr)  // 10 (주소에 있는 값을 꺼냄)
```

### 포인터가 필요한 이유 1 — 함수에서 원본 변경

```go
// ❌ 값 전달: 복사본이 넘어감 → 원본 변경 불가
func printUser(u User) {
    u.Name = "변경시도"       // 복사본만 바뀜
}

user := User{Name: "홍길동"}
printUser(user)
fmt.Println(user.Name)        // "홍길동" → 원본 그대로!

// ✅ 포인터 전달: 주소가 넘어감 → 원본 직접 변경 가능
func updateUser(u *User) {
    u.Name = "변경성공"       // 주소로 찾아가서 원본 변경
}

user := User{Name: "홍길동"}
updateUser(&user)             // &user = user의 주소를 넘김
fmt.Println(user.Name)        // "변경성공" → 원본이 바뀜!
```

### 포인터가 필요한 이유 2 — 복사 비용 절감

```go
type BigData struct {
    Records [100000]int  // 큰 데이터
}

// ❌ 값 전달: 함수 호출마다 100000개 int를 통째로 복사 → 느림
func process(d BigData) { }

// ✅ 포인터 전달: 주소(8바이트)만 전달 → 빠름
func process(d *BigData) { }
```

### 메서드에서의 포인터 리시버 vs 값 리시버

```go
// 값 리시버: 복사본에서 실행 → 원본 변경 안 됨 (읽기 전용)
func (u User) GetName() string {
    return u.Name
}

// 포인터 리시버: 원본에서 실행 → 원본 변경 가능
func (u *User) ChangeName(name string) {
    u.Name = name
}

// ── 호출 예시 ──
user := User{Name: "홍길동"}

fmt.Println(user.GetName())  // "홍길동" (읽기만)

user.ChangeName("김철수")    // 원본 변경
fmt.Println(user.Name)       // "김철수"
```

---

## 4. 객체지향 vs 컴포지션

### 클래스 → Struct

```java
// Java
public class User {
    private int id;
    private String name;

    public User(int id, String name) {  // 생성자 (언어 기능)
        this.id = id;
        this.name = name;
    }

    public String getName() { return name; }
    public void setName(String name) { this.name = name; }
}
```

```go
// Go
type User struct {
    ID   int
    Name string
    age  int    // private (소문자)
}

// 생성자 없음 → 관례적으로 NewXxx() 함수 사용
func NewUser(id int, name string) *User {
    return &User{ID: id, Name: name}
}

// getter/setter 대신: public 필드 직접 접근 또는 메서드
func (u *User) SetAge(age int) { u.age = age }
func (u User) GetAge() int     { return u.age }
```

### 상속 → 컴포지션 (임베딩)

```java
// Java: 상속
public class Admin extends User {
    private int level;
    // User의 모든 것을 물려받음
}

Admin admin = new Admin();
admin.getName();  // User 메서드 사용 가능
```

```go
// Go: 임베딩 (컴포지션)
type Admin struct {
    User        // 임베딩 → User의 필드/메서드를 그대로 사용
    Level int
}

admin := Admin{User: User{ID: 1, Name: "관리자"}, Level: 5}
admin.Name      // User.Name 직접 접근
admin.GetAge()  // User 메서드 직접 호출
admin.Level     // Admin 자체 필드
```

**왜 상속보다 컴포지션인가?**
```
상속의 문제점:
- 강한 결합 (부모 바뀌면 자식 다 영향)
- 다중 상속 문제 (Diamond Problem)
- 불필요한 메서드까지 물려받음

컴포지션의 장점:
- 필요한 것만 가져다 씀
- 유연하게 교체 가능
- 테스트하기 쉬움
```

### public / private

```java
// Java: 접근 제어자 키워드 사용
public String name;
private int age;
protected String email;
```

```go
// Go: 대소문자로만 구분
Name  string  // 대문자 = public (패키지 외부 접근 가능)
age   int     // 소문자 = private (패키지 내부에서만)
// protected 개념 없음
```

---

## 5. 에러 처리

Go에서 가장 사고방식 전환이 필요한 부분.

```java
// Java: 예외(Exception) 기반
try {
    User user = userRepository.findById(id);
    // 비즈니스 로직
} catch (UserNotFoundException e) {
    // 에러 처리
} catch (DatabaseException e) {
    // DB 에러 처리
} finally {
    // 정리
}
```

```go
// Go: 에러는 값(Value)
// 예외 없음, try/catch 없음
user, err := userRepository.FindByID(id)
if err != nil {
    return fmt.Errorf("사용자 조회 실패: %w", err)  // 에러 래핑
}
// 여기서부터는 user가 정상임이 보장됨
```

**Go 에러 처리 패턴:**

```go
// 1. 기본 패턴
result, err := someFunction()
if err != nil {
    return err
}

// 2. 에러 래핑 (원인 추적)
if err != nil {
    return fmt.Errorf("createOrder 실패: %w", err)
}

// 3. sentinel error (특정 에러 구분)
var ErrNotFound = errors.New("not found")

if errors.Is(err, ErrNotFound) {
    // 404 처리
}

// 4. 에러 타입 구분
var dbErr *DatabaseError
if errors.As(err, &dbErr) {
    // DB 에러 전용 처리
}
```

**📌 핵심 마인드셋:**
```
Java: "예외는 특별한 상황" → try/catch로 분리
Go:   "에러는 그냥 값"    → 반환값으로 처리, 절대 무시 금지
```

---

## 6. 동시성

Java vs Go에서 가장 큰 차이.

### Java: 스레드 기반

```java
// Java: Thread / ExecutorService
Thread thread = new Thread(() -> {
    // 작업
});
thread.start();

// 또는
ExecutorService executor = Executors.newFixedThreadPool(10);
executor.submit(() -> process());

// 공유 메모리 → 동기화 필요
synchronized (lock) {
    sharedData++;
}
```

### Go: Goroutine + Channel

```go
// Go: goroutine (스레드보다 훨씬 가벼움)
go process()  // 그냥 go 키워드 하나로 끝

// goroutine vs thread 차이
// Thread:    1MB+ 메모리, OS가 관리
// Goroutine: 2KB 시작, Go 런타임이 관리 → 수만개 동시 실행 가능

// 채널로 통신 (공유 메모리 대신)
ch := make(chan int)

go func() {
    ch <- 42  // 채널에 전송
}()

value := <-ch  // 채널에서 수신

// Go 철학:
// "공유 메모리로 통신하지 말고, 통신으로 메모리를 공유해라"
// "Do not communicate by sharing memory;
//  share memory by communicating."
```

### WaitGroup (Java의 CountDownLatch)

```go
var wg sync.WaitGroup

for i := 0; i < 5; i++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done()
        process(id)
    }(i)
}

wg.Wait()  // 모든 goroutine 완료 대기
```

---

## 7. 인터페이스

```java
// Java: implements 명시 필요
public interface Stringer {
    String toString();
}

public class User implements Stringer {  // 명시적 선언 필수
    @Override
    public String toString() {
        return name;
    }
}
```

```go
// Go: implements 키워드 없음 → 자동 만족 (Duck Typing)
type Stringer interface {
    String() string
}

type User struct { Name string }

func (u User) String() string {  // 메서드만 있으면 자동으로 Stringer
    return u.Name
}
// User가 Stringer를 구현한다고 어디에도 선언 안 함!
// → 메서드 시그니처가 맞으면 자동으로 인터페이스 만족
```

**Go 인터페이스 철학:**

```go
// ❌ Java스러운 Go (나쁜 예)
type UserServiceInterface interface {
    CreateUser(...)
    UpdateUser(...)
    DeleteUser(...)
    FindUser(...)
    ListUsers(...)
    // 너무 큼
}

// ✅ Go다운 인터페이스 (좋은 예) → 작게, 소비자 입장에서
type UserFinder interface {
    FindByID(id int) (*User, error)
}

type UserCreator interface {
    Create(user *User) error
}
```

---

## 8. 패키지 & 모듈

```java
// Java
package com.company.project.service;

import com.company.project.repository.UserRepository;
import org.springframework.stereotype.Service;

@Service
public class UserService { ... }
```

```go
// Go
package service

import (
    "github.com/company/project/repository"
    "fmt"
)

type UserService struct {
    repo repository.UserRepository  // 인터페이스 의존
}
// 어노테이션 없음, DI 프레임워크 없음 (직접 주입)
```

**의존성 주입 비교:**

```java
// Java: Spring이 알아서 주입
@Autowired
private UserRepository userRepository;
```

```go
// Go: 직접 주입 (명시적)
func NewUserService(repo UserRepository) *UserService {
    return &UserService{repo: repo}
}

// main.go에서
repo := repository.NewUserRepository(db)
service := service.NewUserService(repo)
```

---

## 9. 빌드 & 배포

| | Java | Go |
|---|---|---|
| 빌드 결과물 | `.jar` (JVM 필요) | **단일 바이너리** (의존성 없음) |
| 실행 환경 | JRE 설치 필요 | 그냥 실행 가능 |
| 빌드 시간 | 느림 | **매우 빠름** |
| 크로스 컴파일 | 어려움 | `GOOS=linux go build` 한 줄 |
| 도커 이미지 | 300~500MB | **10~30MB** |

```bash
# Go 크로스 컴파일 (Mac에서 Linux 바이너리 생성)
GOOS=linux GOARCH=amd64 go build -o server .

# 도커파일도 단순해짐
FROM scratch          # 아무것도 없는 베이스
COPY server /server
CMD ["/server"]
```

---

## 10. 실무 체감 차이 요약

| 상황 | Java (Spring) | Go |
|---|---|---|
| 서버 시작 시간 | 10~30초 | **0.1초 이하** |
| 메모리 사용량 | 500MB~1GB | **50~100MB** |
| GC 멈춤 | 가끔 튈 수 있음 | 거의 없음 |
| 도커 이미지 | 크고 무거움 | 작고 가벼움 |
| 코드 양 | 많음 (보일러플레이트) | 적음 |
| DI/AOP | 프레임워크 의존 | 직접 구현 (명시적) |
| 동시성 처리 | 스레드풀 관리 필요 | goroutine으로 단순하게 |
| 에러 처리 | try/catch | if err != nil |

---

## 📌 Go로 넘어올 때 버려야 할 Java 습관

1. **try/catch 사고방식** → `if err != nil` 패턴으로
2. **상속 설계** → 컴포지션(임베딩)으로
3. **인터페이스 먼저 만들기** → 필요할 때 추출하기
4. **어노테이션/매직 의존** → 명시적 코드로
5. **getter/setter 무조건 만들기** → public 필드 직접 접근 or 필요한 메서드만
6. **DI 프레임워크 찾기** → 생성자 함수로 직접 주입
7. **checked exception 설계** → error 값으로 반환
8. **Spring 구조 그대로 복사** → Go 철학에 맞게 단순하게
