// "Build Web Application with Golang"의 1.2장에 대한 예제 코드
// 목적 :이 파일을 실행하여 작업 공간이 올바르게 설정되었는지 확인하기위함
// 실행하려면 콘솔에서 현재 디렉토리로 이동하고 `go run main.go`를 입력하십시오.
// 텍스트 "Hello World"가 표시되지 않으면 작업 영역을 다시 설정하십시오.
package main

import (
	"fmt"
	"mymath"
)

func main() {
	fmt.Printf("Hello, world.  Sqrt(2) = %v\n", mymath.Sqrt(2))
}
