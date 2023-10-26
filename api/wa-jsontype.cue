Transactions: {
	result: type: "string"
	cursor: type: "string"
	count: type: "integer"
	transactions: {
		type: "array"
		items: $ref: "#/Transaction"
	}
}
Transaction: {
	blockchain: type: "string"
	symbol: type: "string"
	id: type: "string"
	transaction_type: type: "string"
	hash: type: "string"
	from: $ref: "#/AddressOwner"
	to: $ref: "#/AddressOwner"
	timestamp: type: "integer"
	amount: type: "number"
	amount_usd: type: "number"
	transaction_count: type: "integer"
}
AddressOwner: {
	address: type: "string"
	owner: type: "string"
	owner_type: type: "string"
}
