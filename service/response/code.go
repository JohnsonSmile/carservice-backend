package response

/*
	1. logger.Error("param error", err)
	2. unauthorized, userId 获取失败的那种, 在middleware中需要做处理
	3. get unfinished highway order error
	4. wrong qrcode, should go to the entrance
	5. get position error
	6. wrong qrcode, you should go to the exit
	7. get position error
	8. get fee error
	9. you should pay previous order first
	10. 已经准备支付了ready to pay
	12. create order failed 创建订单失败

*/

const (
	SUCCESS       = 200
	UNKNOWN_ERROR = 5001

	PARAM_ERROR                     = iota + 1000 // 参数错误
	UNAUTHORIZED_ERROR                            // jwt失效错误
	GET_UNFINISHED_ORDER_ERROR                    // 获取未完成订单错误
	WRONG_QRCODE_ENTRANCE_ERROR                   // 错误的二维码进入错误
	WRONG_QRCODE_EXIT_ERROR                       // 错误的二维码出口错误
	WRONG_QRCODE_CHARGE_START_ERROR               // 错误的开始充电二维码
	WRONG_QRCODE_CHARGE_END_ERROR                 // 错误的结束充电二维码
	GET_POSITION_ERROR                            // 获取position错误
	GET_FEE_ERROR                                 // 获取fee错误
	PREVIOUS_ORDER_FIRST_ERROR                    // 应该支付上一个订单
	READY_TO_PAY_ERROR                            // 已经准备支付了ready to pay
	CREATE_ORDER_FAILED                           // 创建订单失败
)

func MsgForCode(code int) string {

	switch code {
	case SUCCESS:
		return "success"
	case UNKNOWN_ERROR:
		return "未知错误"
	case PARAM_ERROR:
		return "参数错误"

	case UNAUTHORIZED_ERROR:
		return "登录失效，请重新登录"
	case GET_UNFINISHED_ORDER_ERROR:
		return "获取未完成订单错误"
	case WRONG_QRCODE_ENTRANCE_ERROR:
		return "当前二维码不是入口的二维码"
	case WRONG_QRCODE_EXIT_ERROR:
		return "当前二维码不是出口的二维码"
	case WRONG_QRCODE_CHARGE_START_ERROR:
		return "当前二维码不是开始充电的二维码"
	case WRONG_QRCODE_CHARGE_END_ERROR:
		return "当前二维码不是结束充电的二维码"
	case GET_POSITION_ERROR:
		return "获取当前位置信息错误"
	case GET_FEE_ERROR:
		return "获取费用错误"
	case PREVIOUS_ORDER_FIRST_ERROR:
		return "您还有未支付的订单，请先支付"
	case READY_TO_PAY_ERROR:
		return "已经结束订单，请支付"
	default:
		return "未知错误"

	}

}
