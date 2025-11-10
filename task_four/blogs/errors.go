package blogs

import (
	"log"
	"net/http"
)

// 错误处理与日志记录
// 对可能出现的错误进行统一处理，如数据库连接错误、用户认证失败、文章或评论不存在等，返回合适的 HTTP 状态码和错误信息。
func handleError(w http.ResponseWriter, err error) {
	if err != nil {
		// 记录错误信息
		log.Println(err)
		// 返回错误信息
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} 
}
// 使用日志库记录系统的运行信息和错误信息，方便后续的调试和维护。
func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
