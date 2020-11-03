package mock

var UserDataCreate = []byte(`{"FirstName": "selmi","LastName": "hamza","Email": "kilxdc@gmail.com","Password": "123456","Type":"Manager","PhoneNumber": "21655664455","HomeAddress": "nabeul tunisie","HomeLatitude":38.8951,"HomeLongitude": -77.0364,   "Country": "tunis","LicenceNumber":"789654123","Office":1,"SaasCompanyID":1}`)
var CompanyDataCreate = []byte(`map[string]string{"CompanyName": "Company_Name", "Registration": "111-xxx-111", "LicenseNumber": "98979897", "VatNumber": "TN-071010107", "VatRate": "20", "PcoNumber": "EVTC107107","DriverID": "17364"}`)
var CompanyDataUpdate = []byte(`map[string]string{"CompanyName": "CompanyName", "Registration": "222-xxx-222", "LicenseNumber": "123654789", "VatNumber": "FR-071010107", "VatRate": "10", "PcoNumber": "EVTC701701"}`)
var VehicleDataCreate = []byte(`map[string]string{"Color":"#000000","Matricule":"AB-107-fk","Year":"2020","Model":"BMW","Mark":"X1","Location":"Location","CarInUse":"true","VehicleCategoryID":"590""DriverID":"11469","DriverCompanyID":"10725","SaasOfficeID":"1"}`)
