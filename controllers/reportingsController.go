package controllers

import (
	"driver/Helpers"
	"driver/Models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Invoices struct
type Invoices struct {
	Invoice       Models.DriverInvoice
	PaymentPeriod Models.PaymentPeriod
}

// GetListReportingsByID function
// @Summary retrieve a reporting by driver ID
// @Description retrieve a reporting in database with the driver ID
// @Tags driver_invoice
// @Accept  json
// @Produce  json
// @Param DriverID path int true "enter driver id :11469:"
// @Success 200 {object} Models.DriverInvoice
// @Router /reportings/{DriverID} [get]
func GetListReportingsByID(w http.ResponseWriter, r *http.Request) {
	db := Helpers.DbConnect()
	params := mux.Vars(r)
	DriverID := params["DriverID"]
	var invoices []Models.DriverInvoice
	var ListInvoices []Invoices
	db.Model(&invoices).Where("driver_id = ?", DriverID).Order("created_at desc").Find(&invoices)
	for i := 0; i < len(invoices); i++ {
		var period Models.PaymentPeriod
		db.First(&period, invoices[i].PeriodID)

		invoice := Invoices{invoices[i], period}
		ListInvoices = append(ListInvoices, invoice)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ListInvoices)

}
