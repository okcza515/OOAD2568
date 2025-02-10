//65070501008
public class AccountsReceivable {
	private CustomerTransaction transactionObject;
	
	public AccountsReceivable(CustomerTransaction aTransaction){
		transactionObject = aTransaction;
	}
	
	public void postPayment(){
		System.out.println("Payment processed for " + transactionObject.getName());
	}

	public void sendInvoice(){
		System.out.println("Invoice sent to " + transactionObject.getName());
	}
}
