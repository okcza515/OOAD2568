
public class AccountsReceivable {
	private CustomerTransaction transactionObject;
	
	public AccountsReceivable(CustomerTransaction aTransaction){
		transactionObject = aTransaction;
	}
	
	public void postPayment(){
		System.out.println("Payment posted for transaction...");
	}

	public void sendInvoice(){
		System.out.println("Prepare invoice...");
		// sends the invoice
	}
}
