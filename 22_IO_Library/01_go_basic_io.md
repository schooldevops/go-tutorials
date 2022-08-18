# Go IO 

- go io 는 가장 기본적인 input/output 처리를 위한 기본 인터페이스를 제공한다. 

## 파일 복사하기.

```go
func Copy(dst Writer, src Reader) (written int64, err error)
```

- 상기 인터페이스를 이용하면 파일 복사를 수행할 수 있다. 
- src에서 EOF가 만날때까지 혹은 오류를 만날때까지 파일을 복사한다. 
- 반환값은 복사한 바이트 총수 혹은 에러 코드를 반환한다. 

```go

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	sourceFile, err := os.Open("source.txt")

	if err != nil {
		fmt.Println("Error occur during read source file")
	}

	defer sourceFile.Close()

	destFile, err := os.OpenFile("dest.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("Error make dest file")
	}

	defer destFile.Close()

	readByteLen, err := io.Copy(io.Writer(destFile), io.Reader(sourceFile))

	fmt.Println("Read File and Copy File to DestFile. readByte: ", readByteLen)

}
```

- os.Open을 이용하여 이미 존재하는 파일 핸들러를 얻어낸다.
- os.OpenFile을 이용하여 파일 핸들러를 생성하거나 핸들러를 얻는다. (파일명, 파일플래그, 모드)
- io.Copy(dest Writer, src Reader) 를 이용하염 소스의 내용을 dest로 복사한다. 최종 반환값은 복사한 바이트수이다. 

- 상기 내용은 "source.txt" 파일내용을 읽어 "dest.txt" 파일에 쓰기를 수행한다.
- 만약 dest.txt 파일이 존재하지 않는다면 해당 파일을 생성하도록 OpenFile을 이용했다. 

## Buffer를 이용한 파일 복사 

- 위 내용은 내용을 읽고, 바로 대상 파일에 쓰는 작업을 수행한다. 
- 그러나 파일 용량이 매우 크다면 버퍼를 활용하여 파일 쓰기를 수행하면 작은 메모리를 활용하여 파일을 복사할 수 있다. 

```go
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	buff := make([]byte, 1024)

	sourceFile, err := os.Open("source.txt")
	if err != nil {
		fmt.Errorf("Fail to oepn source file/n", err)
	}

	defer sourceFile.Close()

	destFile, err := os.OpenFile("dest.txt", os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		fmt.Errorf("Fail to open dest file/n", err)
	}

	defer destFile.Close()

	for {
		readSize, err := sourceFile.Read(buff)
		if err != nil && err != io.EOF {
			fmt.Errorf("Error occur when read file data.", err)
		}

		if readSize == 0 {
			fmt.Println("size is 0")
			break
		}

		writeSize, err := destFile.Write(buff[:readSize])
		if err != nil {
			fmt.Errorf("Error occur write data to dest file", err)
		}
		fmt.Printf("WriteFile Size is %d \n", writeSize)
	}

	fmt.Println("Read File and Copy File to DestFile was done.")

}
```

- io.Open을 이용하여 소스 파일을 읽어 들인다. 
- io.OpenFile을 이용하여 파일이 존재하지 않는경우 생성하고, 읽기/쓰기 모드로 핸들러를 가져온다. 
- buffer를 생성하여 파일을 읽는다. 
- source 파일 핸들러를 Read로 buffer 크기만큼 읽어 들인다. 
- desc 파일 핸들러를 이용하여 버퍼의 파일을 쓰기한다. 이때 중요한 것은 버퍼를 읽은 크기만큼 슬라이스해야한다. 아니면 찌꺼기 파일을 쓰게 될 수도 있다.

## MultiReader 이용하기. 

- MultiReader를 이용하면 복수개의 Reader를 순서대로 읽어들인다. 

```go
package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r1 := strings.NewReader("First Data.....")
	r2 := strings.NewReader("Second Data.....")
	r3 := strings.NewReader("Third Data.....")

	r := io.MultiReader(r1, r2, r3)

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
}

```

- strings 에서 새로운 리더를 3개 생성하고, 문자열을 기술하였다. 이렇게 하면 각각 문자열을 리더로 읽을 수 있다. 
- io.MultiReader를 이용하여 Reader를 순서로 작성한다. 
- io.Copy를 이용하여 표준 출력으로 읽어들인 데이터를 출력한다. 

- 결과는 다음과 같다. 

```go
go run 02_nulti_reader.go 
First Data.....Second Data.....Third Data....
```