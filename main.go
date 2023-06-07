package main

import (
	"fmt"
	"regexp"
)

func main() {
	// 正则表达式模式
	pattern := `\nCVE-\d+-\d+.*(\d+\.\d+).*\n`

	// 待匹配的文本
	text := `漏洞评分( 漏洞编号 危害程度 CVSS 3.1 评分 漏洞类型)
CVE-2023-0266 重要 7.8 其他
漏洞评分向量：CVSS:3.1/AV:L/AC:L/PR:L/UI:N/S:U/C:H/I:H/A:H`

	// 编译正则表达式
	re := regexp.MustCompile(pattern)

	// 使用正则表达式进行匹配
	match := re.FindStringSubmatch(text)

	if len(match) >= 4 {
		cveID := match[1]
		severity := match[2]
		score := match[3]

		fmt.Println("CVE ID:", cveID)
		fmt.Println("Severity:", severity)
		fmt.Println("Score:", score)
	}
	fmt.Println("end,", match)
}
