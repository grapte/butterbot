Transactions: {
	result: "success" | "error"
	if result == "error" {
		message: string
	}
	if result == "success" {
		cursor: string
		count: int
		transaction: [...Transaction]
	}
}

Transaction: {
	blockchain: string
	symbol: string
	id: string
	transaction_type: string
	hash: string
	from: AddressOwner
	to: AddressOwner
	timestamp: int
	amount: number
	amount_usd: string
	transaction_count: int
}

AddressOwner: {
	address: string
	owner: string
	owner_type?: string
}