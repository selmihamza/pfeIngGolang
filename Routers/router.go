package routers

import (
	"driver/controllers"

	"github.com/gorilla/mux"
)

// InitializeRouter function
func InitializeRouter() *mux.Router {

	router := mux.NewRouter()

	// http.Handle("/", http.StripPrefix(strings.TrimRight(path, "/"), http.FileServer(http.Dir(directory))))

	router.HandleFunc("/", controllers.Hello).Methods("GET")
	router.HandleFunc("/login", controllers.Auth).Methods("POST")
	// get offices by saas compnay
	router.HandleFunc("/{saas_company}/offices", controllers.GetOfficesBySaasCompany).Methods("GET")
	router.HandleFunc("/saascompanyAndoffice/{driver_id}", controllers.GetSaasCompanyAndOffice).Methods("GET")
	// user routers
	router.HandleFunc("/register", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/driver/profile/{id}", controllers.GetDriverProfileByID).Methods("GET")
	router.HandleFunc("/driver/{id}", controllers.GetDriverInformationsByID).Methods("GET")
	router.HandleFunc("/update_driver_status/{driverID}", controllers.UpdateDriverStatus).Methods("PUT")
	router.HandleFunc("/update_password/{driverID}", controllers.UpdatePassword).Methods("PUT")
	// company routers
	router.HandleFunc("/company", controllers.CreateCompany).Methods("POST")
	router.HandleFunc("/company/{id}", controllers.GetCompanyByID).Methods("GET")
	router.HandleFunc("/company", controllers.UpdateCompany).Methods("PUT")
	// vehicule routers
	router.HandleFunc("/vehicle_types/{saas_officeID}", controllers.GetVehicleTypes).Methods("GET")
	router.HandleFunc("/vehicle_options/{saas_officeID}", controllers.GetVehicleOptions).Methods("GET")
	router.HandleFunc("/vehicle_categories/{saas_officeID}", controllers.GetVehicleCategories).Methods("GET")
	router.HandleFunc("/vehicle_zones/{saas_officeID}", controllers.GetVehicleZones).Methods("GET")
	router.HandleFunc("/vehicle", controllers.CreateVehicle).Methods("POST")
	// bookings routers
	router.HandleFunc("/request/{DriverID}", controllers.GetRequestByID).Methods("GET")
	router.HandleFunc("/listrequest/{DriverID}", controllers.GetListRequestByID).Methods("GET")
	// reporting routers
	router.HandleFunc("/reportings/{DriverID}", controllers.GetListReportingsByID).Methods("GET")
	// agenda routers
	router.HandleFunc("/agenda", controllers.CreateAgenda).Methods("POST")
	router.HandleFunc("/agenda/{DriverID}", controllers.GetAgendaByID).Methods("GET")
	router.HandleFunc("/agenda/confirm/{ID}", controllers.ConfirmPeriod).Methods("PUT")
	router.HandleFunc("/agenda/delete/{ID}", controllers.DeletePeriod).Methods("PUT")
	router.HandleFunc("/agenda/update/{ID}", controllers.UpdatePeriod).Methods("PUT")

	return router
}
