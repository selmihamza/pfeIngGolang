package controllers

import (
	"driver/Helpers"
	"driver/Models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// CreateCompany Handler
// @Summary Create a new company
// @Description Create a new company with the input payload
// @Tags driver_companies
// @Accept  mpfd
// @Produce  json
// @Param DriverID query int true "enter driver id (tester avec l'id : 17357 ) : "
// @Param CompanyName query string false "enter CompanyName:"
// @Param PcoNumber query string false "enter CompanyDre:"
// @Param VatNumber query string false "enter CompanyTva:"
// @Param Registration query string false "enter CompanySiren:"
// @Param VatRate query float64 false "enter VatPercentage:"
// @Param RIB formData file false "enter rib file:"
// @Param INSEE formData file false "enter insee file:"
// @Param DRE formData file false "enter dre file:"
// @Success 200 {object} Models.DriverCompanie
// @Router /company [post]
func CreateCompany(w http.ResponseWriter, r *http.Request) {
	db := Helpers.DbConnect()
	var driver Models.Driver
	DriverID, _ := strconv.Atoi(r.FormValue("DriverID"))
	var countID int
	db.Table("drivers").Where("id = ?", DriverID).Count(&countID)

	if countID != 0 {
		Vat, err := strconv.ParseFloat(r.FormValue("VatRate"), 64)
		if err != nil {
			w.WriteHeader(http.StatusNotAcceptable)
			w.Write([]byte("TVA rate must be a number!"))
			return
		}
		company := Models.DriverCompanie{CompanyName: r.FormValue("CompanyName"), CompanyDre: r.FormValue("PcoNumber"), CompanyTva: r.FormValue("VatNumber"), CompanySiren: r.FormValue("Registration"), CreatedAt: time.Now(), UpdatedAt: time.Now(), VatPercentage: Vat}
		db.NewRecord(company)
		db.Create(&company)
		db.Model(&driver).Where("id = ?", DriverID).Update("driver_company_id", company.ID)

		//---START UPLOAD RIB FILE--------------------------------------------------------------
		file, _, err := r.FormFile("RIB")
		if file != nil {
			if err != nil {
				fmt.Println("Error Retrieving the File")
				fmt.Println(err)
				return
			}
			defer file.Close()
			tempFile, err := ioutil.TempFile("public/company/temp-images", "rib-*.png")
			nameRib := tempFile.Name()[strings.Index(tempFile.Name(), "rib-"):]
			if err != nil {
				fmt.Println(err)
			}
			defer tempFile.Close()
			fileBytes, err := ioutil.ReadAll(file)
			if err != nil {
				fmt.Println(err)
			}
			tempFile.Write(fileBytes)
			fileTypeRib := Models.AdministrativeFileType{Designation: "Relevé d'Identité Bancaire", Code: "RIB", CreatedAt: time.Now(), UpdatedAt: time.Now()}
			db.NewRecord(fileTypeRib)
			db.Create(&fileTypeRib)
			fileCompanyRib := Models.FileCompanie{Ref: nameRib, FileType: fileTypeRib.ID, CreatedAt: time.Now(), UpdatedAt: time.Now(), DriverCompanyID: company.ID}
			db.NewRecord(fileCompanyRib)
			db.Create(&fileCompanyRib)

		}
		//---END UPLOAD RIB FILE---------------------------------------------------------------
		//---START UPLOAD INSEE FILE-----------------------------------------------------------
		fileInsee, _, err := r.FormFile("INSEE")
		if fileInsee != nil {
			if err != nil {
				fmt.Println("Error Retrieving the File")
				fmt.Println(err)
				return
			}
			defer fileInsee.Close()
			tempFileInsee, err := ioutil.TempFile("public/company/temp-images", "insee-*.png")
			nameInsee := tempFileInsee.Name()[strings.Index(tempFileInsee.Name(), "insee-"):]
			if err != nil {
				fmt.Println(err)
			}
			defer tempFileInsee.Close()
			fileBytesInsee, err := ioutil.ReadAll(fileInsee)
			if err != nil {
				fmt.Println(err)
			}
			tempFileInsee.Write(fileBytesInsee)
			fileTypeInsee := Models.AdministrativeFileType{Designation: "Kbis si société / INSEE si autoentrepreneur", Code: "INSEE", CreatedAt: time.Now(), UpdatedAt: time.Now()}
			db.NewRecord(fileTypeInsee)
			db.Create(&fileTypeInsee)
			fileCompanyInsee := Models.FileCompanie{Ref: nameInsee, FileType: fileTypeInsee.ID, CreatedAt: time.Now(), UpdatedAt: time.Now(), DriverCompanyID: company.ID}
			db.NewRecord(fileCompanyInsee)
			db.Create(&fileCompanyInsee)

		}
		//---END UPLOAD INSEE FILE--------------------------------------------------------------
		//---START UPLOAD DRE FILE--------------------------------------------------------------
		fileLicense, _, err := r.FormFile("DRE")
		if fileLicense != nil {
			if err != nil {
				fmt.Println("Error Retrieving the File")
				fmt.Println(err)
				return
			}
			defer fileLicense.Close()
			tempFileLicense, err := ioutil.TempFile("public/company/temp-images", "dre-*.png")
			nameDre := tempFileLicense.Name()[strings.Index(tempFileLicense.Name(), "dre-"):]
			if err != nil {
				fmt.Println(err)
			}
			defer tempFileLicense.Close()
			fileBytesLicense, err := ioutil.ReadAll(fileLicense)
			if err != nil {
				fmt.Println(err)
			}
			tempFileLicense.Write(fileBytesLicense)
			fileTypeLicense := Models.AdministrativeFileType{Designation: "Licence Atout France ou Capacité DRE", Code: "DRE", CreatedAt: time.Now(), UpdatedAt: time.Now()}
			db.NewRecord(fileTypeLicense)
			db.Create(&fileTypeLicense)
			fileCompanyLicense := Models.FileCompanie{Ref: nameDre, FileType: fileTypeLicense.ID, CreatedAt: time.Now(), UpdatedAt: time.Now(), DriverCompanyID: company.ID}
			db.NewRecord(fileCompanyLicense)
			db.Create(&fileCompanyLicense)

		}
		//---END UPLOAD INSEE FILE--------------------------------------------------------------
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(company.ID)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized!"))
		return
	}

}

// FileData struct
type FileData struct {
	URL  string
	File Models.FileCompanie
	Type Models.AdministrativeFileType
}

// CompanyData struct
type CompanyData struct {
	Company Models.DriverCompanie
	Files   []FileData
}

// GetCompanyByID Handler
// @Summary retrieve a company by ID
// @Description retrieve a company by ID in database with the ID in request parameter
// @Tags driver_companies
// @Accept  json
// @Produce  json
// @Param id query int true "enter id:"
// @Success 200 {object} Models.DriverCompanie
// @Router /company [get]
func GetCompanyByID(w http.ResponseWriter, r *http.Request) {
	db := Helpers.DbConnect()
	params := mux.Vars(r)
	id := params["id"]

	var company Models.DriverCompanie
	db.First(&company, id)
	if company.ID != 0 {
		var files []Models.FileCompanie
		db.Where("driver_company_id = ?", id).Find(&files)
		var FileTemp []FileData
		if len(files) > 0 {
			for i := 0; i < len(files); i++ {
				var fileType Models.AdministrativeFileType
				db.Where("id = ?", files[i].FileType).Find(&fileType)
				temp := FileData{URL: "public/company/temp-images/" + files[i].Ref, File: files[i], Type: fileType}

				FileTemp = append(FileTemp, temp)
			}
		}
		data := CompanyData{Company: company, Files: FileTemp}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - data not found"))
	}

}

// UpdateCompany Handler
// @Summary update a company
// @Description update a company by id with the input payload
// @Tags driver_companies
// @Accept  mpfd
// @Produce  json
// @Param id query int true "enter id:"
// @Param CompanyName query string false "enter CompanyName:"
// @Param PcoNumber query string false "enter CompanyDre:"
// @Param VatNumber query string false "enter CompanyTva:"
// @Param Registration query string false "enter CompanySiren:"
// @Param VatRate query float64 false "enter VatPercentage:"
// @Success 200 {object} Models.DriverCompanie
// @Router /company [put]
func UpdateCompany(w http.ResponseWriter, r *http.Request) {
	db := Helpers.DbConnect()

	id := r.URL.Query().Get("id")
	var company Models.DriverCompanie
	db.First(&company, id)
	if company.ID != 0 {
		var FileTemp []FileData
		// update file rib start ----------------------------------------------
		var ribFile Models.FileCompanie
		r.FormValue("CompanyName")
		db.First(&ribFile, r.FormValue("ribFileID"))
		path := "public/company/temp-images/" + ribFile.Ref
		err := os.Remove(path)
		file, _, err := r.FormFile("RIB")
		if file != nil {
			if err != nil {
				fmt.Println("Error Retrieving the File")
				fmt.Println(err)
				return
			}
			defer file.Close()
			tempFile, err := ioutil.TempFile("public/company/temp-images", "rib-*.png")
			if err != nil {
				fmt.Println(err)
			}
			nameRib := tempFile.Name()[strings.Index(tempFile.Name(), "rib-"):]
			defer tempFile.Close()
			fileBytes, err := ioutil.ReadAll(file)
			if err != nil {
				fmt.Println(err)
			}
			tempFile.Write(fileBytes)
			db.Model(&ribFile).Updates(Models.FileCompanie{Ref: nameRib, UpdatedAt: time.Now()})

			temp := FileData{URL: "public/company/temp-images/" + ribFile.Ref, File: ribFile}
			FileTemp = append(FileTemp, temp)
		}
		// update file rib end -------------------------------------------------
		// update file insee start ---------------------------------------------
		var inseeFile Models.FileCompanie
		r.FormValue("CompanyName")
		db.First(&inseeFile, r.FormValue("inseeFileID"))
		pathInsee := "public/company/temp-images/" + inseeFile.Ref
		os.Remove(pathInsee)
		fileInsee, _, err := r.FormFile("INSEE")
		if fileInsee != nil {
			if err != nil {
				fmt.Println("Error Retrieving the File")
				fmt.Println(err)
				return
			}
			defer fileInsee.Close()
			tempFile, err := ioutil.TempFile("public/company/temp-images", "insee-*.png")
			if err != nil {
				fmt.Println(err)
			}
			nameInsee := tempFile.Name()[strings.Index(tempFile.Name(), "insee-"):]
			defer tempFile.Close()
			fileBytes, err := ioutil.ReadAll(fileInsee)
			if err != nil {
				fmt.Println(err)
			}
			tempFile.Write(fileBytes)
			db.Model(&inseeFile).Updates(Models.FileCompanie{Ref: nameInsee, UpdatedAt: time.Now()})

			temp := FileData{URL: "public/company/temp-images/" + inseeFile.Ref, File: inseeFile}
			FileTemp = append(FileTemp, temp)
		}
		// update file INSEE end ------------------------------------------------
		// update file DRE start ------------------------------------------------
		var dreFile Models.FileCompanie
		r.FormValue("CompanyName")
		db.First(&dreFile, r.FormValue("dreFileID"))
		pathDre := "public/company/temp-images/" + dreFile.Ref
		os.Remove(pathDre)

		fileDre, _, err := r.FormFile("DRE")
		if fileDre != nil {
			if err != nil {
				fmt.Println("Error Retrieving the File")
				fmt.Println(err)
				return
			}
			defer fileDre.Close()
			tempFile, err := ioutil.TempFile("public/company/temp-images", "dre-*.png")
			if err != nil {
				fmt.Println(err)
			}
			nameDre := tempFile.Name()[strings.Index(tempFile.Name(), "dre-"):]
			defer tempFile.Close()
			fileBytes, err := ioutil.ReadAll(fileDre)
			if err != nil {
				fmt.Println(err)
			}
			tempFile.Write(fileBytes)
			db.Model(&dreFile).Updates(Models.FileCompanie{Ref: nameDre, UpdatedAt: time.Now()})
			temp := FileData{URL: "public/company/temp-images/" + dreFile.Ref, File: dreFile}
			FileTemp = append(FileTemp, temp)
		}
		// update file DRE end ---------------------------------------------------
		Vat, _ := strconv.ParseFloat(r.FormValue("VatRate"), 64)
		db.Model(&company).Updates(Models.DriverCompanie{CompanyName: r.FormValue("CompanyName"), CompanyDre: r.FormValue("PcoNumber"), CompanyTva: r.FormValue("VatNumber"), CompanySiren: r.FormValue("Registration"), UpdatedAt: time.Now(), VatPercentage: Vat})
		w.WriteHeader(http.StatusOK)

		d := CompanyData{Company: company, Files: FileTemp}
		json.NewEncoder(w).Encode(d)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - page not found"))
	}
}
