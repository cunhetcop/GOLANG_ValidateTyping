package methods

import (
	"fmt"
	"strconv"
	"time"

	"github.com/fatih/color"
)

func CheckCondition() {
	_, err := getValidNumber3times()//chỗ ẩn này là n
	if err != nil {
		fmt.Println(err)
		return
	}

	//hàm gì đó rồi truyền n vào đây, ví dụ như tính giai thừa: printFactorial(n)
}

func getValidNumber3times() (n int, err error) {
	//dừng chương trình sau 5 giây nếu không có phản hồi
	color.Yellow("ĐÂY LÀ CHƯƠNG TRÌNH TÌM N SỐ NGUYÊN TỐ ĐẦU TIÊN")
	color.Green("Nhập một số nguyên dương: ")
	time.AfterFunc(5*time.Second, func() {
		color.Red("Chương trình đã tự động kết thúc sau 5 giây không có phản hồi")
		panic("timeout")
	})

	//bắt đầu nhập từ bàn phím
	var input string
	count := 0
	for count < 3 {
		//nhập số
		_, err = fmt.Scanln(&input)
		if err != nil {
			fmt.Printf("Nhập sai %d lần \n", count+1)
			count++
			continue
		}
		//nhập string
		n, err = strconv.Atoi(input)
		if err != nil || n <= 0 {
			fmt.Printf("Nhập sai %d lần \n", count+1)
			count++
			continue
		}

		return n, nil
	}
	color.Red("Bạn đã nhập sai quá 3 lần. Game over.")
	return
}
