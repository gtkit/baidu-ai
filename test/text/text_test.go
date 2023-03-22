// @Author xiaozhaofu 2023/3/22 10:02:00
package text

import (
	"fmt"
	"testing"

	"github.com/Baidu-AIP/golang-sdk/aip/censor"
)

func TestText(t *testing.T) {
	client := censor.NewClient("AK", "SK")
	res := client.TextCensor("")
	fmt.Println(res)
}
