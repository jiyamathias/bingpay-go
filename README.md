# Bingpay-go
Bingpay-go is a Go library that allows you to integrate BINGPAY into your Go project. Bingpay is a FinTech company which allows for Cheap Airtime &amp; Data Topup, Send &amp; Recieve Cash, Pay Utility Bills, Purchase Giftcards, Trade Airtime, Paypal &amp; Giftcards and International Topup.
# Please ensure to create issues in this repo if :
- You encounter any error while using this package and that issue would be attended to immediately.
# Get Started
- In other to use this package, you need to first create an account with bingpay via https://bingpay.ng/register
- After your account have been successfully created, you would need an API Token which is used to authorize users to use the bingpay api.

## How to get the Bingpay API token
- After you signin to your Bingpay dashboard, click on the top right corner which has your profile picture and username
- On the pop up, click on the "developer api" option and you would get a form to fill which would request for two things: your website address and the reason why you need the developer API. After those have been inputed then you can submit your request which would be under review for some day(s) before you get an API token. 

# Installation
To install this bingpay package, you need to install [Go](https://golang.org/) and set your Go workspace first.
1. You can use the below Go command to install BingPay
```sh
$ go get -u github.com/jiyamathias/bingpay-go
```
2. Import it in your code:
```sh
import "github.com/jiyamathias/bingpay-go"
```
## Note : All methods in this package returns three (3) things:
- [x] The object of the response
- [x] An int (status code) i.e  status 200 or status 400
- [x] An error (if any)

# Quick start
```sh
# assume the following codes in example.go file
$ touch example.go
# open the just created example.go file in the text editor of your choice
```
## Wallet 
- ### Check Balance
Use this to fetch your bingpay wallet balance
```go
package main

import (
	"fmt"
	bingpay "github.com/jiyamathias/bingpay-go"
	"github.com/jiyamathias/bingpay-go/bingpayBal"
)

func main() {
	bingpay.Token = "your bingpay api secret key" // to add your secret key for the api requests 

	response, status, err := bingpayBal.CheckBalance()
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(status)
	fmt.Println(response)
}
```
## Airtime
- ### All Networks
Use this to fetch the list of all networks supported by bingpay

```go
package main

import (
	"fmt"
	bingpay "github.com/jiyamathias/bingpay-go"
	"github.com/jiyamathias/bingpay-go/airtime"
)

func main() {
	bingpay.Token = "your bingpay api secret key" // to add your secret key for the api requests 

	response, status, err := airtime.AllNetworks()
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(status)
	fmt.Println(response)
}
```
- ### Verify Phone Number
Use this to verify customer's phone number.
Only country ISO codes are allowed. example NG for Nigeria, US for United States of America etc.
You can find list of all ISO codes [here](https://www.nationsonline.org/oneworld/country_code_list.htm)

```go
package main

import (
	"fmt"
	bingpay "github.com/jiyamathias/bingpay-go"
	"github.com/jiyamathias/bingpay-go/airtime"
)

func main() {
	bingpay.Token = "your bingpay api secret key" // to add your secret key for the api requests 

	country := "NG" // Country ISO code e.g NG for nigeria
	number := "08000000000" // Customer's phone number.
	response, status, err := airtime.VerifyPhoneNumber(country, number)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(status)
	fmt.Println(response)
}
```
- ### Buy Airtime
Use this to perform airtime purchase.
You get 2% discount on every airtime purchase instantly.

```go
package main

import (
	"fmt"
	bingpay "github.com/jiyamathias/bingpay-go"
	"github.com/jiyamathias/bingpay-go/airtime"
)

func main() {
	bingpay.Token = "your bingpay api secret key" // to add your secret key for the api requests 

	phone := "08000000000" // Phone number to recharge.
	amount := 100 // Amount to recharge. e.g 100 for 100 naria topup
	network_id := 0 // Network (As seen in the all-networks endpoint).
	response, status, err := airtime.BuyAirtime(phone,amount,network_id)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(status)
	fmt.Println(response)
}
```

## Data
- ### All data plans
Use this to fetch all data plans

```go
package main

import (
	"fmt"
	bingpay "github.com/jiyamathias/bingpay-go"
	"github.com/jiyamathias/bingpay-go/data"
)

func main() {
	bingpay.Token = "your bingpay api secret key" // to add your secret key for the api requests 

	response, status, err := data.AllDataPlans()
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(status)
	fmt.Println(response)
}
```
- ### Data plans
Fetch Data plans for a specific network.
Network id can be gotten from All networks endpoint.

```go
package main

import (
	"fmt"
	bingpay "github.com/jiyamathias/bingpay-go"
	"github.com/jiyamathias/bingpay-go/data"
)

func main() {
	bingpay.Token = "your bingpay api secret key" // to add your secret key for the api requests 

	/* network ID's
			MTN = 1
			Airtel = 2
			9mobile = 3
			Glo = 4
	*/

	network_id := 0 // Network (As seen in the all-networks endpoint).
	response, status, err := data.DataPlans(network_id)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(status)
	fmt.Println(response)
}
```
- ### Buy Data
Use this to perform a data purchase.
You get 2% discount on every data purchase instantly.

```go
package main

import (
	"fmt"
	bingpay "github.com/jiyamathias/bingpay-go"
	"github.com/jiyamathias/bingpay-go/data"
)

func main() {
	bingpay.Token = "your bingpay api secret key" // to add your secret key for the api requests 

	/* network ID's
			MTN = 1
			Airtel = 2
			9mobile = 3
			Glo = 4
	*/
	phone := "08000000000" // Phone number to recharge.
	plan := 0 // Data Plan (As seen as in the data-plans endpoint as 'uniq_id' ).
	network_id := 0 // // Network (As seen in the all-networks endpoint).
	response, status, err := data.BuyData(phone, plan, network_id)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(status)
	fmt.Println(response)
}
```
## Utility Bills
- ### All Services
Fetch all available bills services from this endpoint.

```go
package main

import (
	"fmt"
	bingpay "github.com/jiyamathias/bingpay-go"
	"github.com/jiyamathias/bingpay-go/utilitybills"
)

func main() {
	bingpay.Token = "your bingpay api secret key" // to add your secret key for the api requests 

	response, status, err := utilitybills.AllServices()
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(status)
	fmt.Println(response)
}
```
- ### Service Variation
Fetch variation for a bill service.
Service id can be gotten from all-services endpoint.

```go
package main

import (
	"fmt"
	bingpay "github.com/jiyamathias/bingpay-go"
	"github.com/jiyamathias/bingpay-go/utilitybills"
)

func main() {
	bingpay.Token = "your bingpay api secret key" // to add your secret key for the api requests 

	service_id := 0 //as gotten from the "Fetch all available bills services" endpoint
	response, status, err := utilitybills.ServiceVariation(service_id)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(status)
	fmt.Println(response)
}
```
- ### Verify Customer
Use this to validate customer's details before performing a bill purchase.

```go
package main

import (
	"fmt"
	bingpay "github.com/jiyamathias/bingpay-go"
	"github.com/jiyamathias/bingpay-go/utilitybills"
)

func main() {
	bingpay.Token = "your bingpay api secret key" // to add your secret key for the api requests 

	service_id := "" // Service id (As seen in the all-services endpoint).
	customer_id := "" // Customer meter number, smart card number, etc
	meter_type := "" // Meter type, (prepaid or postpaid) required for electricity bills only.
	response, status, err := utilitybills.VerifyCustomer(service_id, customer_id, meter_type)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(status)
	fmt.Println(response)
}
```
- ### Purchase Bills
Use this to perform a bill purchase.

```go
package main

import (
	"fmt"
	bingpay "github.com/jiyamathias/bingpay-go"
	"github.com/jiyamathias/bingpay-go/utilitybills"
)

func main() {
	bingpay.Token = "your bingpay api secret key" // to add your secret key for the api requests 

	service_id := "" // Service id (As seen in the all-services endpoint).
	customer_id := "" // Customer meter number, smart card number, etc
	variation := "" // For services that has variation (As seen in the service variation endpoint).
	amount := "" // For services has custom amount. Example: electricity bill paymen
	response, status, err := utilitybills.PurchaseBill(service_id, customer_id, variation, amount)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(status)
	fmt.Println(response)
}
```
## Airtime To Cash
- ### Network Fee
Fetch network charges and recipient phone number for each network before processing airtime to cash transactions.
The data->charge response is the percentage of charge and the data->value response is the amount you'll receive if the conversion is approved.

```go
package main

import (
	"fmt"
	bingpay "github.com/jiyamathias/bingpay-go"
	"github.com/jiyamathias/bingpay-go/airtimetocash"
)

func main() {
	bingpay.Token = "your bingpay api secret key" // to add your secret key for the api requests 

	amount := 0 // Amount to convert.
	network_id := 0 // Network id (As seen in the all-networks endpoint).

	response, status, err := airtimetocash.NetworkFee(amount, network_id)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(status)
	fmt.Println(response)
}
```
- ### Process Airtime To Cash
Use this to perform airtime to cash transactions.
Transactions may take 5-10 mins to be approved, You can automatically get notification if the transaction is approved or declined by using our Webhooks.

```go
package main

import (
	"fmt"
	bingpay "github.com/jiyamathias/bingpay-go"
	"github.com/jiyamathias/bingpay-go/airtimetocash"
)

func main() {
	bingpay.Token = "your bingpay api secret key" // to add your secret key for the api requests 

	amount := 0 // Amount to convert.
	network_id := 0 // Network id (As seen in the all-networks endpoint).
	phone := "080000000000" // Sender's phone number.
	response, status, err := airtimetocash.AirtimeToCash(amount, network_id, phone)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(status)
	fmt.Println(response)
}
```
## Identity Verification
- ### Bvn Verification
Bank Verification Number or BVN is a biometric identification widely used by Banks in Nigeria. This endpoint allows you to fetch a user's information from their BVN. It supports only BVNs generated by Nigerian Banks.

```go
package main

import (
	"fmt"
	bingpay "github.com/jiyamathias/bingpay-go"
	"github.com/jiyamathias/bingpay-go/indentifyVerifications"
)

func main() {
	bingpay.Token = "your bingpay api secret key" // to add your secret key for the api requests 

	firstName := "" // Customer's first name
	lastName := "" // Customer's last name.
	phone := "" // Customer's phone number.
	bvn := "" // Customer's bvn.
	response, status, err := indentifyverifications.VerifyBvn(firstName, lastName, phone, bvn)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(status)
	fmt.Println(response)
}
```
- ### Nin Verification
This endpoint allows you to verify a Nigeria National ID.
Service charge of NGN80 applies for every successful data fetch.

```go
package main

import (
	"fmt"
	bingpay "github.com/jiyamathias/bingpay-go"
	"github.com/jiyamathias/bingpay-go/indentifyverifications"
)

func main() {
	bingpay.Token = "your bingpay api secret key" // to add your secret key for the api requests 

	firstName := "" // Customer's first name
	lastName := "" // Customer's last name.
	phone := "" // Customer's phone number.
	bvn := "" // Customer's bvn.
	response, status, err := indentifyverifications.VerifyNin(firstName, lastName, phone, bvn)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(status)
	fmt.Println(response)
}
```
## Banks
- ### All Banks
Fetch all Nigerian banks with this endpoint

```go
package main

import (
	"fmt"
	bingpay "github.com/jiyamathias/bingpay-go"
	"github.com/jiyamathias/bingpay-go/bank"
)

func main() {
	bingpay.Token = "your bingpay api secret key" // to add your secret key for the api requests 

	response, status, err := bank.AllBanks()
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(status)
	fmt.Println(response)
}
```
- ### Resolve Account
Use this to resolve customer's account details.

```go
package main

import (
	"fmt"
	bingpay "github.com/jiyamathias/bingpay-go"
	"github.com/jiyamathias/bingpay-go/bank"
)

func main() {
	bingpay.Token = "your bingpay api secret key" // to add your secret key for the api requests 

	bank_code := "" // Bank Code (As seen in All-banks endpoint).
	acct_num := "" // Customer's account number.
	response, status, err := bank.ResolveAccount(bank_code, acct_num)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(status)
	fmt.Println(response)
}
```
```sh
# run example.go 
$ go run example.go
```