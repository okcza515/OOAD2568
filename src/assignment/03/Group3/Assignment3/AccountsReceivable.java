
public class AccountsReceivable {
	private CustomerTransaction transactionObject;
	
	public AccountsReceivable(CustomerTransaction aTransaction){
		transactionObject = aTransaction;
	}
	
	public void postPayment(){
        System.out.println(String.format("Payment received from %s", transactionObject.getName()));
	}

	public void sendInvoice(){
        System.out.println(String.format("Invoice sent to %s", transactionObject.getName()));
	}
}

// 65070501023 Thanaphol Thangthaweesuk