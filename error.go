package main 

import "fmt"

type MyError struct {
	message string
}

func (err MyError) Error() string {
	return err.message
}

func main() {
	fmt.Printf("'0':%d\n",'0')
	fmt.Printf("'a':%d\n",'a')
	fmt.Printf("'f':%d\n",'f')

	val, err:= hex2int("10")
	fmt.Println(val, err)

	val, err= hex2int("ff")
	fmt.Println(val, err)

	val, err= hex2int("z")
	fmt.Println(val, err)
}



func hex2int(hex string) (val int, err error) {
	for _, r:= range hex {
		fmt.Println("r:",r)

        fmt.Printf("val前:%d\n",val)
		val *= 16
		fmt.Printf("val後:%d\n",val)
		switch {
		case '0' <= r && r <= '9':
			val += int(r - '0')

        case 'a' <= r && r <= 'f':
            val += int(r - 'a')+10

		default:
			return 0, MyError{"不正な文字列です。:" + string(r) }
		}
	}
	return
}