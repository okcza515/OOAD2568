
public class AccountsReceivable {
	private CustomerTransaction transactionObject;
	
	public AccountsReceivable(CustomerTransaction aTransaction){
		transactionObject = aTransaction;
	}
	
	public void postPayment(){
		transactionObject.chargeCustomer();
	}

	public void sendInvoice(){
		transactionObject.prepareInvoice();
		// sends the invoice
	}
}
