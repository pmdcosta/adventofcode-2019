package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var companies = []string{"c4aee7a2-9e2f-4bf5-9da3-9881a6b1676a", "71aae7e8-4b9c-45f0-a5d6-350eea69f069", "03e000a6-0d21-4eec-b328-9102013187d7", "5ffe07d5-25ab-461e-93de-4655733bbaf2", "3513597e-4597-41e7-9e8d-46d58c4104d8", "62bb4af9-2b4e-429b-b4c4-da389eff1e69", "c16834da-f8c7-45b3-936c-12796e44d0fb", "4cd3959c-71f7-445c-b47e-30b9097c1ff1", "f629808e-8352-4bd8-bca2-4d58e99f2718", "8537ba4e-f574-43e9-b86a-4a50d0e24590", "f19aa50f-8050-48be-9a82-859c4d83ae22", "80313100-ff39-485c-be13-bce49ff82702", "596795a8-cafd-4025-84eb-7a3283c0b490", "a205218d-abe9-4ce7-b757-59a69132a2f2", "d0e0f8ef-3d96-4d6a-a796-e95a5b5adade", "ca6b30d5-b826-4fdb-82be-efa204fb8c9d", "514122e7-184c-4de2-9da1-673644051a33", "1bbf2bc2-76e7-4463-bc11-d5f81ffea027", "71190655-590b-4602-a3e0-74650fd46220", "589d15f4-f57d-457c-ba34-bc70e581c424", "a7d40ed4-4487-415a-9b2f-24e5ecc16e6d", "e7487525-5c6c-48f4-bdd4-21f60a14591c", "ce0b65a4-f822-4b78-9d5f-db5f71c8887d", "aa3e6e79-2ecc-4bff-bb9a-24bb2809f1ea", "44eecb90-2e4e-453f-9fc7-6c89a968a07d", "f03c6122-0107-48d4-94ea-c938974eaf4b", "d4183d0b-2c0b-4586-8536-b39888328401", "f02405bc-005c-4cfd-94fa-451cb5a279bf", "d4f2ce51-3fb2-444a-be7a-ec58e13e4b4f", "73eefd48-5fbc-47cb-9d9a-76a2573ce86b", "02858ab0-83d8-4cbc-b07c-ceafbd28b5cc", "bcc80f95-9425-4b38-9122-d34626a154bb", "6dc3745b-9c74-44ec-8128-df0672b7d77d", "ab996bc4-f95d-4af7-8b1f-5e70c2d8f9d4", "bf38b7d7-4fe1-403e-bead-2c2d4febed51", "3ce6349f-13e3-4d5c-a9e8-e6aebbeca9b5", "2b0f3c66-b514-446d-9f1d-382ef00d2261", "55a0e45c-4b86-476a-9fbe-639d4bbd7941", "199a0709-f6a9-4385-b86b-92b66c70cce5", "518428df-193d-4d88-a5b0-b675507269cc", "18215cda-7b1f-4e85-b499-f005d17d5bf1", "89dd645f-ecf0-4ad5-9a90-9eed825ff521", "f5545e76-52c0-42b6-a0ac-eaf97f9699e2", "88b1ea7c-837a-4419-acc3-827e4054fcf3", "1d0c4c43-c7dc-4258-96be-95ef12a73278", "855381cf-6611-4c07-a50e-b6b39ede4e7d", "3dd131be-87af-47f4-b61b-6690df0c5a79", "69291d87-fb4c-4a42-bcd7-79e4dde247c5", "3e8abcec-e2c8-4bd3-bfbe-95477d1e41de", "5f1bc96a-c543-45c7-8b35-d85e93a4dced", "bcdfa855-c0f5-47d3-a1f5-659061103479", "512c0e54-4e37-4ec6-9c90-cb80f7c1e939", "733f581d-d2fd-4aa6-88a5-971e6bfd6edc", "45c72769-93c9-48db-b71c-27d96ae27762", "7b13ad21-a475-465a-80ae-24c72521b0df", "09df38b2-a46f-441e-8e28-0d6a4ff54df8", "ca5e2304-8e70-4952-9e7c-038b238267ba", "d1ca338c-b91a-4a7b-a41a-472c99ea5d76", "c6547c2e-6463-40eb-b43c-17ca2f1a9d16", "39e731b5-3db4-4811-8974-6a113704c027", "d43a5adb-1043-41cc-8c17-33235cc230b1", "d7eae0d6-36df-4a3e-bce1-186433118e21", "ca52280a-f0fe-4060-8a5d-1b2682596bb2", "b9748c1d-0b63-4786-aad1-d658e7156c06", "41fda195-5bf9-4bf9-ab2d-ae79ba1ea71f", "41fda195-5bf9-4bf9-ab2d-ae79ba1ea71f", "8ee5c5a1-e375-4c39-92aa-85f39529c249", "baaa60d9-6078-4b0d-ac3c-0bdb1ba0543a", "034fa726-4ce1-4e37-9e31-0f52aae4992b", "02eaf233-0459-485d-8121-61e05cda058c"}

func main() {
	user := "billing"
	password := "perkbox"
	host := "localhost:3306"
	database := "billing_contracts"
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, host, database))
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	for _, c := range companies {
		q := `"` + c + `"`
		results, err := db.Query(`SELECT id, stripe_id, subscription from customers WHERE id = ` + q)
		if err != nil {
			panic(fmt.Sprintf("%s - %s", err.Error(), q))
		}

		var r int
		for results.Next() {
			r++
			var company Company
			err = results.Scan(&company.ID, &company.StripeID, &company.Subscription)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			// and then print out the tag's Name attribute
			fmt.Printf("%+v\n", company)
		}
	}
}

type Company struct {
	ID           string `json:"id"`
	StripeID     string `json:"stripe_id"`
	Subscription string `json:"subscription"`
}
