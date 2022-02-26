package controllers

import "caseMajoo/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/api/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	// Merchant Route
	s.Router.HandleFunc("/api/merchant/report", middlewares.SetMiddlewareAuthentication(middlewares.SetMiddlewareAuthentication(s.GetMerchantReport))).Methods("POST")

	// Outlet Route
	s.Router.HandleFunc("/api/outlet/report", middlewares.SetMiddlewareAuthentication(middlewares.SetMiddlewareAuthentication(s.GetOutletReport))).Methods("POST")

	// Area Route
	s.Router.HandleFunc("/api/area", middlewares.SetMiddlewareJSON(s.Area)).Methods("GET")

	s.Router.HandleFunc("/deret", middlewares.SetMiddlewareJSON(s.Deret)).Methods("GET")

}
