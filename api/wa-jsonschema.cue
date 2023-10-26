@jsonschema(schema="http://json-schema.org/draft-06/schema#")
#Transactions

#Transactions: {
	result: string
	message?: string
	cursor?: string
	count?:  int
	transactions?: [...#Transaction]
}

#Transaction: {
	blockchain?:        string
	symbol?:            string
	id?:                string
	transaction_type?:  string
	hash?:              string
	from?:              #AddressOwner
	to?:                #AddressOwner
	timestamp?:         int
	amount?:            number
	amount_usd?:        number
	transaction_count?: int
}

#AddressOwner: {
	address?:    string
	owner?:      string
	owner_type?: string
}
