type error interface {
	Error() string //Error() 메소드가 있고, string만 반환하면 무슨 타입이든 에러 타입이 될 수 있음.
}