package constant

const (
	/* Customer */
	INSERT_CUSTOMER   = "INSERT INTO m_customer(customer_id,username,password,customer_name,customer_address,customer_balance) VALUES (?,?,?,?,?,?)"
	CUSTOMER_LIST     = "SELECT customer_id,username,customer_name,customer_address,customer_balance FROM m_customer"
	CUSTOMER_GET      = "SELECT customer_id,username,customer_name,customer_address,customer_balance FROM m_customer WHERE customer_id=?"
	CUSTOMER_UPDATE   = "UPDATE m_customer SET customer_balance=? WHERE customer_id=?"
	CUSTOMER_GET_USER = "SELECT username,password from m_customer where username=?"

	/*Transaction*/
	INSERT_TRANSACTION = "INSERT INTO t_transaction(transaction_id,transaction_date,sender_id,reciever_id,transaction_amount) VALUES (?,?,?,?,?)"
)
