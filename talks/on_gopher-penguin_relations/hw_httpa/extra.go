package main

/*
// START1 OMIT
type Handler interface {
	ServeHTTP(ResponseWriter, *Request) // HL
}
// END1 OMIT
*/

/*
// START2 OMIT
type ServeMux
	func NewServeMux() *ServeMux
	func (mux *ServeMux) Handle(pattern string, handler Handler)
	func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string)
	func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request) // HL
// END2 OMIT
*/

/*
// START3 OMIT
type HandlerFunc func(ResponseWriter, *Request)
// END3 OMIT
*/
