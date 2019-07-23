package main

import (
	"awesomeProject/endpoints"
	"awesomeProject/endpoints/cookie"
	transportCookie "awesomeProject/endpoints/cookie/transport"
	"awesomeProject/endpoints/count"
	transportCount "awesomeProject/endpoints/count/transport"
	"awesomeProject/endpoints/uppercase"
	transportUpperCase "awesomeProject/endpoints/uppercase/transport"

	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	svc := endpoints.New()

	uppercaseHandler := httptransport.NewServer(
		uppercase.MakeUppercaseEndpoint(svc),
		transportUpperCase.DecodeUppercaseRequest,
		transportUpperCase.EncodeResponse,
	)

	countHandler := httptransport.NewServer(
		count.MakeCountEndpoint(svc),
		transportCount.DecodeCountRequest,
		transportCount.EncodeResponse,
	)

	cookieHandler := httptransport.NewServer(
		cookie.MakeCookieEndpoint(svc),
		transportCookie.DecodeCookieRequest,
		transportCookie.EncodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)

	http.Handle("/cookie", cookieHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
